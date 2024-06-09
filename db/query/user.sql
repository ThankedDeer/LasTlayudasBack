-- name: CreateUser :one
INSERT INTO users (
    firstname,
    lastname,
    password,
    email
) VALUES (
    $1,$2,$3,$4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email =$1;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE email =$1
FOR UPDATE;

-- name: UpdatPassword :one
UPDATE users SET password = $2
WHERE email = $1
RETURNING *;

