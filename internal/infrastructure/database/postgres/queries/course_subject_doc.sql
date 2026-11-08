-- name: CreateCourseSubjectDoc :one
INSERT INTO course_subject_doc (
    guid,
    course_subject_guid,
    title_i18n,
    description_i18n,
    subject_doc_type,
    example_meta,
    submission_meta,
    literature_meta,
    doc_index,
    published_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: UpdateCourseSubjectDoc :one
UPDATE course_subject_doc
SET
    course_subject_guid = $2,
    title_i18n = $3,
    description_i18n = $4,
    subject_doc_type = $5,
    example_meta = $6,
    submission_meta = $7,
    literature_meta = $8,
    doc_index = $9,
    published_at = $10
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveCourseSubjectDoc :one
INSERT INTO course_subject_doc (
    guid,
    course_subject_guid,
    title_i18n,
    description_i18n,
    subject_doc_type,
    example_meta,
    submission_meta,
    literature_meta,
    doc_index,
    published_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
ON CONFLICT (guid) DO UPDATE
SET
    course_subject_guid = EXCLUDED.course_subject_guid,
    title_i18n = EXCLUDED.title_i18n,
    description_i18n = EXCLUDED.description_i18n,
    subject_doc_type = EXCLUDED.subject_doc_type,
    example_meta = EXCLUDED.example_meta,
    submission_meta = EXCLUDED.submission_meta,
    literature_meta = EXCLUDED.literature_meta,
    doc_index = EXCLUDED.doc_index,
    published_at = EXCLUDED.published_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteCourseSubjectDoc :exec
UPDATE course_subject_doc
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetCourseSubjectDoc :one
SELECT *
FROM course_subject_doc
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
