-- name: CreateProject :one
INSERT INTO projects (project_name, created_at, updated_at, creator)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: CheckIfProjectExistByName :one
SELECT EXISTS (Select 1 from projects
where project_name = $1);

-- name: CheckIfProjectExistByUserIdAndName :one
SELECT EXISTS (Select 1 from projects
where creator = $1 and project_name = $2);

-- name: GetProjectByNameAndCreator :one
SELECT * FROM projects WHERE project_name = $1 and creator = $2;

-- name: GetProjectById :one
SELECT * FROM projects WHERE id = $1;

-- name: GetProjectsByUser :many
SELECT * FROM projects WHERE creator = $1;