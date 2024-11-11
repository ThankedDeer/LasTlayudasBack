-- name: CreateProduct :one
INSERT INTO "product" 
("name", "purchase_price", "sale_price", "stock", "category_id", "provider_id")
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING "product_id", "created_at", "updated_at";

-- name: GetProductByID :one
SELECT "product_id", "name", "purchase_price", "sale_price", "stock", "category_id", "provider_id", "created_at", "updated_at"
FROM "product"
WHERE "product_id" = $1;

-- name: GetAllProducts :many
SELECT "product_id", "name", "purchase_price", "sale_price", "stock", "category_id", "provider_id", "created_at", "updated_at"
FROM "product";

-- name: UpdateProduct :one
UPDATE "product"
SET "name" = $1, 
    "purchase_price" = $2, 
    "sale_price" = $3, 
    "stock" = $4, 
    "category_id" = $5, 
    "provider_id" = $6, 
    "updated_at" = current_timestamp
WHERE "product_id" = $7
RETURNING "updated_at";

-- name: DeleteProduct :exec
DELETE FROM "product"
WHERE "product_id" = $1;
