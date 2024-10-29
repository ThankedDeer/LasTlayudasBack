// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: restaurant_table.sql

package sqlc

import (
	"context"
)

const createRestaurantTable = `-- name: CreateRestaurantTable :one
INSERT INTO "restaurant_table" (number, waiter_id, status_id)
VALUES ($1, $2, $3)
RETURNING table_id, number, waiter_id, status_id, created_at, updated_at
`

type CreateRestaurantTableParams struct {
	Number   int32 `json:"number"`
	WaiterID int32 `json:"waiter_id"`
	StatusID int32 `json:"status_id"`
}

func (q *Queries) CreateRestaurantTable(ctx context.Context, arg CreateRestaurantTableParams) (RestaurantTable, error) {
	row := q.queryRow(ctx, q.createRestaurantTableStmt, createRestaurantTable, arg.Number, arg.WaiterID, arg.StatusID)
	var i RestaurantTable
	err := row.Scan(
		&i.TableID,
		&i.Number,
		&i.WaiterID,
		&i.StatusID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteRestaurantTable = `-- name: DeleteRestaurantTable :exec
DELETE FROM "restaurant_table"
WHERE table_id = $1
`

func (q *Queries) DeleteRestaurantTable(ctx context.Context, tableID int32) error {
	_, err := q.exec(ctx, q.deleteRestaurantTableStmt, deleteRestaurantTable, tableID)
	return err
}

const getAllRestaurantTables = `-- name: GetAllRestaurantTables :many
SELECT table_id, number, waiter_id, status_id, created_at, updated_at
FROM "restaurant_table"
ORDER BY number
`

func (q *Queries) GetAllRestaurantTables(ctx context.Context) ([]RestaurantTable, error) {
	rows, err := q.query(ctx, q.getAllRestaurantTablesStmt, getAllRestaurantTables)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RestaurantTable
	for rows.Next() {
		var i RestaurantTable
		if err := rows.Scan(
			&i.TableID,
			&i.Number,
			&i.WaiterID,
			&i.StatusID,
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

const getRestaurantTableByID = `-- name: GetRestaurantTableByID :one
SELECT table_id, number, waiter_id, status_id, created_at, updated_at
FROM "restaurant_table"
WHERE table_id = $1
`

func (q *Queries) GetRestaurantTableByID(ctx context.Context, tableID int32) (RestaurantTable, error) {
	row := q.queryRow(ctx, q.getRestaurantTableByIDStmt, getRestaurantTableByID, tableID)
	var i RestaurantTable
	err := row.Scan(
		&i.TableID,
		&i.Number,
		&i.WaiterID,
		&i.StatusID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateRestaurantTable = `-- name: UpdateRestaurantTable :exec
UPDATE "restaurant_table"
SET number = $2,
    waiter_id = $3,
    status_id = $4,
    updated_at = current_timestamp
WHERE table_id = $1
`

type UpdateRestaurantTableParams struct {
	TableID  int32 `json:"table_id"`
	Number   int32 `json:"number"`
	WaiterID int32 `json:"waiter_id"`
	StatusID int32 `json:"status_id"`
}

func (q *Queries) UpdateRestaurantTable(ctx context.Context, arg UpdateRestaurantTableParams) error {
	_, err := q.exec(ctx, q.updateRestaurantTableStmt, updateRestaurantTable,
		arg.TableID,
		arg.Number,
		arg.WaiterID,
		arg.StatusID,
	)
	return err
}