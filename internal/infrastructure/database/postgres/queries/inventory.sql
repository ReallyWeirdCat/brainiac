-- name: CreateInventory :one
INSERT INTO inventory (
    app_user_guid,
    max_slots
) VALUES (
    $1, $2
)
RETURNING *;

-- name: UpdateInventory :one
UPDATE inventory
SET
    max_slots = $2
WHERE app_user_guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveInventory :one
INSERT INTO inventory (
    app_user_guid,
    max_slots
) VALUES (
    $1, $2
)
ON CONFLICT (app_user_guid) DO UPDATE
SET
    max_slots = EXCLUDED.max_slots
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteInventory :exec
UPDATE inventory
SET deleted_at = now()
WHERE app_user_guid = $1 AND deleted_at IS NULL;

-- name: GetInventory :one
SELECT *
FROM inventory
WHERE app_user_guid = $1 AND deleted_at IS NULL
LIMIT 1;
