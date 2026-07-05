-- name: CreateAppUserSession :one
INSERT INTO app_user_session (
    guid,
    app_user_guid,
    last_ipv4,
    last_ipv6,
    last_agent,
    last_seen_at,
    expire_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateAppUserSession :one
UPDATE app_user_session
SET
    app_user_guid = $2,
    last_ipv4 = $3,
    last_ipv6 = $4,
    last_agent = $5,
    last_seen_at = $6,
    expire_at = $7
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveAppUserSession :one
INSERT INTO app_user_session (
    guid,
    app_user_guid,
    last_ipv4,
    last_ipv6,
    last_agent,
    last_seen_at,
    expire_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
ON CONFLICT (guid) DO UPDATE
SET
    app_user_guid = EXCLUDED.app_user_guid,
    last_ipv4 = EXCLUDED.last_ipv4,
    last_ipv6 = EXCLUDED.last_ipv6,
    last_agent = EXCLUDED.last_agent,
    last_seen_at = EXCLUDED.last_seen_at,
    expire_at = EXCLUDED.expire_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteAppUserSession :exec
UPDATE app_user_session
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetAppUserSession :one
SELECT *
FROM app_user_session
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetAllAppUserSessions :many
SELECT *
FROM app_user_session
WHERE deleted_at IS NULL;

-- name: CountAppUserSessions :one
SELECT COUNT(*)
FROM app_user_session
WHERE deleted_at IS NULL;

-- name: ExistsAppUserSession :one
SELECT EXISTS (
    SELECT 1
    FROM app_user_session
    WHERE guid = $1 AND deleted_at IS NULL
);

-- name: IsDeletedAppUserSession :one
SELECT coalesce(deleted_at IS NOT NULL, false)::boolean
FROM app_user_session
WHERE guid = $1;

-- name: CreateAppUserSessionBatch :batchmany
INSERT INTO app_user_session (
    guid,
    app_user_guid,
    last_ipv4,
    last_ipv6,
    last_agent,
    last_seen_at,
    expire_at
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateAppUserSessionBatch :batchmany
UPDATE app_user_session
SET
    app_user_guid = $2,
    last_ipv4 = $3,
    last_ipv6 = $4,
    last_agent = $5,
    last_seen_at = $6,
    expire_at = $7
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveAppUserSessionBatch :batchmany
INSERT INTO app_user_session (
    guid,
    app_user_guid,
    last_ipv4,
    last_ipv6,
    last_agent,
    last_seen_at,
    expire_at
) VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (guid) DO UPDATE
SET
    app_user_guid = EXCLUDED.app_user_guid,
    last_ipv4 = EXCLUDED.last_ipv4,
    last_ipv6 = EXCLUDED.last_ipv6,
    last_agent = EXCLUDED.last_agent,
    last_seen_at = EXCLUDED.last_seen_at,
    expire_at = EXCLUDED.expire_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteAppUserSessionBatch :batchexec
UPDATE app_user_session
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetAppUserSessionBatch :batchmany
SELECT * FROM app_user_session
WHERE guid = $1 AND deleted_at IS NULL;

-- name: ExistsAppUserSessionBatch :batchmany
SELECT guid FROM app_user_session
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetAllActiveSessionsByUsername :many
SELECT s.*
FROM app_user_session s
JOIN app_user u ON s.app_user_guid = u.guid
WHERE u.username = $1
  AND u.deleted_at IS NULL
  AND s.deleted_at IS NULL
  AND (s.expire_at IS NULL OR s.expire_at > now())
ORDER BY s.created_at DESC;

-- name: GetAllInactiveSessionsByUsername :many
SELECT s.*
FROM app_user_session s
JOIN app_user u ON s.app_user_guid = u.guid
WHERE u.username = $1
  AND u.deleted_at IS NULL
  AND s.deleted_at IS NULL
  AND s.expire_at IS NOT NULL
  AND s.expire_at <= now()
ORDER BY s.created_at DESC;

-- name: GetAllSessionsByLastIP :many
SELECT *
FROM app_user_session
WHERE (last_ipv4 = $1 OR last_ipv6 = $1)
  AND deleted_at IS NULL;

-- name: GetAllAppUserGUIDsByIP :many
SELECT DISTINCT app_user_guid
FROM app_user_session
WHERE (last_ipv4 = $1 OR last_ipv6 = $1)
  AND deleted_at IS NULL;
