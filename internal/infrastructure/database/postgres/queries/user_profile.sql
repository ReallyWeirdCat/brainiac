-- name: CreateAppUserProfile :one
INSERT INTO app_user_profile (
    app_user_guid,
    name,
    surname,
    patronymic,
    nickname,
    bio,
    preferred_language,
    profile_discovery,
    avatar_url,
    editing_locked_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: UpdateAppUserProfile :one
UPDATE app_user_profile
SET
    name = $2,
    surname = $3,
    patronymic = $4,
    nickname = $5,
    bio = $6,
    preferred_language = $7,
    profile_discovery = $8,
    avatar_url = $9,
    editing_locked_at = $10
WHERE app_user_guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveAppUserProfile :one
INSERT INTO app_user_profile (
    app_user_guid,
    name,
    surname,
    patronymic,
    nickname,
    bio,
    preferred_language,
    profile_discovery,
    avatar_url,
    editing_locked_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
ON CONFLICT (app_user_guid) DO UPDATE
SET
    name = EXCLUDED.name,
    surname = EXCLUDED.surname,
    patronymic = EXCLUDED.patronymic,
    nickname = EXCLUDED.nickname,
    bio = EXCLUDED.bio,
    preferred_language = EXCLUDED.preferred_language,
    profile_discovery = EXCLUDED.profile_discovery,
    avatar_url = EXCLUDED.avatar_url,
    editing_locked_at = EXCLUDED.editing_locked_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteAppUserProfile :exec
UPDATE app_user_profile
SET deleted_at = now()
WHERE app_user_guid = $1 AND deleted_at IS NULL;

-- name: GetAppUserProfile :one
SELECT *
FROM app_user_profile
WHERE app_user_guid = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetAllAppUserProfiles :many
SELECT *
FROM app_user_profile
WHERE deleted_at IS NULL;

-- name: CountAppUserProfiles :one
SELECT COUNT(*)
FROM app_user_profile
WHERE deleted_at IS NULL;

-- name: ExistsAppUserProfile :one
SELECT EXISTS (
    SELECT 1
    FROM app_user_profile
    WHERE app_user_guid = $1 AND deleted_at IS NULL
);

-- name: IsDeletedAppUserProfile :one
SELECT coalesce(deleted_at IS NOT NULL, false)::boolean
FROM app_user_profile
WHERE app_user_guid = $1;

-- name: CreateAppUserProfileBatch :batchmany
INSERT INTO app_user_profile (
    app_user_guid,
    name,
    surname,
    patronymic,
    nickname,
    bio,
    preferred_language,
    profile_discovery,
    avatar_url,
    editing_locked_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: UpdateAppUserProfileBatch :batchmany
UPDATE app_user_profile
SET
    name = $2,
    surname = $3,
    patronymic = $4,
    nickname = $5,
    bio = $6,
    preferred_language = $7,
    profile_discovery = $8,
    avatar_url = $9,
    editing_locked_at = $10
WHERE app_user_guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveAppUserProfileBatch :batchmany
INSERT INTO app_user_profile (
    app_user_guid,
    name,
    surname,
    patronymic,
    nickname,
    bio,
    preferred_language,
    profile_discovery,
    avatar_url,
    editing_locked_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT (app_user_guid) DO UPDATE
SET
    name = EXCLUDED.name,
    surname = EXCLUDED.surname,
    patronymic = EXCLUDED.patronymic,
    nickname = EXCLUDED.nickname,
    bio = EXCLUDED.bio,
    preferred_language = EXCLUDED.preferred_language,
    profile_discovery = EXCLUDED.profile_discovery,
    avatar_url = EXCLUDED.avatar_url,
    editing_locked_at = EXCLUDED.editing_locked_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteAppUserProfileBatch :batchexec
UPDATE app_user_profile
SET deleted_at = now()
WHERE app_user_guid = $1 AND deleted_at IS NULL;

-- name: GetAppUserProfileBatch :batchmany
SELECT * FROM app_user_profile
WHERE app_user_guid = $1 AND deleted_at IS NULL;

-- name: ExistsAppUserProfileBatch :batchmany
SELECT app_user_guid FROM app_user_profile
WHERE app_user_guid = $1 AND deleted_at IS NULL;

-- name: GetAppUserProfileByUsername :one
SELECT p.*
FROM app_user_profile p
JOIN app_user u ON p.app_user_guid = u.guid
WHERE u.username = $1
  AND u.deleted_at IS NULL
  AND p.deleted_at IS NULL
LIMIT 1;
