-- name: CreateAppUserCredential :one
INSERT INTO app_user_credential (app_user_guid, email, password_hash, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAppUserCredentialByAppUserGUID :one
SELECT * FROM app_user_credential
WHERE app_user_guid = $1 
    AND deleted_at IS NULL
LIMIT 1;

-- name: GetAppUserCredentialByEmail :one
SELECT * FROM app_user_credential
WHERE email = $1 
    AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateAppUserCredentialPassword :one
UPDATE app_user_credential
SET password_hash = $2
WHERE app_user_guid = $1 
    AND deleted_at IS NULL
RETURNING *;

-- name: UpdateAppUserCredentialEmail :one
UPDATE app_user_credential
SET email = $2
WHERE app_user_guid = $1 
    AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteAppUserCredential :exec
UPDATE app_user_credential
SET deleted_at = $2
WHERE app_user_guid = $1 
    AND deleted_at IS NULL;

-- name: HardDeleteAppUserCredential :exec
DELETE FROM app_user_credential
WHERE app_user_guid = $1;
