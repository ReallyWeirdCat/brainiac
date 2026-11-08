-- name: CreateItem :one
INSERT INTO item (
    guid,
    title_i18n,
    description_i18n,
    resource_url,
    meta,
    stack_size,
    rarity,
    allow_exchange,
    shop_price,
    one_time_purchase,
    shop_quantity,
    level_required,
    on_sale_since,
    on_sale_until,
    published_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
)
RETURNING *;

-- name: UpdateItem :one
UPDATE item
SET
    title_i18n = $2,
    description_i18n = $3,
    resource_url = $4,
    meta = $5,
    stack_size = $6,
    rarity = $7,
    allow_exchange = $8,
    shop_price = $9,
    one_time_purchase = $10,
    shop_quantity = $11,
    level_required = $12,
    on_sale_since = $13,
    on_sale_until = $14,
    published_at = $15
WHERE guid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SaveItem :one
INSERT INTO item (
    guid,
    title_i18n,
    description_i18n,
    resource_url,
    meta,
    stack_size,
    rarity,
    allow_exchange,
    shop_price,
    one_time_purchase,
    shop_quantity,
    level_required,
    on_sale_since,
    on_sale_until,
    published_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
)
ON CONFLICT (guid) DO UPDATE
SET
    title_i18n = EXCLUDED.title_i18n,
    description_i18n = EXCLUDED.description_i18n,
    resource_url = EXCLUDED.resource_url,
    meta = EXCLUDED.meta,
    stack_size = EXCLUDED.stack_size,
    rarity = EXCLUDED.rarity,
    allow_exchange = EXCLUDED.allow_exchange,
    shop_price = EXCLUDED.shop_price,
    one_time_purchase = EXCLUDED.one_time_purchase,
    shop_quantity = EXCLUDED.shop_quantity,
    level_required = EXCLUDED.level_required,
    on_sale_since = EXCLUDED.on_sale_since,
    on_sale_until = EXCLUDED.on_sale_until,
    published_at = EXCLUDED.published_at
WHERE deleted_at IS NULL
RETURNING *;

-- name: DeleteItem :exec
UPDATE item
SET deleted_at = now()
WHERE guid = $1 AND deleted_at IS NULL;

-- name: GetItem :one
SELECT *
FROM item
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
