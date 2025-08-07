-- name: StoreCPE :one
INSERT INTO cpe_per_project (cpe, created_at, updated_at, project_id)
VALUES (
    ?,
    ?,
    ?,
    ?
)
RETURNING *;

-- name: CheckIfCPEExistByProjectName :one
SELECT EXISTS (Select 1 from cpe_per_project
where cpe = ? and project_id = ?);

-- name: GetAllCPEByProject :many
SELECT * FROM cpe_per_project WHERE project_id = ?;

-- name: GetCPEById :one
SELECT * FROM cpe_per_project WHERE id = ?;