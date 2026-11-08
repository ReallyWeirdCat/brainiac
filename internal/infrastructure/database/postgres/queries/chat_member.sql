-- name: CreateChatMember :one
INSERT INTO chat_member (
    guid,
    app_user_guid,
    chat_guid,
    chat_role,
    name,
    custom_role_name,
    left_at,
    seen_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateChatMember :one
UPDATE chat_member
SET
    app_user_guid = $2,
    chat_guid = $3,
    chat_role = $4,
    name = $5,
    custom_role_name = $6,
    left_at = $7,
    seen_at = $8
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveChatMember :one
INSERT INTO chat_member (
    guid,
    app_user_guid,
    chat_guid,
    chat_role,
    name,
    custom_role_name,
    left_at,
    seen_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
ON CONFLICT (guid) DO UPDATE
SET
    app_user_guid = EXCLUDED.app_user_guid,
    chat_guid = EXCLUDED.chat_guid,
    chat_role = EXCLUDED.chat_role,
    name = EXCLUDED.name,
    custom_role_name = EXCLUDED.custom_role_name,
    left_at = EXCLUDED.left_at,
    seen_at = EXCLUDED.seen_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteChatMember :exec
UPDATE chat_member
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetChatMember :one
SELECT *
FROM chat_member
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
