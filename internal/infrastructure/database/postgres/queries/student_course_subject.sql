-- name: CreateStudentCourseSubject :one
INSERT INTO student_course_subject (
    guid,
    app_user_guid,
    course_subject_guid,
    completed_at,
    subject_status,
    is_favorite,
    meta
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateStudentCourseSubject :one
UPDATE student_course_subject
SET
    app_user_guid = $2,
    course_subject_guid = $3,
    completed_at = $4,
    subject_status = $5,
    is_favorite = $6,
    meta = $7
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveStudentCourseSubject :one
INSERT INTO student_course_subject (
    guid,
    app_user_guid,
    course_subject_guid,
    completed_at,
    subject_status,
    is_favorite,
    meta
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
ON CONFLICT (guid) DO UPDATE
SET
    app_user_guid = EXCLUDED.app_user_guid,
    course_subject_guid = EXCLUDED.course_subject_guid,
    completed_at = EXCLUDED.completed_at,
    subject_status = EXCLUDED.subject_status,
    is_favorite = EXCLUDED.is_favorite,
    meta = EXCLUDED.meta
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteStudentCourseSubject :exec
UPDATE student_course_subject
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetStudentCourseSubject :one
SELECT *
FROM student_course_subject
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
