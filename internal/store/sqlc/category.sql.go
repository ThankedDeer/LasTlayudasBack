package sqlc

import (
	"context"
	"database/sql"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO category (name, description, active)
VALUES ($1, $2, COALESCE($3, true))
RETURNING category_id, name, description, active, created_at, updated_at
`

type CreateCategoryParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Column3     interface{}    `json:"column_3"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.queryRow(ctx, q.createCategoryStmt, createCategory, arg.Name, arg.Description, arg.Column3)
	var i Category
	err := row.Scan(
		&i.CategoryID,
		&i.Name,
		&i.Description,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM category
WHERE category_id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, categoryID int32) error {
	_, err := q.exec(ctx, q.deleteCategoryStmt, deleteCategory, categoryID)
	return err
}

const getAllCategories = `-- name: GetAllCategories :many
SELECT category_id, name, description, active, created_at, updated_at
FROM category
ORDER BY name
`

func (q *Queries) GetAllCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.query(ctx, q.getAllCategoriesStmt, getAllCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.CategoryID,
			&i.Name,
			&i.Description,
			&i.Active,
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

const getCategoryByID = `-- name: GetCategoryByID :one
SELECT category_id, name, description, active, created_at, updated_at
FROM category
WHERE category_id = $1
`

func (q *Queries) GetCategoryByID(ctx context.Context, categoryID int32) (Category, error) {
	row := q.queryRow(ctx, q.getCategoryByIDStmt, getCategoryByID, categoryID)
	var i Category
	err := row.Scan(
		&i.CategoryID,
		&i.Name,
		&i.Description,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCategoryByName = `-- name: GetCategoryByName :one
SELECT category_id, name, description, active, created_at, updated_at
FROM category
WHERE name = $1
`

func (q *Queries) GetCategoryByName(ctx context.Context, name string) (Category, error) {
	row := q.queryRow(ctx, q.getCategoryByNameStmt, getCategoryByName, name)
	var i Category
	err := row.Scan(
		&i.CategoryID,
		&i.Name,
		&i.Description,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE category
SET name = $2,
    description = $3,
    active = $4,
    updated_at = current_timestamp
WHERE category_id = $1
`

type UpdateCategoryParams struct {
	CategoryID  int32          `json:"category_id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Active      sql.NullBool   `json:"active"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.exec(ctx, q.updateCategoryStmt, updateCategory,
		arg.CategoryID,
		arg.Name,
		arg.Description,
		arg.Active,
	)
	return err
}
