-- name: CreateProject :one
INSERT INTO projects (project_name, created_at, updated_at, creator)
VALUES (
    ?,
    ?,
    ?,
    ?
)
RETURNING *;

-- name: CheckIfProjectExistByName :one
SELECT EXISTS (Select 1 from projects
where project_name = ?);

-- name: CheckIfProjectExistByUserIdAndName :one
SELECT EXISTS (Select 1 from projects
where creator = ? and project_name = ?);

-- name: GetProjectByNameAndCreator :one
SELECT * FROM projects WHERE project_name = ? and creator = ?;

-- name: GetProjectById :one
SELECT * FROM projects WHERE id = ?;

-- name: GetProjectsByUser :many
SELECT * FROM projects WHERE creator = ?;