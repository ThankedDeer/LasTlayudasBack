// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: role_permission.sql

package sqlc

import (
	"context"
)

const createRolePermission = `-- name: CreateRolePermission :one
INSERT INTO "role_permission" (role_id, permission_id)
VALUES ($1, $2)
RETURNING role_permission_id, role_id, permission_id, created_at
`

type CreateRolePermissionParams struct {
	RoleID       int32 `json:"role_id"`
	PermissionID int32 `json:"permission_id"`
}

func (q *Queries) CreateRolePermission(ctx context.Context, arg CreateRolePermissionParams) (RolePermission, error) {
	row := q.queryRow(ctx, q.createRolePermissionStmt, createRolePermission, arg.RoleID, arg.PermissionID)
	var i RolePermission
	err := row.Scan(
		&i.RolePermissionID,
		&i.RoleID,
		&i.PermissionID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteRolePermission = `-- name: DeleteRolePermission :exec
DELETE FROM "role_permission"
WHERE role_permission_id = $1
`

func (q *Queries) DeleteRolePermission(ctx context.Context, rolePermissionID int32) error {
	_, err := q.exec(ctx, q.deleteRolePermissionStmt, deleteRolePermission, rolePermissionID)
	return err
}

const getPermissionsByRoleID = `-- name: GetPermissionsByRoleID :many
SELECT rp.role_permission_id, rp.role_id, rp.permission_id, rp.created_at
FROM "role_permission" rp
JOIN "permission" p ON rp.permission_id = p.permission_id
WHERE rp.role_id = $1
ORDER BY p.name
`

func (q *Queries) GetPermissionsByRoleID(ctx context.Context, roleID int32) ([]RolePermission, error) {
	rows, err := q.query(ctx, q.getPermissionsByRoleIDStmt, getPermissionsByRoleID, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RolePermission
	for rows.Next() {
		var i RolePermission
		if err := rows.Scan(
			&i.RolePermissionID,
			&i.RoleID,
			&i.PermissionID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRolePermissionByID = `-- name: GetRolePermissionByID :one
SELECT role_permission_id, role_id, permission_id, created_at
FROM "role_permission"
WHERE role_permission_id = $1
`

func (q *Queries) GetRolePermissionByID(ctx context.Context, rolePermissionID int32) (RolePermission, error) {
	row := q.queryRow(ctx, q.getRolePermissionByIDStmt, getRolePermissionByID, rolePermissionID)
	var i RolePermission
	err := row.Scan(
		&i.RolePermissionID,
		&i.RoleID,
		&i.PermissionID,
		&i.CreatedAt,
	)
	return i, err
}

const getRolesByPermissionID = `-- name: GetRolesByPermissionID :many
SELECT rp.role_permission_id, rp.role_id, rp.permission_id, rp.created_at
FROM "role_permission" rp
JOIN "role" r ON rp.role_id = r.role_id
WHERE rp.permission_id = $1
ORDER BY r.name
`

func (q *Queries) GetRolesByPermissionID(ctx context.Context, permissionID int32) ([]RolePermission, error) {
	rows, err := q.query(ctx, q.getRolesByPermissionIDStmt, getRolesByPermissionID, permissionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RolePermission
	for rows.Next() {
		var i RolePermission
		if err := rows.Scan(
			&i.RolePermissionID,
			&i.RoleID,
			&i.PermissionID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
