-- name: CreateStudentCourseSubjectAssessmentAttempt :one
INSERT INTO student_course_subject_assessment_attempt (
    guid,
    app_user_guid,
    course_subject_assessment_guid,
    attempt_status,
    score,
    is_success,
    answers
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateStudentCourseSubjectAssessmentAttempt :one
UPDATE student_course_subject_assessment_attempt
SET
    app_user_guid = $2,
    course_subject_assessment_guid = $3,
    attempt_status = $4,
    score = $5,
    is_success = $6,
    answers = $7
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveStudentCourseSubjectAssessmentAttempt :one
INSERT INTO student_course_subject_assessment_attempt (
    guid,
    app_user_guid,
    course_subject_assessment_guid,
    attempt_status,
    score,
    is_success,
    answers
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
ON CONFLICT (guid) DO UPDATE
SET
    app_user_guid = EXCLUDED.app_user_guid,
    course_subject_assessment_guid = EXCLUDED.course_subject_assessment_guid,
    attempt_status = EXCLUDED.attempt_status,
    score = EXCLUDED.score,
    is_success = EXCLUDED.is_success,
    answers = EXCLUDED.answers
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteStudentCourseSubjectAssessmentAttempt :exec
UPDATE student_course_subject_assessment_attempt
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetStudentCourseSubjectAssessmentAttempt :one
SELECT *
FROM student_course_subject_assessment_attempt
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
