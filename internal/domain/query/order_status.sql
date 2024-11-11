-- name: CreateOrderStatus :one
INSERT INTO "order_status" (name, description)
VALUES ($1, $2)
RETURNING order_status_id, name, description, created_at;

-- name: GetOrderStatusByID :one
SELECT order_status_id, name, description, created_at
FROM "order_status"
WHERE order_status_id = $1;

-- name: GetAllOrderStatuses :many
SELECT order_status_id, name, description, created_at
FROM "order_status"
ORDER BY name;

-- name: UpdateOrderStatus :exec
UPDATE "order_status"
SET name = $2,
    description = $3
WHERE order_status_id = $1;

-- name: DeleteOrderStatus :exec
DELETE FROM "order_status"
WHERE order_status_id = $1;
