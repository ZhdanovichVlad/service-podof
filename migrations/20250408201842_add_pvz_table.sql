-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pvzs (
    id UUID PRIMARY KEY,
    registration_date TIMESTAMP NOT NULL,
    city VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (city) REFERENCES cities(name)
);  

CREATE INDEX IF NOT EXISTS idx_pvzs_city ON pvzs(city);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_pvzs_city;
DROP TABLE IF EXISTS pvzs;
-- +goose StatementEnd
