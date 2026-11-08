-- name: CreateLanguage :one
INSERT INTO language (
    guid,
    language_code,
    english_title,
    local_title,
    emoji,
    is_supported,
    is_beta
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateLanguage :one
UPDATE language
SET
    language_code = $2,
    english_title = $3,
    local_title = $4,
    emoji = $5,
    is_supported = $6,
    is_beta = $7
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveLanguage :one
INSERT INTO language (
    guid,
    language_code,
    english_title,
    local_title,
    emoji,
    is_supported,
    is_beta
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
ON CONFLICT (guid) DO UPDATE
SET
    language_code = EXCLUDED.language_code,
    english_title = EXCLUDED.english_title,
    local_title = EXCLUDED.local_title,
    emoji = EXCLUDED.emoji,
    is_supported = EXCLUDED.is_supported,
    is_beta = EXCLUDED.is_beta
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteLanguage :exec
UPDATE language
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetLanguage :one
SELECT *
FROM language
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
