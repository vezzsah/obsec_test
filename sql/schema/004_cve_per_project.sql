-- +goose Up
CREATE TABLE cve_per_project (
    id TEXT UNIQUE NOT NULL,
    cve TEXT NOT NULL,
    descrip TEXT NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    cpe TEXT NOT NULL,
    project TEXT NOT NULL,
    solved BOOLEAN NOT NULL DEFAULT false,
    
    PRIMARY KEY(id),
    FOREIGN KEY (cpe) REFERENCES projects(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE cve_per_cpe;