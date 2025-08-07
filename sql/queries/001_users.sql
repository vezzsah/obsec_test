-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, email, hashedP)
VALUES (
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING id, created_at, updated_at, email;

-- name: CheckIfUserExistByEmail :one
SELECT EXISTS (Select 1 from users
where email = $1);

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: DeleteAllUsers :exec
DELETE FROM users;