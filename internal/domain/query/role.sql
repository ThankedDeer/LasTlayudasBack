-- name: CreateRole :one
INSERT INTO "role" (name, description)
VALUES ($1, $2)
RETURNING role_id, name, description, created_at, updated_at;

-- name: GetRoleByID :one
SELECT role_id, name, description, created_at, updated_at
FROM "role"
WHERE role_id = $1;

-- name: GetRoleByName :one
SELECT role_id, name, description, created_at, updated_at
FROM "role"
WHERE name = $1;

-- name: GetAllRoles :many
SELECT role_id, name, description, created_at, updated_at
FROM "role"
ORDER BY name;

-- name: UpdateRole :exec
UPDATE "role"
SET name = $2,
    description = $3,
    updated_at = current_timestamp
WHERE role_id = $1;

-- name: DeleteRole :exec
DELETE FROM "role"
WHERE role_id = $1;
