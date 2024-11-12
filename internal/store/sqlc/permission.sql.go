// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: permission.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createPermission = `-- name: CreatePermission :one
INSERT INTO "permission" (name, description)
VALUES ($1, $2)
RETURNING permission_id, name, description, created_at, updated_at
`

type CreatePermissionParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) CreatePermission(ctx context.Context, arg CreatePermissionParams) (Permission, error) {
	row := q.queryRow(ctx, q.createPermissionStmt, createPermission, arg.Name, arg.Description)
	var i Permission
	err := row.Scan(
		&i.PermissionID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePermission = `-- name: DeletePermission :exec
DELETE FROM "permission"
WHERE permission_id = $1
`

func (q *Queries) DeletePermission(ctx context.Context, permissionID int32) error {
	_, err := q.exec(ctx, q.deletePermissionStmt, deletePermission, permissionID)
	return err
}

const getAllPermissions = `-- name: GetAllPermissions :many
SELECT permission_id, name, description, created_at, updated_at
FROM "permission"
ORDER BY name
`

func (q *Queries) GetAllPermissions(ctx context.Context) ([]Permission, error) {
	rows, err := q.query(ctx, q.getAllPermissionsStmt, getAllPermissions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Permission
	for rows.Next() {
		var i Permission
		if err := rows.Scan(
			&i.PermissionID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getPermissionByID = `-- name: GetPermissionByID :one
SELECT permission_id, name, description, created_at, updated_at
FROM "permission"
WHERE permission_id = $1
`

func (q *Queries) GetPermissionByID(ctx context.Context, permissionID int32) (Permission, error) {
	row := q.queryRow(ctx, q.getPermissionByIDStmt, getPermissionByID, permissionID)
	var i Permission
	err := row.Scan(
		&i.PermissionID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPermissionByName = `-- name: GetPermissionByName :one
SELECT permission_id, name, description, created_at, updated_at
FROM "permission"
WHERE name = $1
`

func (q *Queries) GetPermissionByName(ctx context.Context, name string) (Permission, error) {
	row := q.queryRow(ctx, q.getPermissionByNameStmt, getPermissionByName, name)
	var i Permission
	err := row.Scan(
		&i.PermissionID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePermission = `-- name: UpdatePermission :exec
UPDATE "permission"
SET name = $2,
    description = $3,
    updated_at = current_timestamp
WHERE permission_id = $1
`

type UpdatePermissionParams struct {
	PermissionID int32          `json:"permission_id"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
}

func (q *Queries) UpdatePermission(ctx context.Context, arg UpdatePermissionParams) error {
	_, err := q.exec(ctx, q.updatePermissionStmt, updatePermission, arg.PermissionID, arg.Name, arg.Description)
	return err
}
