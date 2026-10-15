-- name: CreateChat :one
INSERT INTO chat (
    guid,
    title,
    topic,
    avatar_url,
    students_can_text,
    students_see_members,
    preserve_messages,
    meta
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateChat :one
UPDATE chat
SET
    title = $2,
    topic = $3,
    avatar_url = $4,
    students_can_text = $5,
    students_see_members = $6,
    preserve_messages = $7,
    meta = $8
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveChat :one
INSERT INTO chat (
    guid,
    title,
    topic,
    avatar_url,
    students_can_text,
    students_see_members,
    preserve_messages,
    meta
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
ON CONFLICT (guid) DO UPDATE
SET
    title = EXCLUDED.title,
    topic = EXCLUDED.topic,
    avatar_url = EXCLUDED.avatar_url,
    students_can_text = EXCLUDED.students_can_text,
    students_see_members = EXCLUDED.students_see_members,
    preserve_messages = EXCLUDED.preserve_messages,
    meta = EXCLUDED.meta
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteChat :exec
UPDATE chat
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetChat :one
SELECT *
FROM chat
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
