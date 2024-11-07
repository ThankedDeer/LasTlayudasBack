-- name: CreateRolePermission :one
INSERT INTO "role_permission" (role_id, permission_id)
VALUES ($1, $2)
RETURNING role_permission_id, role_id, permission_id, created_at;

-- name: GetRolePermissionByID :one
SELECT role_permission_id, role_id, permission_id, created_at
FROM "role_permission"
WHERE role_permission_id = $1;

-- name: GetPermissionsByRoleID :many
SELECT rp.role_permission_id, rp.role_id, rp.permission_id, rp.created_at
FROM "role_permission" rp
JOIN "permission" p ON rp.permission_id = p.permission_id
WHERE rp.role_id = $1
ORDER BY p.name;

-- name: GetRolesByPermissionID :many
SELECT rp.role_permission_id, rp.role_id, rp.permission_id, rp.created_at
FROM "role_permission" rp
JOIN "role" r ON rp.role_id = r.role_id
WHERE rp.permission_id = $1
ORDER BY r.name;

-- name: DeleteRolePermission :exec
DELETE FROM "role_permission"
WHERE role_permission_id = $1;
