-- name: CreateTOTPSecret :one
INSERT INTO totp_secret (
    app_user_guid,
    secret_base32,
    last_used_at
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateTOTPSecret :one
UPDATE totp_secret
SET
    secret_base32 = $2,
    last_used_at = $3
WHERE app_user_guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveTOTPSecret :one
INSERT INTO totp_secret (
    app_user_guid,
    secret_base32,
    last_used_at
) VALUES (
    $1, $2, $3
)
ON CONFLICT (app_user_guid) DO UPDATE
SET
    secret_base32 = EXCLUDED.secret_base32,
    last_used_at = EXCLUDED.last_used_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteTOTPSecret :exec
UPDATE totp_secret
SET deleted_at = now()
WHERE app_user_guid = $1 AND deleted_at IS NULL;

-- name: GetTOTPSecret :one
SELECT *
FROM totp_secret
WHERE app_user_guid = $1 AND deleted_at IS NULL
LIMIT 1;
