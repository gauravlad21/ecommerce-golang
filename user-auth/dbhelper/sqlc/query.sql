-- name: InsertUser :one
INSERT INTO users(
    email, password, user_type
)VALUES(
    $1, $2, $3
)
RETURNING id;

-- name: GetUser :many
SELECT * FROM users
WHERE email=$1;

-- name: GetUserById :many
SELECT * FROM users
WHERE id=$1;