-- name: CreateCourseSubjectItem :one
INSERT INTO course_subject_item (
    guid,
    course_subject_guid,
    item_guid,
    amount
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateCourseSubjectItem :one
UPDATE course_subject_item
SET
    course_subject_guid = $2,
    item_guid = $3,
    amount = $4
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveCourseSubjectItem :one
INSERT INTO course_subject_item (
    guid,
    course_subject_guid,
    item_guid,
    amount
) VALUES (
    $1, $2, $3, $4
)
ON CONFLICT (guid) DO UPDATE
SET
    course_subject_guid = EXCLUDED.course_subject_guid,
    item_guid = EXCLUDED.item_guid,
    amount = EXCLUDED.amount
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteCourseSubjectItem :exec
UPDATE course_subject_item
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetCourseSubjectItem :one
SELECT *
FROM course_subject_item
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
