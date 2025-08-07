-- +goose Up
CREATE TABLE users (
    id TEXT UNIQUE DEFAULT(uuid()),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    email TEXT UNIQUE NOT NULL,
    hashedP TEXT NOT NULL,
    
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE users;