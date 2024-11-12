-- name: CreateUser :one
INSERT INTO "user" (role_id, first_name, last_name, email, password, is_active)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    COALESCE(sqlc.arg('is_active'), true)
)
RETURNING user_id, role_id, first_name, last_name, email, is_active, created_at, updated_at;


-- name: GetUserByID :one
SELECT user_id, role_id, first_name, last_name, email, is_active, created_at, updated_at
FROM "user"
WHERE user_id = $1;

-- name: GetUserByEmail :one
SELECT user_id, role_id, first_name, last_name, email, is_active, created_at, updated_at
FROM "user"
WHERE email = $1;

-- name: GetAllUsers :many
SELECT user_id, role_id, first_name, last_name, email, is_active, created_at, updated_at
FROM "user"
ORDER BY last_name, first_name;

-- name: UpdateUser :exec
UPDATE "user"
SET role_id = $2,
    first_name = $3,
    last_name = $4,
    email = $5,
    password = $6,
    is_active = $7,
    updated_at = current_timestamp
WHERE user_id = $1;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE user_id = $1;
