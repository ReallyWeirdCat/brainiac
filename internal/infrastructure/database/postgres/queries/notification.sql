-- name: CreateNotification :one
INSERT INTO notification (
    guid,
    app_user_guid,
    title_i18n,
    content_i18n,
    resource_url,
    urgency,
    meta,
    seen_at,
    expire_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateNotification :one
UPDATE notification
SET
    app_user_guid = $2,
    title_i18n = $3,
    content_i18n = $4,
    resource_url = $5,
    urgency = $6,
    meta = $7,
    seen_at = $8,
    expire_at = $9
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveNotification :one
INSERT INTO notification (
    guid,
    app_user_guid,
    title_i18n,
    content_i18n,
    resource_url,
    urgency,
    meta,
    seen_at,
    expire_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
ON CONFLICT (guid) DO UPDATE
SET
    app_user_guid = EXCLUDED.app_user_guid,
    title_i18n = EXCLUDED.title_i18n,
    content_i18n = EXCLUDED.content_i18n,
    resource_url = EXCLUDED.resource_url,
    urgency = EXCLUDED.urgency,
    meta = EXCLUDED.meta,
    seen_at = EXCLUDED.seen_at,
    expire_at = EXCLUDED.expire_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteNotification :exec
UPDATE notification
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetNotification :one
SELECT *
FROM notification
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
