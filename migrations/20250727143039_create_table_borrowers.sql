-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS borrowers (
    id          UUID        PRIMARY KEY,
    name        VARCHAR     NOT NULL,
    phone       VARCHAR     NOT NULL,
    created_at  TIMESTAMP   NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP   NOT NULL    DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS borrowers
-- +goose StatementEnd
