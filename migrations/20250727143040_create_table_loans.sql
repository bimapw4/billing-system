-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS loans (
    id              UUID            PRIMARY KEY,
    borrower_id     UUID            NOT NULL REFERENCES borrowers(id) ON DELETE CASCADE,
    principal       INT             NOT NULL,                          
    interest_rate   INT             NOT NULL,            
    total_weeks     INT             NOT NULL,
    start_date      DATE            NOT NULL,
    created_at      TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS loans
-- +goose StatementEnd
