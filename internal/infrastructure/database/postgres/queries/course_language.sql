-- name: CreateCourseLanguage :one
INSERT INTO course_language (
    guid,
    course_guid,
    language_guid,
    calculated_support_percentage,
    is_beta,
    is_new
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateCourseLanguage :one
UPDATE course_language
SET
    course_guid = $2,
    language_guid = $3,
    calculated_support_percentage = $4,
    is_beta = $5,
    is_new = $6
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveCourseLanguage :one
INSERT INTO course_language (
    guid,
    course_guid,
    language_guid,
    calculated_support_percentage,
    is_beta,
    is_new
) VALUES (
    $1, $2, $3, $4, $5, $6
)
ON CONFLICT (guid) DO UPDATE
SET
    course_guid = EXCLUDED.course_guid,
    language_guid = EXCLUDED.language_guid,
    calculated_support_percentage = EXCLUDED.calculated_support_percentage,
    is_beta = EXCLUDED.is_beta,
    is_new = EXCLUDED.is_new
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteCourseLanguage :exec
UPDATE course_language
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetCourseLanguage :one
SELECT *
FROM course_language
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
