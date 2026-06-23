-- name: CreateAppUser :one
INSERT INTO app_user(username)
VALUES ($1)
RETURNING *;

-- name: SaveAppUser :one
INSERT INTO app_user (guid, username, created_at)
VALUES ($1, $2, $3)
ON CONFLICT (guid) 
DO UPDATE SET 
    username = EXCLUDED.username
RETURNING *;

-- name: DeleteAppUser :exec
DELETE FROM app_user
WHERE guid = $1;

-- name: GetAppUserByGUID :one
SELECT * FROM app_user
WHERE guid = $1
LIMIT 1;

-- name: GetAppUserByUsername :one
SELECT * FROM app_user
WHERE username = $1
LIMIT 1;

-- name: ExistsAppUserByUsername :one
SELECT EXISTS(
    SELECT 1 FROM app_user
    WHERE username = $1
) AS exists;
