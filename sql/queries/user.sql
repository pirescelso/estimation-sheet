-- name: InsertUser :exec
INSERT INTO
    users (
        user_id,
        email,
        user_name,
        name,
        user_type,
        created_at
    )
VALUES ($1, $2, $3, $4, $5, $6);

-- name: FindUserById :one
SELECT * FROM users WHERE user_id = $1;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: FindAllUsers :many
SELECT * FROM users;

-- name: UpdateUser :exec
UPDATE users
SET
    email = $2,
    user_name = $3,
    name = $4,
    user_type = $5,
    updated_at = $6
WHERE
    user_id = $1;

-- name: DeleteUser :one
DELETE FROM users WHERE user_id = $1 RETURNING *;