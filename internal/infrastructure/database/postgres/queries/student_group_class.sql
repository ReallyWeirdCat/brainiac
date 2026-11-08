-- name: CreateStudentGroupClass :one
INSERT INTO student_group_class (
    guid,
    student_group_guid,
    title,
    start_time,
    end_time,
    objectives,
    attestation_assessments_enabled,
    attendance_closed_at,
    marks_approved_at,
    totp_secret,
    room,
    class_type
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING *;

-- name: UpdateStudentGroupClass :one
UPDATE student_group_class
SET
    student_group_guid = $2,
    title = $3,
    start_time = $4,
    end_time = $5,
    objectives = $6,
    attestation_assessments_enabled = $7,
    attendance_closed_at = $8,
    marks_approved_at = $9,
    totp_secret = $10,
    room = $11,
    class_type = $12
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveStudentGroupClass :one
INSERT INTO student_group_class (
    guid,
    student_group_guid,
    title,
    start_time,
    end_time,
    objectives,
    attestation_assessments_enabled,
    attendance_closed_at,
    marks_approved_at,
    totp_secret,
    room,
    class_type
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
ON CONFLICT (guid) DO UPDATE
SET
    student_group_guid = EXCLUDED.student_group_guid,
    title = EXCLUDED.title,
    start_time = EXCLUDED.start_time,
    end_time = EXCLUDED.end_time,
    objectives = EXCLUDED.objectives,
    attestation_assessments_enabled = EXCLUDED.attestation_assessments_enabled,
    attendance_closed_at = EXCLUDED.attendance_closed_at,
    marks_approved_at = EXCLUDED.marks_approved_at,
    totp_secret = EXCLUDED.totp_secret,
    room = EXCLUDED.room,
    class_type = EXCLUDED.class_type
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteStudentGroupClass :exec
UPDATE student_group_class
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetStudentGroupClass :one
SELECT *
FROM student_group_class
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
