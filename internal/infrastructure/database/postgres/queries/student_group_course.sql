-- name: CreateStudentGroupCourse :one
INSERT INTO student_group_course (
    guid,
    student_group_guid,
    course_guid,
    full_access,
    close_submissions
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateStudentGroupCourse :one
UPDATE student_group_course
SET
    student_group_guid = $2,
    course_guid = $3,
    full_access = $4,
    close_submissions = $5
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveStudentGroupCourse :one
INSERT INTO student_group_course (
    guid,
    student_group_guid,
    course_guid,
    full_access,
    close_submissions
) VALUES (
    $1, $2, $3, $4, $5
)
ON CONFLICT (guid) DO UPDATE
SET
    student_group_guid = EXCLUDED.student_group_guid,
    course_guid = EXCLUDED.course_guid,
    full_access = EXCLUDED.full_access,
    close_submissions = EXCLUDED.close_submissions
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteStudentGroupCourse :exec
UPDATE student_group_course
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetStudentGroupCourse :one
SELECT *
FROM student_group_course
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
