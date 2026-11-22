-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS accounts (
    id             BIGSERIAL     PRIMARY KEY,
    user_id        BIGINT        NOT NULL UNIQUE REFERENCES users (id) ON DELETE CASCADE,
    account_number VARCHAR(50)   NOT NULL UNIQUE,
    balance        NUMERIC(18,2) NOT NULL DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS accounts;
-- +goose StatementEnd
