-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cities (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

INSERT INTO cities (name) VALUES ('Москва');
INSERT INTO cities (name) VALUES ('Санкт-Петербург');
INSERT INTO cities (name) VALUES ('Казань');   

CREATE INDEX IF NOT EXISTS idx_cities_name ON cities(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_cities_name;
DROP TABLE IF EXISTS cities;
-- +goose StatementEnd
