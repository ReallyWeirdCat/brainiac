-- name: CreateAppUser :one
INSERT INTO app_user(username)
VALUES ($1)
RETURNING *;

-- name: CreateAppUserProfile :one
INSERT INTO app_user_profile(
    app_user_guid,
    name,
    surname,
    patronymic,
    nickname,
    bio,
    profile_discovery,
    avatar_url
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;
