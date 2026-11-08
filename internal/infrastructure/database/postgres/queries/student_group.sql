-- name: CreateStudentGroup :one
INSERT INTO student_group (
    guid,
    title,
    description,
    enable_gamification
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateStudentGroup :one
UPDATE student_group
SET
    title = $2,
    description = $3,
    enable_gamification = $4
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveStudentGroup :one
INSERT INTO student_group (
    guid,
    title,
    description,
    enable_gamification
) VALUES (
    $1, $2, $3, $4
)
ON CONFLICT (guid) DO UPDATE
SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    enable_gamification = EXCLUDED.enable_gamification
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteStudentGroup :exec
UPDATE student_group
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetStudentGroup :one
SELECT *
FROM student_group
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
