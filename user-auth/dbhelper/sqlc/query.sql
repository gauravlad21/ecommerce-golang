-- name: InsertUser :one
INSERT INTO users(
    email, password
)VALUES(
    $1, $2
)
RETURNING id;

-- name: GetUser :many
SELECT * FROM users
WHERE email=$1;

-- name: GetUserById :many
SELECT * FROM users
WHERE id=$1;