-- name: CreateCourseSubjectDocMd :one
INSERT INTO course_subject_doc_md (
    guid,
    course_subject_doc_guid,
    language_code,
    md,
    published_at
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateCourseSubjectDocMd :one
UPDATE course_subject_doc_md
SET
    course_subject_doc_guid = $2,
    language_code = $3,
    md = $4,
    published_at = $5
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveCourseSubjectDocMd :one
INSERT INTO course_subject_doc_md (
    guid,
    course_subject_doc_guid,
    language_code,
    md,
    published_at
) VALUES (
    $1, $2, $3, $4, $5
)
ON CONFLICT (guid) DO UPDATE
SET
    course_subject_doc_guid = EXCLUDED.course_subject_doc_guid,
    language_code = EXCLUDED.language_code,
    md = EXCLUDED.md,
    published_at = EXCLUDED.published_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteCourseSubjectDocMd :exec
UPDATE course_subject_doc_md
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetCourseSubjectDocMd :one
SELECT *
FROM course_subject_doc_md
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
