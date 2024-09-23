-- name: CreateTestimonial :one
INSERT INTO testimonials (
    title,
    testimonial,
    user_id,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, DEFAULT, DEFAULT
) RETURNING *;

-- name: GetTestimonial :one
SELECT * FROM testimonials
WHERE testimonial_id = $1 LIMIT 1;

-- name: GetTestimonials :many
SELECT * FROM testimonials
ORDER BY created_at DESC;

-- name: GetTestimonialsByUser :many
SELECT * FROM testimonials
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateTestimonial :exec
UPDATE testimonials SET 
    title = $2,
    testimonial = $3,
    updated_at = DEFAULT
WHERE testimonial_id = $1
RETURNING *;

-- name: DeleteTestimonial :exec
DELETE FROM testimonials WHERE testimonial_id = $1;
