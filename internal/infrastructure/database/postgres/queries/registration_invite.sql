-- name: CreateRegistrationInvite :one
INSERT INTO registration_invite (
    app_user_guid,
    invited_by_app_user_guid,
    invite_code,
    email,
    message,
    name,
    surname,
    patronymic,
    nickname,
    username,
    student_groups,
    teacher_groups,
    chats,
    expire_at,
    used_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
)
RETURNING *;

-- name: UpdateRegistrationInvite :one
UPDATE registration_invite
SET
    invited_by_app_user_guid = $2,
    invite_code = $3,
    email = $4,
    message = $5,
    name = $6,
    surname = $7,
    patronymic = $8,
    nickname = $9,
    username = $10,
    student_groups = $11,
    teacher_groups = $12,
    chats = $13,
    expire_at = $14,
    used_at = $15
WHERE app_user_guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveRegistrationInvite :one
INSERT INTO registration_invite (
    app_user_guid,
    invited_by_app_user_guid,
    invite_code,
    email,
    message,
    name,
    surname,
    patronymic,
    nickname,
    username,
    student_groups,
    teacher_groups,
    chats,
    expire_at,
    used_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
)
ON CONFLICT (app_user_guid) DO UPDATE
SET
    invited_by_app_user_guid = EXCLUDED.invited_by_app_user_guid,
    invite_code = EXCLUDED.invite_code,
    email = EXCLUDED.email,
    message = EXCLUDED.message,
    name = EXCLUDED.name,
    surname = EXCLUDED.surname,
    patronymic = EXCLUDED.patronymic,
    nickname = EXCLUDED.nickname,
    username = EXCLUDED.username,
    student_groups = EXCLUDED.student_groups,
    teacher_groups = EXCLUDED.teacher_groups,
    chats = EXCLUDED.chats,
    expire_at = EXCLUDED.expire_at,
    used_at = EXCLUDED.used_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteRegistrationInvite :exec
UPDATE registration_invite
SET deleted_at = now()
WHERE app_user_guid = $1 AND deleted_at IS NULL;

-- name: GetRegistrationInvite :one
SELECT *
FROM registration_invite
WHERE app_user_guid = $1 AND deleted_at IS NULL
LIMIT 1;
