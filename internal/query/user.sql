-- name: CreateUser :one
INSERT INTO users (
    firstname,
    lastname,
    password,
    email
) VALUES (
    $1,$2,$3,$4
) RETURNING *;

-- name: CreateUserRole :one
INSERT INTO user_roles (
    user_id,
    role_id,
    assigned_at
) VALUES (
    $1, $2, DEFAULT
) RETURNING *;


-- name: GetUser :one
SELECT * FROM users
WHERE email =$1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY user_id;


-- name: GetUsersWithRoles :many
SELECT 
    u.user_id,
    u.firstname,
    u.lastname,
    u.password,
    u.email,
    r.role_id,
    r.name AS role_name,
    r.description AS role_description
FROM 
    users u
INNER JOIN 
    user_roles ur ON u.user_id = ur.user_id
INNER JOIN 
    roles r ON ur.role_id = r.role_id
ORDER BY 
    u.user_id;


-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE user_id =$1
FOR UPDATE;

-- name: UpdatUser :exec
UPDATE users SET 
firstname = $2,
lastname = $3,
password = $4,
email = $5
WHERE user_id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1;


-- name: UpdatPassword :one
UPDATE users SET password = $2
WHERE email = $1
RETURNING *;

