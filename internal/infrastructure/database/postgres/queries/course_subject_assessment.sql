-- name: CreateCourseSubjectAssessment :one
INSERT INTO course_subject_assessment (
    guid,
    course_subject_guid,
    title_i18n,
    description_i18n,
    is_attestation,
    hide_answers,
    hide_answer_validity,
    hide_score,
    use_timed_questions,
    max_time,
    max_attempts,
    scoring_method,
    required_score,
    short_questions_count,
    normal_questions_count,
    long_questions_count,
    published_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
)
RETURNING *;

-- name: UpdateCourseSubjectAssessment :one
UPDATE course_subject_assessment
SET
    course_subject_guid = $2,
    title_i18n = $3,
    description_i18n = $4,
    is_attestation = $5,
    hide_answers = $6,
    hide_answer_validity = $7,
    hide_score = $8,
    use_timed_questions = $9,
    max_time = $10,
    max_attempts = $11,
    scoring_method = $12,
    required_score = $13,
    short_questions_count = $14,
    normal_questions_count = $15,
    long_questions_count = $16,
    published_at = $17
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveCourseSubjectAssessment :one
INSERT INTO course_subject_assessment (
    guid,
    course_subject_guid,
    title_i18n,
    description_i18n,
    is_attestation,
    hide_answers,
    hide_answer_validity,
    hide_score,
    use_timed_questions,
    max_time,
    max_attempts,
    scoring_method,
    required_score,
    short_questions_count,
    normal_questions_count,
    long_questions_count,
    published_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
)
ON CONFLICT (guid) DO UPDATE
SET
    course_subject_guid = EXCLUDED.course_subject_guid,
    title_i18n = EXCLUDED.title_i18n,
    description_i18n = EXCLUDED.description_i18n,
    is_attestation = EXCLUDED.is_attestation,
    hide_answers = EXCLUDED.hide_answers,
    hide_answer_validity = EXCLUDED.hide_answer_validity,
    hide_score = EXCLUDED.hide_score,
    use_timed_questions = EXCLUDED.use_timed_questions,
    max_time = EXCLUDED.max_time,
    max_attempts = EXCLUDED.max_attempts,
    scoring_method = EXCLUDED.scoring_method,
    required_score = EXCLUDED.required_score,
    short_questions_count = EXCLUDED.short_questions_count,
    normal_questions_count = EXCLUDED.normal_questions_count,
    long_questions_count = EXCLUDED.long_questions_count,
    published_at = EXCLUDED.published_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteCourseSubjectAssessment :exec
UPDATE course_subject_assessment
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetCourseSubjectAssessment :one
SELECT *
FROM course_subject_assessment
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
