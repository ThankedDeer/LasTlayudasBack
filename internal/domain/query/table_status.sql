-- name: CreateTableStatus :one
INSERT INTO "table_status" (name, description)
VALUES ($1, $2)
RETURNING table_status_id, name, description, created_at;

-- name: GetTableStatusByID :one
SELECT table_status_id, name, description, created_at
FROM "table_status"
WHERE table_status_id = $1;

-- name: GetAllTableStatuses :many
SELECT table_status_id, name, description, created_at
FROM "table_status"
ORDER BY name;

-- name: UpdateTableStatus :exec
UPDATE "table_status"
SET name = $2,
    description = $3
WHERE table_status_id = $1;

-- name: DeleteTableStatus :exec
DELETE FROM "table_status"
WHERE table_status_id = $1;
