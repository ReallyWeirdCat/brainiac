-- name: CreateCourseSubject :one
INSERT INTO course_subject (
    guid,
    course_guid,
    parent_guid,
    title_i18n,
    description_i18n,
    reward_experience,
    completion_condition,
    hide_children,
    style,
    meta,
    published_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: UpdateCourseSubject :one
UPDATE course_subject
SET
    course_guid = $2,
    parent_guid = $3,
    title_i18n = $4,
    description_i18n = $5,
    reward_experience = $6,
    completion_condition = $7,
    hide_children = $8,
    style = $9,
    meta = $10,
    published_at = $11
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveCourseSubject :one
INSERT INTO course_subject (
    guid,
    course_guid,
    parent_guid,
    title_i18n,
    description_i18n,
    reward_experience,
    completion_condition,
    hide_children,
    style,
    meta,
    published_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
ON CONFLICT (guid) DO UPDATE
SET
    course_guid = EXCLUDED.course_guid,
    parent_guid = EXCLUDED.parent_guid,
    title_i18n = EXCLUDED.title_i18n,
    description_i18n = EXCLUDED.description_i18n,
    reward_experience = EXCLUDED.reward_experience,
    completion_condition = EXCLUDED.completion_condition,
    hide_children = EXCLUDED.hide_children,
    style = EXCLUDED.style,
    meta = EXCLUDED.meta,
    published_at = EXCLUDED.published_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteCourseSubject :exec
UPDATE course_subject
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetCourseSubject :one
SELECT *
FROM course_subject
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
