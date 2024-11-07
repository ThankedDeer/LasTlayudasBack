-- name: CreateProvider :one
INSERT INTO provider (name, phone, email, address)
VALUES ($1, $2, $3, $4)
RETURNING provider_id, name, phone, email, address, created_at, updated_at;

-- name: GetProviderByID :one
SELECT provider_id, name, phone, email, address, created_at, updated_at
FROM provider
WHERE provider_id = $1;

-- name: GetProviderByEmail :one
SELECT provider_id, name, phone, email, address, created_at, updated_at
FROM provider
WHERE email = $1;

-- name: GetAllProviders :many
SELECT provider_id, name, phone, email, address, created_at, updated_at
FROM provider
ORDER BY name;

-- name: UpdateProvider :exec
UPDATE provider
SET name = $2,
    phone = $3,
    email = $4,
    address = $5,
    updated_at = current_timestamp
WHERE provider_id = $1;

-- name: DeleteProvider :exec
DELETE FROM provider
WHERE provider_id = $1;
