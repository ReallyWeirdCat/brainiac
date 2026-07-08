-- name: CreateAppUserCredential :one
INSERT INTO app_user_credential (
    app_user_guid,
    email,
    password_hash
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateAppUserCredential :one
UPDATE app_user_credential
SET
    email = $2,
    password_hash = $3
WHERE app_user_guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveAppUserCredential :one
INSERT INTO app_user_credential (
    app_user_guid,
    email,
    password_hash
) VALUES (
    $1, $2, $3
)
ON CONFLICT (app_user_guid) DO UPDATE
SET
    email = EXCLUDED.email,
    password_hash = EXCLUDED.password_hash
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteAppUserCredential :exec
UPDATE app_user_credential
SET deleted_at = now()
WHERE app_user_guid = $1 AND deleted_at IS NULL;

-- name: GetAppUserCredential :one
SELECT *
FROM app_user_credential
WHERE app_user_guid = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetAllAppUserCredentials :many
SELECT *
FROM app_user_credential
WHERE deleted_at IS NULL;

-- name: CountAppUserCredentials :one
SELECT COUNT(*)
FROM app_user_credential
WHERE deleted_at IS NULL;

-- name: ExistsAppUserCredential :one
SELECT EXISTS (
    SELECT 1
    FROM app_user_credential
    WHERE app_user_guid = $1 AND deleted_at IS NULL
);

-- name: IsDeletedAppUserCredential :one
SELECT coalesce(deleted_at IS NOT NULL, false)::boolean
FROM app_user_credential
WHERE app_user_guid = $1;

-- name: CreateAppUserCredentialBatch :batchmany
INSERT INTO app_user_credential (
    app_user_guid,
    email,
    password_hash
) VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateAppUserCredentialBatch :batchmany
UPDATE app_user_credential
SET
    email = $2,
    password_hash = $3
WHERE app_user_guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveAppUserCredentialBatch :batchmany
INSERT INTO app_user_credential (
    app_user_guid,
    email,
    password_hash
) VALUES ($1, $2, $3)
ON CONFLICT (app_user_guid) DO UPDATE
SET
    email = EXCLUDED.email,
    password_hash = EXCLUDED.password_hash
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteAppUserCredentialBatch :batchexec
UPDATE app_user_credential
SET deleted_at = now()
WHERE app_user_guid = $1 AND deleted_at IS NULL;

-- name: GetAppUserCredentialBatch :batchmany
SELECT * FROM app_user_credential
WHERE app_user_guid = $1 AND deleted_at IS NULL;

-- name: ExistsAppUserCredentialBatch :batchmany
SELECT app_user_guid FROM app_user_credential
WHERE app_user_guid = $1 AND deleted_at IS NULL;

-- name: GetAppUserCredentialByEmail :one
SELECT *
FROM app_user_credential
WHERE email = $1
  AND deleted_at IS NULL
LIMIT 1;

-- name: ExistsAppUserCredentialByEmail :one
SELECT EXISTS (
    SELECT 1
    FROM app_user_credential
    WHERE email = $1 AND deleted_at IS NULL
);
