-- +goose Up
CREATE TABLE users (
    id TEXT UNIQUE DEFAULT(uuid7()),
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    hashedP TEXT NOT NULL,
    
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE users;