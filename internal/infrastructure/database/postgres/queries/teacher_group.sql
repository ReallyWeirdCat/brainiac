-- name: CreateTeacherGroup :one
INSERT INTO teacher_group (
    guid,
    teacher_guid,
    student_group_guid,
    active_since,
    manage_attendance,
    manage_results,
    manage_students,
    manage_achievements,
    teacher_until
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateTeacherGroup :one
UPDATE teacher_group
SET
    teacher_guid = $2,
    student_group_guid = $3,
    active_since = $4,
    manage_attendance = $5,
    manage_results = $6,
    manage_students = $7,
    manage_achievements = $8,
    teacher_until = $9
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveTeacherGroup :one
INSERT INTO teacher_group (
    guid,
    teacher_guid,
    student_group_guid,
    active_since,
    manage_attendance,
    manage_results,
    manage_students,
    manage_achievements,
    teacher_until
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
ON CONFLICT (guid) DO UPDATE
SET
    teacher_guid = EXCLUDED.teacher_guid,
    student_group_guid = EXCLUDED.student_group_guid,
    active_since = EXCLUDED.active_since,
    manage_attendance = EXCLUDED.manage_attendance,
    manage_results = EXCLUDED.manage_results,
    manage_students = EXCLUDED.manage_students,
    manage_achievements = EXCLUDED.manage_achievements,
    teacher_until = EXCLUDED.teacher_until
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteTeacherGroup :exec
UPDATE teacher_group
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetTeacherGroup :one
SELECT *
FROM teacher_group
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
