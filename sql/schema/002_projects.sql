-- +goose Up
CREATE TABLE projects (
    id TEXT UNIQUE DEFAULT (uuid7()),
    project_name TEXT NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    creator TEXT NOT NULL,
    
    PRIMARY KEY(id),
    FOREIGN KEY (creator) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE projects;