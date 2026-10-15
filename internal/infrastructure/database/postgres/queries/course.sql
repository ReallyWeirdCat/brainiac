-- name: CreateCourse :one
INSERT INTO course (
    guid,
    title_i18n,
    description_i18n,
    style,
    meta,
    published_at
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateCourse :one
UPDATE course
SET
    title_i18n = $2,
    description_i18n = $3,
    style = $4,
    meta = $5,
    published_at = $6
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveCourse :one
INSERT INTO course (
    guid,
    title_i18n,
    description_i18n,
    style,
    meta,
    published_at
) VALUES (
    $1, $2, $3, $4, $5, $6
)
ON CONFLICT (guid) DO UPDATE
SET
    title_i18n = EXCLUDED.title_i18n,
    description_i18n = EXCLUDED.description_i18n,
    style = EXCLUDED.style,
    meta = EXCLUDED.meta,
    published_at = EXCLUDED.published_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteCourse :exec
UPDATE course
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetCourse :one
SELECT *
FROM course
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
