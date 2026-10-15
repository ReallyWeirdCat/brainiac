-- name: CreateStudentGlobalStats :one
INSERT INTO student_global_stats (
    app_user_guid,
    level,
    other_experience,
    money,
    money_earned,
    subjects_completed,
    assessments_completed,
    assessments_failed,
    assessments_terminated,
    works_submitted,
    items_collected,
    items_used,
    items_sold,
    items_exchanged,
    boxes_opened,
    max_daily_streak,
    messages_sent
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
)
RETURNING *;

-- name: UpdateStudentGlobalStats :one
UPDATE student_global_stats
SET
    level = $2,
    other_experience = $3,
    money = $4,
    money_earned = $5,
    subjects_completed = $6,
    assessments_completed = $7,
    assessments_failed = $8,
    assessments_terminated = $9,
    works_submitted = $10,
    items_collected = $11,
    items_used = $12,
    items_sold = $13,
    items_exchanged = $14,
    boxes_opened = $15,
    max_daily_streak = $16,
    messages_sent = $17
WHERE app_user_guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveStudentGlobalStats :one
INSERT INTO student_global_stats (
    app_user_guid,
    level,
    other_experience,
    money,
    money_earned,
    subjects_completed,
    assessments_completed,
    assessments_failed,
    assessments_terminated,
    works_submitted,
    items_collected,
    items_used,
    items_sold,
    items_exchanged,
    boxes_opened,
    max_daily_streak,
    messages_sent
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
)
ON CONFLICT (app_user_guid) DO UPDATE
SET
    level = EXCLUDED.level,
    other_experience = EXCLUDED.other_experience,
    money = EXCLUDED.money,
    money_earned = EXCLUDED.money_earned,
    subjects_completed = EXCLUDED.subjects_completed,
    assessments_completed = EXCLUDED.assessments_completed,
    assessments_failed = EXCLUDED.assessments_failed,
    assessments_terminated = EXCLUDED.assessments_terminated,
    works_submitted = EXCLUDED.works_submitted,
    items_collected = EXCLUDED.items_collected,
    items_used = EXCLUDED.items_used,
    items_sold = EXCLUDED.items_sold,
    items_exchanged = EXCLUDED.items_exchanged,
    boxes_opened = EXCLUDED.boxes_opened,
    max_daily_streak = EXCLUDED.max_daily_streak,
    messages_sent = EXCLUDED.messages_sent
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteStudentGlobalStats :exec
UPDATE student_global_stats
SET deleted_at = now()
WHERE app_user_guid = $1 AND deleted_at IS NULL;

-- name: GetStudentGlobalStats :one
SELECT *
FROM student_global_stats
WHERE app_user_guid = $1 AND deleted_at IS NULL
LIMIT 1;
