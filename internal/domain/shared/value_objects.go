package shared

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

type Email struct {
	value string
}

func NewEmail(email string) (Email, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return Email{}, errors.New("email cannot be empty")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return Email{}, errors.New("invalid email format")
	}

	if strings.Contains(email, "..") {
		return Email{}, errors.New("email cannot contain consecutive dots")
	}

	return Email{value: email}, nil
}

func (e Email) String() string {
	return e.value
}

func (e Email) Equals(other Email) bool {
	return strings.EqualFold(e.value, other.value)
}

func (e Email) Valid() bool {
	_, err := NewEmail(e.value)
	return err == nil
}

func (e Email) Domain() string {
	parts := strings.Split(e.value, "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

func (e Email) LocalPart() string {
	parts := strings.Split(e.value, "@")
	if len(parts) == 2 {
		return parts[0]
	}
	return ""
}

type URL struct {
	value  string
	scheme string
	host   string
	path   string
}

func NewURL(rawURL string) (URL, error) {
	rawURL = strings.TrimSpace(rawURL)
	if rawURL == "" {
		return URL{}, errors.New("URL cannot be empty")
	}

	if !strings.Contains(rawURL, "://") {
		rawURL = "https://" + rawURL
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return URL{}, errors.New("invalid URL format")
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return URL{}, errors.New("URL must use HTTP or HTTPS scheme")
	}

	if parsedURL.Host == "" {
		return URL{}, errors.New("URL must have a host")
	}

	host := parsedURL.Hostname()
	if host == "" {
		return URL{}, errors.New("invalid domain in URL")
	}

	domainRegex := regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)
	if !domainRegex.MatchString(host) {
		return URL{}, errors.New("invalid domain format")
	}

	return URL{
		value:  parsedURL.String(),
		scheme: parsedURL.Scheme,
		host:   parsedURL.Host,
		path:   parsedURL.Path,
	}, nil
}

func (u URL) String() string {
	return u.value
}

func (u URL) Equals(other URL) bool {
	return u.value == other.value
}

func (u URL) Valid() bool {
	_, err := NewURL(u.value)
	return err == nil
}

func (u URL) Scheme() string {
	return u.scheme
}

func (u URL) Host() string {
	return u.host
}

func (u URL) Path() string {
	return u.path
}

func (u URL) Domain() string {
	host := u.Hostname()
	if idx := strings.Index(host, ":"); idx != -1 {
		host = host[:idx]
	}
	return host
}

func (u URL) Hostname() string {
	host := u.host
	if idx := strings.Index(host, ":"); idx != -1 {
		host = host[:idx]
	}
	return host
}

type Username struct {
	value string
}

func NewUsername(username string) (Username, error) {
	username = strings.TrimSpace(username)

	if len(username) == 0 {
		return Username{}, errors.New("username cannot be empty")
	}
	if len(username) > 18 || len(username) < 3 {
		return Username{}, errors.New("username must be between 3 and 18 characters in length")
	}

	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !usernameRegex.MatchString(username) {
		return Username{}, errors.New("username can only contain Latin letters, numbers, and underscores")
	}

	if strings.HasPrefix(username, "_") || strings.HasSuffix(username, "_") {
		return Username{}, errors.New("username cannot start or end with underscore")
	}

	if strings.Contains(username, "__") {
		return Username{}, errors.New("username cannot contain consecutive underscores")
	}

	return Username{value: username}, nil
}

func (u Username) String() string {
	return u.value
}

func (u Username) Equals(other Username) bool {
	return u.value == other.value
}

func (u Username) Valid() bool {
	_, err := NewUsername(u.value)
	return err == nil
}

func (u Username) Length() int {
	return len(u.value)
}
