-- name: CreateStudentGroupMemberClassAttendance :one
INSERT INTO student_group_member_class_attendance (
    guid,
    student_group_member_guid,
    student_group_class_guid,
    first_seen_at,
    last_seen_at,
    present_at,
    objectives_completion,
    comment,
    grade
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateStudentGroupMemberClassAttendance :one
UPDATE student_group_member_class_attendance
SET
    student_group_member_guid = $2,
    student_group_class_guid = $3,
    first_seen_at = $4,
    last_seen_at = $5,
    present_at = $6,
    objectives_completion = $7,
    comment = $8,
    grade = $9
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveStudentGroupMemberClassAttendance :one
INSERT INTO student_group_member_class_attendance (
    guid,
    student_group_member_guid,
    student_group_class_guid,
    first_seen_at,
    last_seen_at,
    present_at,
    objectives_completion,
    comment,
    grade
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
ON CONFLICT (guid) DO UPDATE
SET
    student_group_member_guid = EXCLUDED.student_group_member_guid,
    student_group_class_guid = EXCLUDED.student_group_class_guid,
    first_seen_at = EXCLUDED.first_seen_at,
    last_seen_at = EXCLUDED.last_seen_at,
    present_at = EXCLUDED.present_at,
    objectives_completion = EXCLUDED.objectives_completion,
    comment = EXCLUDED.comment,
    grade = EXCLUDED.grade
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteStudentGroupMemberClassAttendance :exec
UPDATE student_group_member_class_attendance
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetStudentGroupMemberClassAttendance :one
SELECT *
FROM student_group_member_class_attendance
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
