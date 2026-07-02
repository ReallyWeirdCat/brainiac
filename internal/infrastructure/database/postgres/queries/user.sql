-- name: CreateAppUser :one
INSERT INTO app_user (
    guid,
    username,
    activated_at
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateAppUser :one
UPDATE app_user
SET
    username = $2,
    activated_at = $3
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveAppUser :one
INSERT INTO app_user (
    guid,
    username,
    activated_at
) VALUES (
    $1, $2, $3
)
ON CONFLICT (guid) DO UPDATE
SET
    username = EXCLUDED.username,
    activated_at = EXCLUDED.activated_at
RETURNING *;

-- name: DeleteAppUser :exec
UPDATE app_user
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetAppUser :one
SELECT *
FROM app_user
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetAllAppUsers :many
SELECT *
FROM app_user
WHERE deleted_at IS NULL;

-- name: CountAppUsers :one
SELECT COUNT(*)
FROM app_user
WHERE deleted_at IS NULL;

-- name: ExistsAppUser :one
SELECT EXISTS (
    SELECT 1
    FROM app_user
    WHERE guid = $1 AND deleted_at IS NULL
);

-- name: IsDeletedAppUser :one
SELECT coalesce(deleted_at IS NOT NULL, false)::boolean
FROM app_user
WHERE guid = $1;

-- name: CreateAppUserBatch :batchmany
INSERT INTO app_user (
    guid,
    username,
    activated_at
) VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateAppUserBatch :batchmany
UPDATE app_user
SET 
    username = $2,
    activated_at = $3
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveAppUserBatch :batchmany
INSERT INTO app_user (
    guid,
    username,
    activated_at
) VALUES (
    $1, $2, $3
)
ON CONFLICT (guid) DO UPDATE
SET
    username = EXCLUDED.username,
    activated_at = EXCLUDED.activated_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteAppUserBatch :batchexec
UPDATE app_user
SET 
    deleted_at = true
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetAppUserBatch :batchmany
SELECT * FROM app_user
WHERE guid = $1 AND deleted_at IS NULL;

-- name: ExistsAppUserBatch :batchmany
SELECT guid FROM app_user
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetAppUserByUsername :one
SELECT *
FROM app_user
WHERE username = $1
  AND deleted_at IS NULL;

-- name: GetAppUserByEmail :one
SELECT au.*
FROM app_user au
INNER JOIN app_user_credential auc ON au.guid = auc.app_user_guid
WHERE auc.email = $1
  AND auc.deleted_at IS NULL
  AND au.deleted_at IS NULL;
