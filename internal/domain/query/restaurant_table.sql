-- name: CreateRestaurantTable :one
INSERT INTO "restaurant_table" (number, waiter_id, status_id)
VALUES ($1, $2, $3)
RETURNING table_id, number, waiter_id, status_id, created_at, updated_at;

-- name: GetRestaurantTableByID :one
SELECT table_id, number, waiter_id, status_id, created_at, updated_at
FROM "restaurant_table"
WHERE table_id = $1;

-- name: GetAllRestaurantTables :many
SELECT table_id, number, waiter_id, status_id, created_at, updated_at
FROM "restaurant_table"
ORDER BY number;

-- name: UpdateRestaurantTable :exec
UPDATE "restaurant_table"
SET number = $2,
    waiter_id = $3,
    status_id = $4,
    updated_at = current_timestamp
WHERE table_id = $1;

-- name: DeleteRestaurantTable :exec
DELETE FROM "restaurant_table"
WHERE table_id = $1;
