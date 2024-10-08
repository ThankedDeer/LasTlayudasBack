// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: testimonials.sql

package sqlc

import (
	"context"
)

const createTestimonial = `-- name: CreateTestimonial :one
INSERT INTO testimonials (
    title,
    testimonial,
    user_id,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, DEFAULT, DEFAULT
) RETURNING testimonial_id, title, testimonial, user_id, created_at, updated_at
`

type CreateTestimonialParams struct {
	Title       string `json:"title"`
	Testimonial string `json:"testimonial"`
	UserID      int32  `json:"user_id"`
}

func (q *Queries) CreateTestimonial(ctx context.Context, arg CreateTestimonialParams) (Testimonials, error) {
	row := q.queryRow(ctx, q.createTestimonialStmt, createTestimonial, arg.Title, arg.Testimonial, arg.UserID)
	var i Testimonials
	err := row.Scan(
		&i.TestimonialID,
		&i.Title,
		&i.Testimonial,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTestimonial = `-- name: DeleteTestimonial :exec
DELETE FROM testimonials WHERE testimonial_id = $1
`

func (q *Queries) DeleteTestimonial(ctx context.Context, testimonialID int32) error {
	_, err := q.exec(ctx, q.deleteTestimonialStmt, deleteTestimonial, testimonialID)
	return err
}

const getTestimonial = `-- name: GetTestimonial :one
SELECT testimonial_id, title, testimonial, user_id, created_at, updated_at FROM testimonials
WHERE testimonial_id = $1 LIMIT 1
`

func (q *Queries) GetTestimonial(ctx context.Context, testimonialID int32) (Testimonials, error) {
	row := q.queryRow(ctx, q.getTestimonialStmt, getTestimonial, testimonialID)
	var i Testimonials
	err := row.Scan(
		&i.TestimonialID,
		&i.Title,
		&i.Testimonial,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTestimonials = `-- name: GetTestimonials :many
SELECT testimonial_id, title, testimonial, user_id, created_at, updated_at FROM testimonials
ORDER BY created_at DESC
`

func (q *Queries) GetTestimonials(ctx context.Context) ([]Testimonials, error) {
	rows, err := q.query(ctx, q.getTestimonialsStmt, getTestimonials)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Testimonials
	for rows.Next() {
		var i Testimonials
		if err := rows.Scan(
			&i.TestimonialID,
			&i.Title,
			&i.Testimonial,
			&i.UserID,
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

const getTestimonialsByUser = `-- name: GetTestimonialsByUser :many
SELECT testimonial_id, title, testimonial, user_id, created_at, updated_at FROM testimonials
WHERE user_id = $1
ORDER BY created_at DESC
`

func (q *Queries) GetTestimonialsByUser(ctx context.Context, userID int32) ([]Testimonials, error) {
	rows, err := q.query(ctx, q.getTestimonialsByUserStmt, getTestimonialsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Testimonials
	for rows.Next() {
		var i Testimonials
		if err := rows.Scan(
			&i.TestimonialID,
			&i.Title,
			&i.Testimonial,
			&i.UserID,
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

const updateTestimonial = `-- name: UpdateTestimonial :exec
UPDATE testimonials SET 
    title = $2,
    testimonial = $3,
    updated_at = DEFAULT
WHERE testimonial_id = $1
RETURNING testimonial_id, title, testimonial, user_id, created_at, updated_at
`

type UpdateTestimonialParams struct {
	TestimonialID int32  `json:"testimonial_id"`
	Title         string `json:"title"`
	Testimonial   string `json:"testimonial"`
}

func (q *Queries) UpdateTestimonial(ctx context.Context, arg UpdateTestimonialParams) error {
	_, err := q.exec(ctx, q.updateTestimonialStmt, updateTestimonial, arg.TestimonialID, arg.Title, arg.Testimonial)
	return err
}
