// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: waiter.sql

package sqlc

import (
	"context"
)

const createWaiter = `-- name: CreateWaiter :one
INSERT INTO "waiter" (user_id)
VALUES ($1)
RETURNING waiter_id, user_id, created_at
`

func (q *Queries) CreateWaiter(ctx context.Context, userID int32) (Waiter, error) {
	row := q.db.QueryRowContext(ctx, createWaiter, userID)
	var i Waiter
	err := row.Scan(&i.WaiterID, &i.UserID, &i.CreatedAt)
	return i, err
}

const deleteWaiter = `-- name: DeleteWaiter :exec
DELETE FROM "waiter"
WHERE waiter_id = $1
`

func (q *Queries) DeleteWaiter(ctx context.Context, waiterID int32) error {
	_, err := q.db.ExecContext(ctx, deleteWaiter, waiterID)
	return err
}

const getAllWaiters = `-- name: GetAllWaiters :many
SELECT waiter_id, user_id, created_at
FROM "waiter"
`

func (q *Queries) GetAllWaiters(ctx context.Context) ([]Waiter, error) {
	rows, err := q.db.QueryContext(ctx, getAllWaiters)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Waiter
	for rows.Next() {
		var i Waiter
		if err := rows.Scan(&i.WaiterID, &i.UserID, &i.CreatedAt); err != nil {
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

const getWaiterByID = `-- name: GetWaiterByID :one
SELECT waiter_id, user_id, created_at
FROM "waiter"
WHERE waiter_id = $1
`

func (q *Queries) GetWaiterByID(ctx context.Context, waiterID int32) (Waiter, error) {
	row := q.db.QueryRowContext(ctx, getWaiterByID, waiterID)
	var i Waiter
	err := row.Scan(&i.WaiterID, &i.UserID, &i.CreatedAt)
	return i, err
}

const getWaiterByUserID = `-- name: GetWaiterByUserID :one
SELECT waiter_id, user_id, created_at
FROM "waiter"
WHERE user_id = $1
`

func (q *Queries) GetWaiterByUserID(ctx context.Context, userID int32) (Waiter, error) {
	row := q.db.QueryRowContext(ctx, getWaiterByUserID, userID)
	var i Waiter
	err := row.Scan(&i.WaiterID, &i.UserID, &i.CreatedAt)
	return i, err
}
