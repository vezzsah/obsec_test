-- name: StoreCVE :one
INSERT INTO cve_per_project (id, cve, descrip, created_at, updated_at, cpe, project)
VALUES (
    $1,
    $2,
    $3,
    $4, 
    $5,
    $6,
    $7
)
RETURNING *;

-- name: GetAllCVEByProject :many
SELECT * FROM cve_per_project WHERE project = $1 AND solved = false;

-- name: UpdateCVE :exec
UPDATE cve_per_project SET solved = $1 WHERE cve = $2 AND project = $3 AND CPE = $4;

-- name: GetProjectAndCPE :one
SELECT p.id as projectID, cpe.id as CPEID, cve.id as CVEID, cve.cve as CVEString FROM projects as p 
LEFT JOIN cpe_per_project cpe ON cpe.project_id = p.id
LEFT JOIN cve_per_project cve ON cve.project = p.id
WHERE p.project_name = $1 
AND p.creator = $2
AND cpe.cpe = $3
AND cve.cve = $4;