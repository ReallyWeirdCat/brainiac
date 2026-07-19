package security

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/config"
)

var (
	ErrInvalidToken            = errors.New("invalid token")
	ErrExpiredToken            = errors.New("token has expired")
	ErrUnexpectedScope         = errors.New("unexpected token scope")
	ErrMissingClaims           = errors.New("token claims are missing")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
)

type jwtProvider struct {
	secret          []byte
	issuer          string
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	totpTokenTTL    time.Duration
}

var _ ports.TokenGenerator = &jwtProvider{}

func NewTokenProvider(cfg config.AppConfigProvider) (ports.TokenProvider, error) {
	c := cfg.Get().Security.JWT
	if len(c.Secret) == 0 {
		return nil, errors.New("token provider secret must not be empty")
	}
	if c.AccessTokenTTL <= 0 || c.RefreshTokenTTL <= 0 || c.TOTPTokenTTL <= 0 {
		return nil, errors.New("token TTLs must be positive")
	}
	return &jwtProvider{
		secret:          []byte(c.Secret),
		issuer:          c.Issuer,
		accessTokenTTL:  c.AccessTokenTTL,
		refreshTokenTTL: c.RefreshTokenTTL,
		totpTokenTTL:    c.TOTPTokenTTL,
	}, nil
}

// customClaims extends jwt.StandardClaims with our extra fields.
type customClaims struct {
	jwt.RegisteredClaims
	SessionGUID string `json:"sid"`
	Scope       string `json:"scope"`
}

func (p *jwtProvider) IssueAccessToken(userGUID, sessionGUID string) (string, error) {
	return p.issueToken(userGUID, sessionGUID, ports.ScopeAccess, p.accessTokenTTL)
}

func (p *jwtProvider) IssueRefreshToken(userGUID, sessionGUID string) (string, error) {
	return p.issueToken(userGUID, sessionGUID, ports.ScopeRefresh, p.refreshTokenTTL)
}

func (p *jwtProvider) IssueTOTPToken(userGUID string) (string, error) {
	// TOTP tokens are scoped and session‑less.
	return p.issueToken(userGUID, "", ports.ScopeTOTP, p.totpTokenTTL)
}

func (p *jwtProvider) issueToken(userGUID, sessionGUID, scope string, ttl time.Duration) (string, error) {
	now := time.Now()
	tokenID := uuid.New().String()

	claims := customClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID,
			Subject:   userGUID,
			Issuer:    p.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
		SessionGUID: sessionGUID,
		Scope:       scope,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(p.secret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signed, nil
}

func (p *jwtProvider) ValidateAccessToken(tokenString string) (*ports.TokenClaims, error) {
	return p.validateToken(tokenString, ports.ScopeAccess)
}

func (p *jwtProvider) ValidateRefreshToken(tokenString string) (*ports.TokenClaims, error) {
	return p.validateToken(tokenString, ports.ScopeRefresh)
}

func (p *jwtProvider) ValidateTOTPToken(tokenString string) (*ports.TokenClaims, error) {
	return p.validateToken(tokenString, ports.ScopeTOTP)
}

// validateToken parses the JWT, checks its validity and scope, and returns the claims.
func (p *jwtProvider) validateToken(tokenString string, expectedScope string) (*ports.TokenClaims, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("%w: token string is empty", ErrInvalidToken)
	}

	parser := jwt.NewParser(
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithIssuer(p.issuer),
		jwt.WithLeeway(0),
	)

	var claims customClaims
	token, err := parser.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%w: %v", ErrUnexpectedSigningMethod, t.Header["alg"])
		}
		return p.secret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("%w: %v", ErrExpiredToken, err)
		}
		return nil, fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	if token == nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	if claims.Scope != expectedScope {
		return nil, fmt.Errorf("%w: got %s, want %s", ErrUnexpectedScope, claims.Scope, expectedScope)
	}

	tc := &ports.TokenClaims{
		TokenGUID:   claims.ID,
		SessionGUID: claims.SessionGUID,
		UserGUID:    claims.Subject,
		Scope:       claims.Scope,
		IssuedAt:    claims.IssuedAt.Time,
		ExpiresAt:   claims.ExpiresAt.Time,
	}

	if tc.TokenGUID == "" || tc.UserGUID == "" || tc.Scope == "" {
		return nil, fmt.Errorf("%w: missing mandatory fields", ErrMissingClaims)
	}

	return tc, nil
}
