-- name: CreateTeacher :one
INSERT INTO teacher (
    app_user_guid,
    active_since,
    manage_students,
    manage_attendance,
    manage_results,
    manage_courses,
    is_admin,
    contact_info
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateTeacher :one
UPDATE teacher
SET
    active_since = $2,
    manage_students = $3,
    manage_attendance = $4,
    manage_results = $5,
    manage_courses = $6,
    is_admin = $7,
    contact_info = $8
WHERE app_user_guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveTeacher :one
INSERT INTO teacher (
    app_user_guid,
    active_since,
    manage_students,
    manage_attendance,
    manage_results,
    manage_courses,
    is_admin,
    contact_info
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
ON CONFLICT (app_user_guid) DO UPDATE
SET
    active_since = EXCLUDED.active_since,
    manage_students = EXCLUDED.manage_students,
    manage_attendance = EXCLUDED.manage_attendance,
    manage_results = EXCLUDED.manage_results,
    manage_courses = EXCLUDED.manage_courses,
    is_admin = EXCLUDED.is_admin,
    contact_info = EXCLUDED.contact_info
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteTeacher :exec
UPDATE teacher
SET deleted_at = now()
WHERE app_user_guid = $1 AND deleted_at IS NULL;

-- name: GetTeacher :one
SELECT *
FROM teacher
WHERE app_user_guid = $1 AND deleted_at IS NULL
LIMIT 1;
