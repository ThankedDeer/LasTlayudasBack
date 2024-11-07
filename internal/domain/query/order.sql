-- name: CreateOrder :one
INSERT INTO "order" (order_date, table_id, status_id)
VALUES (current_timestamp, $1, $2)
RETURNING order_id, order_date, table_id, status_id;

-- name: GetOrderByID :one
SELECT order_id, order_date, table_id, status_id
FROM "order"
WHERE order_id = $1;

-- name: GetAllOrders :many
SELECT order_id, order_date, table_id, status_id
FROM "order"
ORDER BY order_date DESC;

-- name: UpdateOrder :exec
UPDATE "order"
SET table_id = $2,
    status_id = $3
WHERE order_id = $1;

-- name: DeleteOrder :exec
DELETE FROM "order"
WHERE order_id = $1;
