-- name: StoreCVE :one
INSERT INTO cve_per_project (id, cve, descrip, created_at, updated_at, cpe, project)
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
)
RETURNING *;

-- name: GetAllCVEByProject :many
SELECT * FROM cve_per_project WHERE project = ? AND solved = false;

-- name: UpdateCVE :exec
UPDATE cve_per_project SET solved = ? WHERE cve = ? AND project = ? AND CPE = ?;

-- name: GetProjectAndCPE :one
SELECT p.id as projectID, cpe.id as CPEID, cve.id as CVEID, cve.cve as CVEString FROM projects as p 
LEFT JOIN cpe_per_project cpe ON cpe.project_id = p.id
LEFT JOIN cve_per_project cve ON cve.project = p.id
WHERE p.project_name = ?
AND p.creator = ?
AND cpe.cpe = ?
AND cve.cve = ?;