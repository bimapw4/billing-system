-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payments (
    id          UUID            PRIMARY KEY,
    loan_id     uuid            NOT NULL REFERENCES loans(id) ON DELETE CASCADE,
    week        INT             NOT NULL,
    paid        INT             NOT NULL DEFAULT 0,
    is_paid     BOOLEAN         NOT NULL DEFAULT FALSE,
    paid_at     TIMESTAMP,
    created_at  TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payments
-- +goose StatementEnd
