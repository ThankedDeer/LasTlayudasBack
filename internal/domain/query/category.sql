-- name: CreateCategory :one
INSERT INTO
    category (name, description, is_active)
VALUES
    ($ 1, $ 2, COALESCE($ 3, true)) RETURNING category_id,
    name,
    description,
    is_active,
    created_at,
    updated_at;

-- name: GetCategoryByID :one
SELECT
    category_id,
    name,
    description,
    is_active,
    created_at,
    updated_at
FROM
    category
WHERE
    category_id = $ 1;

-- name: GetCategoryByName :one
SELECT
    category_id,
    name,
    description,
    is_active,
    created_at,
    updated_at
FROM
    category
WHERE
    name = $ 1;

-- name: GetAllCategories :many
SELECT
    category_id,
    name,
    description,
    is_active,
    created_at,
    updated_at
FROM
    category
ORDER BY
    name;

-- name: UpdateCategory :exec
UPDATE
    category
SET
    name = $ 2,
    description = $ 3,
    is_active = $ 4,
    updated_at = current_timestamp
WHERE
    category_id = $ 1;

-- name: DeleteCategory :exec
DELETE FROM
    category
WHERE
    category_id = $ 1;