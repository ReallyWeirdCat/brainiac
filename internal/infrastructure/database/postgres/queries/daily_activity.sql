-- name: CreateDailyActivity :one
INSERT INTO daily_activity (
    guid,
    app_user_guid,
    day,
    experience_earned,
    levels_earned,
    subjects_completed,
    assessments_completed,
    practices_completed
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateDailyActivity :one
UPDATE daily_activity
SET
    app_user_guid = $2,
    day = $3,
    experience_earned = $4,
    levels_earned = $5,
    subjects_completed = $6,
    assessments_completed = $7,
    practices_completed = $8
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveDailyActivity :one
INSERT INTO daily_activity (
    guid,
    app_user_guid,
    day,
    experience_earned,
    levels_earned,
    subjects_completed,
    assessments_completed,
    practices_completed
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
ON CONFLICT (guid) DO UPDATE
SET
    app_user_guid = EXCLUDED.app_user_guid,
    day = EXCLUDED.day,
    experience_earned = EXCLUDED.experience_earned,
    levels_earned = EXCLUDED.levels_earned,
    subjects_completed = EXCLUDED.subjects_completed,
    assessments_completed = EXCLUDED.assessments_completed,
    practices_completed = EXCLUDED.practices_completed
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteDailyActivity :exec
UPDATE daily_activity
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetDailyActivity :one
SELECT *
FROM daily_activity
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
