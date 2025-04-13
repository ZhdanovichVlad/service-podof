-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS receptions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    date_time TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
    pvz_id UUID NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (pvz_id) REFERENCES pvzs(id)
);  

CREATE INDEX IF NOT EXISTS idx_receptions_pvz_id ON receptions(pvz_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_recep  tions_pvz_id;
DROP TABLE IF EXISTS reception s;
-- +goose StatementEnd
