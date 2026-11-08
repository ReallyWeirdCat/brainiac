-- name: CreateCourseSubjectAssessmentPass :one
INSERT INTO course_subject_assessment_pass (
    guid,
    course_subject_assessment_guid,
    app_user_guid,
    enable_text_obfuscation,
    is_active,
    attempts_left,
    expire_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateCourseSubjectAssessmentPass :one
UPDATE course_subject_assessment_pass
SET
    course_subject_assessment_guid = $2,
    app_user_guid = $3,
    enable_text_obfuscation = $4,
    is_active = $5,
    attempts_left = $6,
    expire_at = $7
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveCourseSubjectAssessmentPass :one
INSERT INTO course_subject_assessment_pass (
    guid,
    course_subject_assessment_guid,
    app_user_guid,
    enable_text_obfuscation,
    is_active,
    attempts_left,
    expire_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
ON CONFLICT (guid) DO UPDATE
SET
    course_subject_assessment_guid = EXCLUDED.course_subject_assessment_guid,
    app_user_guid = EXCLUDED.app_user_guid,
    enable_text_obfuscation = EXCLUDED.enable_text_obfuscation,
    is_active = EXCLUDED.is_active,
    attempts_left = EXCLUDED.attempts_left,
    expire_at = EXCLUDED.expire_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteCourseSubjectAssessmentPass :exec
UPDATE course_subject_assessment_pass
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetCourseSubjectAssessmentPass :one
SELECT *
FROM course_subject_assessment_pass
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
