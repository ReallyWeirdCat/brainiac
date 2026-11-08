-- name: CreateMessage :one
INSERT INTO message (
    guid,
    chat_member_guid,
    chat_guid,
    resource_url,
    content,
    urgency,
    meta,
    seen_by,
    edited_at,
    pinned_at,
    expire_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: UpdateMessage :one
UPDATE message
SET
    chat_member_guid = $2,
    chat_guid = $3,
    resource_url = $4,
    content = $5,
    urgency = $6,
    meta = $7,
    seen_by = $8,
    edited_at = $9,
    pinned_at = $10,
    expire_at = $11
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveMessage :one
INSERT INTO message (
    guid,
    chat_member_guid,
    chat_guid,
    resource_url,
    content,
    urgency,
    meta,
    seen_by,
    edited_at,
    pinned_at,
    expire_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
ON CONFLICT (guid) DO UPDATE
SET
    chat_member_guid = EXCLUDED.chat_member_guid,
    chat_guid = EXCLUDED.chat_guid,
    resource_url = EXCLUDED.resource_url,
    content = EXCLUDED.content,
    urgency = EXCLUDED.urgency,
    meta = EXCLUDED.meta,
    seen_by = EXCLUDED.seen_by,
    edited_at = EXCLUDED.edited_at,
    pinned_at = EXCLUDED.pinned_at,
    expire_at = EXCLUDED.expire_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteMessage :exec
UPDATE message
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetMessage :one
SELECT *
FROM message
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
