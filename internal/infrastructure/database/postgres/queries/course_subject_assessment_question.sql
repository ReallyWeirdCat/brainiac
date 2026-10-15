-- name: CreateCourseSubjectAssessmentQuestion :one
INSERT INTO course_subject_assessment_question (
    guid,
    course_subject_assessment_guid,
    title_i18n,
    description_i18n,
    question_type,
    attachment_url,
    options,
    correct_options,
    max_correct_options,
    is_multiple_choice,
    max_options,
    max_answer_time,
    use_text_answer,
    correct_text_answer,
    example_url,
    published_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
)
RETURNING *;

-- name: UpdateCourseSubjectAssessmentQuestion :one
UPDATE course_subject_assessment_question
SET
    course_subject_assessment_guid = $2,
    title_i18n = $3,
    description_i18n = $4,
    question_type = $5,
    attachment_url = $6,
    options = $7,
    correct_options = $8,
    max_correct_options = $9,
    is_multiple_choice = $10,
    max_options = $11,
    max_answer_time = $12,
    use_text_answer = $13,
    correct_text_answer = $14,
    example_url = $15,
    published_at = $16
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveCourseSubjectAssessmentQuestion :one
INSERT INTO course_subject_assessment_question (
    guid,
    course_subject_assessment_guid,
    title_i18n,
    description_i18n,
    question_type,
    attachment_url,
    options,
    correct_options,
    max_correct_options,
    is_multiple_choice,
    max_options,
    max_answer_time,
    use_text_answer,
    correct_text_answer,
    example_url,
    published_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
)
ON CONFLICT (guid) DO UPDATE
SET
    course_subject_assessment_guid = EXCLUDED.course_subject_assessment_guid,
    title_i18n = EXCLUDED.title_i18n,
    description_i18n = EXCLUDED.description_i18n,
    question_type = EXCLUDED.question_type,
    attachment_url = EXCLUDED.attachment_url,
    options = EXCLUDED.options,
    correct_options = EXCLUDED.correct_options,
    max_correct_options = EXCLUDED.max_correct_options,
    is_multiple_choice = EXCLUDED.is_multiple_choice,
    max_options = EXCLUDED.max_options,
    max_answer_time = EXCLUDED.max_answer_time,
    use_text_answer = EXCLUDED.use_text_answer,
    correct_text_answer = EXCLUDED.correct_text_answer,
    example_url = EXCLUDED.example_url,
    published_at = EXCLUDED.published_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteCourseSubjectAssessmentQuestion :exec
UPDATE course_subject_assessment_question
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetCourseSubjectAssessmentQuestion :one
SELECT *
FROM course_subject_assessment_question
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
