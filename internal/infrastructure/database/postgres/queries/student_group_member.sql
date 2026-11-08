-- name: CreateStudentGroupMember :one
INSERT INTO student_group_member (
    guid,
    app_user_guid,
    student_group_guid,
    student_role,
    notes,
    team,
    team_responsibilities
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateStudentGroupMember :one
UPDATE student_group_member
SET
    app_user_guid = $2,
    student_group_guid = $3,
    student_role = $4,
    notes = $5,
    team = $6,
    team_responsibilities = $7
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveStudentGroupMember :one
INSERT INTO student_group_member (
    guid,
    app_user_guid,
    student_group_guid,
    student_role,
    notes,
    team,
    team_responsibilities
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
ON CONFLICT (guid) DO UPDATE
SET
    app_user_guid = EXCLUDED.app_user_guid,
    student_group_guid = EXCLUDED.student_group_guid,
    student_role = EXCLUDED.student_role,
    notes = EXCLUDED.notes,
    team = EXCLUDED.team,
    team_responsibilities = EXCLUDED.team_responsibilities
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteStudentGroupMember :exec
UPDATE student_group_member
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetStudentGroupMember :one
SELECT *
FROM student_group_member
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
