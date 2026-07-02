-- name: CreateAppUserSession :one 
INSERT INTO app_user_session(guid, app_user_guid, last_ipv4, last_ipv6, last_agent, last_seen_at, expire_at)
VALUES
($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
