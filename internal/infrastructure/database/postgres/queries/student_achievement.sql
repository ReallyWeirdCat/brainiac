-- name: CreateStudentAchievement :one
INSERT INTO student_achievement (
    guid,
    app_user_guid,
    achievement_guid
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateStudentAchievement :one
UPDATE student_achievement
SET
    app_user_guid = $2,
    achievement_guid = $3
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveStudentAchievement :one
INSERT INTO student_achievement (
    guid,
    app_user_guid,
    achievement_guid
) VALUES (
    $1, $2, $3
)
ON CONFLICT (guid) DO UPDATE
SET
    app_user_guid = EXCLUDED.app_user_guid,
    achievement_guid = EXCLUDED.achievement_guid
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteStudentAchievement :exec
UPDATE student_achievement
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetStudentAchievement :one
SELECT *
FROM student_achievement
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
