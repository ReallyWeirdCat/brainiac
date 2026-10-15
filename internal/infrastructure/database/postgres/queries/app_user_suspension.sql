-- name: CreateAppUserSuspension :one
INSERT INTO app_user_suspension (
    guid,
    app_user_guid,
    reason,
    expire_at,
    expired_at
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateAppUserSuspension :one
UPDATE app_user_suspension
SET
    app_user_guid = $2,
    reason = $3,
    expire_at = $4,
    expired_at = $5
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveAppUserSuspension :one
INSERT INTO app_user_suspension (
    guid,
    app_user_guid,
    reason,
    expire_at,
    expired_at
) VALUES (
    $1, $2, $3, $4, $5
)
ON CONFLICT (guid) DO UPDATE
SET
    app_user_guid = EXCLUDED.app_user_guid,
    reason = EXCLUDED.reason,
    expire_at = EXCLUDED.expire_at,
    expired_at = EXCLUDED.expired_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteAppUserSuspension :exec
UPDATE app_user_suspension
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetAppUserSuspension :one
SELECT *
FROM app_user_suspension
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
