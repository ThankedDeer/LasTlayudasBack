-- name: CreateWaiter :one
INSERT INTO "waiter" (user_id)
VALUES ($1)
RETURNING waiter_id, user_id, created_at;

-- name: GetWaiterByID :one
SELECT waiter_id, user_id, created_at
FROM "waiter"
WHERE waiter_id = $1;

-- name: GetWaiterByUserID :one
SELECT waiter_id, user_id, created_at
FROM "waiter"
WHERE user_id = $1;

-- name: GetAllWaiters :many
SELECT waiter_id, user_id, created_at
FROM "waiter";

-- name: DeleteWaiter :exec
DELETE FROM "waiter"
WHERE waiter_id = $1;
