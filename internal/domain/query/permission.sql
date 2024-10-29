-- name: CreatePermission :one
INSERT INTO "permission" (name, description)
VALUES ($1, $2)
RETURNING permission_id, name, description, created_at, updated_at;

-- name: GetPermissionByID :one
SELECT permission_id, name, description, created_at, updated_at
FROM "permission"
WHERE permission_id = $1;

-- name: GetPermissionByName :one
SELECT permission_id, name, description, created_at, updated_at
FROM "permission"
WHERE name = $1;

-- name: GetAllPermissions :many
SELECT permission_id, name, description, created_at, updated_at
FROM "permission"
ORDER BY name;

-- name: UpdatePermission :exec
UPDATE "permission"
SET name = $2,
    description = $3,
    updated_at = current_timestamp
WHERE permission_id = $1;

-- name: DeletePermission :exec
DELETE FROM "permission"
WHERE permission_id = $1;
