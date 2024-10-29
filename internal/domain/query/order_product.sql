-- name: CreateOrderProduct :one
INSERT INTO "order_product" (order_id, product_id, quantity)
VALUES ($1, $2, COALESCE($3, 1))
RETURNING order_product_id, order_id, product_id, quantity;

-- name: GetOrderProductByID :one
SELECT order_product_id, order_id, product_id, quantity
FROM "order_product"
WHERE order_product_id = $1;

-- name: GetProductsByOrderID :many
SELECT order_product_id, order_id, product_id, quantity
FROM "order_product"
WHERE order_id = $1;

-- name: GetOrdersByProductID :many
SELECT order_product_id, order_id, product_id, quantity
FROM "order_product"
WHERE product_id = $1;

-- name: UpdateOrderProduct :exec
UPDATE "order_product"
SET quantity = $2
WHERE order_product_id = $1;

-- name: DeleteOrderProduct :exec
DELETE FROM "order_product"
WHERE order_product_id = $1;
