-- name: CreateStudentCourseSubjectSubmissionAttempt :one
INSERT INTO student_course_subject_submission_attempt (
    guid,
    app_user_guid,
    course_subject_doc_guid,
    submission_status,
    files,
    meta,
    student_comment,
    teacher_comment
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateStudentCourseSubjectSubmissionAttempt :one
UPDATE student_course_subject_submission_attempt
SET
    app_user_guid = $2,
    course_subject_doc_guid = $3,
    submission_status = $4,
    files = $5,
    meta = $6,
    student_comment = $7,
    teacher_comment = $8
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveStudentCourseSubjectSubmissionAttempt :one
INSERT INTO student_course_subject_submission_attempt (
    guid,
    app_user_guid,
    course_subject_doc_guid,
    submission_status,
    files,
    meta,
    student_comment,
    teacher_comment
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
ON CONFLICT (guid) DO UPDATE
SET
    app_user_guid = EXCLUDED.app_user_guid,
    course_subject_doc_guid = EXCLUDED.course_subject_doc_guid,
    submission_status = EXCLUDED.submission_status,
    files = EXCLUDED.files,
    meta = EXCLUDED.meta,
    student_comment = EXCLUDED.student_comment,
    teacher_comment = EXCLUDED.teacher_comment
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteStudentCourseSubjectSubmissionAttempt :exec
UPDATE student_course_subject_submission_attempt
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetStudentCourseSubjectSubmissionAttempt :one
SELECT *
FROM student_course_subject_submission_attempt
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
