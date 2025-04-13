-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
INSERT INTO roles (name) VALUES ('client');
INSERT INTO roles (name) VALUES ('moderator');
INSERT INTO roles (name) VALUES ('pvz_employee');

CREATE INDEX IF NOT EXISTS idx_roles_name ON roles(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_roles_name;
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd
