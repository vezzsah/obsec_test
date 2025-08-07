-- +goose Up
CREATE TABLE cpe_per_project (
    id TEXT UNIQUE DEFAULT (uuid()) NOT NULL,
    cpe TEXT NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    project_id TEXT NOT NULL,
    
    PRIMARY KEY(id),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE cpe_per_project;