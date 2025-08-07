-- name: StoreCPE :one
INSERT INTO cpe_per_project (cpe, created_at, updated_at, project_id)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: CheckIfCPEExistByProjectName :one
SELECT EXISTS (Select 1 from cpe_per_project
where cpe = $1 and project_id = $2);

-- name: GetAllCPEByProject :many
SELECT * FROM cpe_per_project WHERE project_id = $1;

-- name: GetCPEById :one
SELECT * FROM cpe_per_project WHERE id = $1;