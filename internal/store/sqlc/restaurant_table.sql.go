// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: restaurant_table.sql

package sqlc

import (
	"context"
)

const createRestaurantTable = `-- name: CreateRestaurantTable :one
INSERT INTO "restaurant_table" (number,status_id)
VALUES ($1, $2)
RETURNING table_id, number, status_id, created_at, updated_at
`

type CreateRestaurantTableParams struct {
	Number   int32
	StatusID int32
}

func (q *Queries) CreateRestaurantTable(ctx context.Context, arg CreateRestaurantTableParams) (RestaurantTable, error) {
	row := q.db.QueryRowContext(ctx, createRestaurantTable, arg.Number, arg.StatusID)
	var i RestaurantTable
	err := row.Scan(
		&i.TableID,
		&i.Number,
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
	_, err := q.db.ExecContext(ctx, deleteRestaurantTable, tableID)
	return err
}

const getAllRestaurantTables = `-- name: GetAllRestaurantTables :many
SELECT table_id, number, status_id, created_at, updated_at
FROM "restaurant_table"
ORDER BY number
`

func (q *Queries) GetAllRestaurantTables(ctx context.Context) ([]RestaurantTable, error) {
	rows, err := q.db.QueryContext(ctx, getAllRestaurantTables)
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
SELECT table_id, number, status_id, created_at, updated_at
FROM "restaurant_table"
WHERE table_id = $1
`

func (q *Queries) GetRestaurantTableByID(ctx context.Context, tableID int32) (RestaurantTable, error) {
	row := q.db.QueryRowContext(ctx, getRestaurantTableByID, tableID)
	var i RestaurantTable
	err := row.Scan(
		&i.TableID,
		&i.Number,
		&i.StatusID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateRestaurantTable = `-- name: UpdateRestaurantTable :exec
UPDATE "restaurant_table"
SET number = $2,
    status_id = $3,
    updated_at = current_timestamp
WHERE table_id = $1
`

type UpdateRestaurantTableParams struct {
	TableID  int32
	Number   int32
	StatusID int32
}

func (q *Queries) UpdateRestaurantTable(ctx context.Context, arg UpdateRestaurantTableParams) error {
	_, err := q.db.ExecContext(ctx, updateRestaurantTable, arg.TableID, arg.Number, arg.StatusID)
	return err
}
