-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, email, hashedP)
VALUES (
    ?,
    ?,
    ?,
    ?
)
RETURNING id, created_at, updated_at, email;

-- name: CheckIfUserExistByEmail :one
SELECT EXISTS (Select 1 from users
where email = ?);

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ?;

-- name: DeleteAllUsers :exec
DELETE FROM users;