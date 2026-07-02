-- name: CreateAppUser :one
INSERT INTO app_user (
    guid,
    username,
    activated_at,
    created_at,
    deleted_at
) VALUES (
    $1, $2, $3, now(), $4
)
RETURNING *;

-- name: UpdateAppUser :one
UPDATE app_user
SET
    username = $2,
    activated_at = $3,
    deleted_at = $4
WHERE guid = $1
RETURNING *;

-- name: SaveAppUser :one
INSERT INTO app_user (
    guid,
    username,
    activated_at,
    created_at,
    deleted_at
) VALUES (
    $1, $2, $3, now(), $4
)
ON CONFLICT (guid) DO UPDATE
SET
    username = EXCLUDED.username,
    activated_at = EXCLUDED.activated_at,
    deleted_at = EXCLUDED.deleted_at
RETURNING *;

-- name: DeleteAppUser :exec
UPDATE app_user
SET deleted_at = now()
WHERE guid = $1;

-- name: GetAppUser :one
SELECT *
FROM app_user
WHERE guid = $1
LIMIT 1;

-- name: GetAllAppUsers :many
SELECT *
FROM app_user;

-- name: CountAppUsers :one
SELECT COUNT(*)
FROM app_user;

-- name: ExistsAppUser :one
SELECT EXISTS (
    SELECT 1
    FROM app_user
    WHERE guid = $1
);

-- name: IsDeletedAppUser :one
SELECT coalesce(deleted_at IS NOT NULL, false)::boolean
FROM app_user
WHERE guid = $1;

-- name: CreateAppUserBatch :batchmany
INSERT INTO app_user (
    guid,
    username,
    activated_at,
    deleted_at
) VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateAppUserBatch :batchmany
UPDATE app_user
SET 
    username = $2,
    activated_at = $3,
    deleted_at = $4
WHERE guid = $1
RETURNING *;

-- name: SaveAppUserBatch :batchmany
INSERT INTO app_user (
    guid,
    username,
    activated_at,
    deleted_at
) VALUES (
    $1, $2, $3, $4
)
ON CONFLICT (guid) DO UPDATE
SET
    username = EXCLUDED.username,
    activated_at = EXCLUDED.activated_at,
    deleted_at = EXCLUDED.deleted_at
RETURNING *;

-- name: DeleteAppUserBatch :batchexec
UPDATE app_user
SET 
    deleted_at = true
WHERE guid = $1;

-- name: GetAppUserBatch :batchmany
SELECT * FROM app_user
WHERE guid = $1;

-- name: ExistsAppUserBatch :batchmany
SELECT guid FROM app_user
WHERE guid = $1;

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
