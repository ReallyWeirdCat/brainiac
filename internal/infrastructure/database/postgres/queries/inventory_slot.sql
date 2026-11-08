-- name: CreateInventorySlot :one
INSERT INTO inventory_slot (
    guid,
    inventory_guid,
    item_guid,
    amount,
    amount_in_exchange,
    expire_at
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateInventorySlot :one
UPDATE inventory_slot
SET
    inventory_guid = $2,
    item_guid = $3,
    amount = $4,
    amount_in_exchange = $5,
    expire_at = $6
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveInventorySlot :one
INSERT INTO inventory_slot (
    guid,
    inventory_guid,
    item_guid,
    amount,
    amount_in_exchange,
    expire_at
) VALUES (
    $1, $2, $3, $4, $5, $6
)
ON CONFLICT (guid) DO UPDATE
SET
    inventory_guid = EXCLUDED.inventory_guid,
    item_guid = EXCLUDED.item_guid,
    amount = EXCLUDED.amount,
    amount_in_exchange = EXCLUDED.amount_in_exchange,
    expire_at = EXCLUDED.expire_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteInventorySlot :exec
UPDATE inventory_slot
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetInventorySlot :one
SELECT *
FROM inventory_slot
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
