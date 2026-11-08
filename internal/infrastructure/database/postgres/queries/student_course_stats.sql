-- name: CreateStudentCourseStats :one
INSERT INTO student_course_stats (
    guid,
    app_user_guid,
    course_guid,
    experience,
    level,
    meta
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateStudentCourseStats :one
UPDATE student_course_stats
SET
    app_user_guid = $2,
    course_guid = $3,
    experience = $4,
    level = $5,
    meta = $6
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveStudentCourseStats :one
INSERT INTO student_course_stats (
    guid,
    app_user_guid,
    course_guid,
    experience,
    level,
    meta
) VALUES (
    $1, $2, $3, $4, $5, $6
)
ON CONFLICT (guid) DO UPDATE
SET
    app_user_guid = EXCLUDED.app_user_guid,
    course_guid = EXCLUDED.course_guid,
    experience = EXCLUDED.experience,
    level = EXCLUDED.level,
    meta = EXCLUDED.meta
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteStudentCourseStats :exec
UPDATE student_course_stats
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetStudentCourseStats :one
SELECT *
FROM student_course_stats
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
