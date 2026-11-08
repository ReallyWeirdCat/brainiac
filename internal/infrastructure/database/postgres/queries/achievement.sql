-- name: CreateAchievement :one
INSERT INTO achievement (
    guid,
    title_i18n,
    description_i18n,
    resource_url,
    is_hidden,
    meta,
    conditions,
    reward_experience,
    reward_money,
    published_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: UpdateAchievement :one
UPDATE achievement
SET
    title_i18n = $2,
    description_i18n = $3,
    resource_url = $4,
    is_hidden = $5,
    meta = $6,
    conditions = $7,
    reward_experience = $8,
    reward_money = $9,
    published_at = $10
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveAchievement :one
INSERT INTO achievement (
    guid,
    title_i18n,
    description_i18n,
    resource_url,
    is_hidden,
    meta,
    conditions,
    reward_experience,
    reward_money,
    published_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
ON CONFLICT (guid) DO UPDATE
SET
    title_i18n = EXCLUDED.title_i18n,
    description_i18n = EXCLUDED.description_i18n,
    resource_url = EXCLUDED.resource_url,
    is_hidden = EXCLUDED.is_hidden,
    meta = EXCLUDED.meta,
    conditions = EXCLUDED.conditions,
    reward_experience = EXCLUDED.reward_experience,
    reward_money = EXCLUDED.reward_money,
    published_at = EXCLUDED.published_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteAchievement :exec
UPDATE achievement
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetAchievement :one
SELECT *
FROM achievement
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
