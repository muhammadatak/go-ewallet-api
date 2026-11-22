-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transactions (
    id                   BIGSERIAL     PRIMARY KEY,
    account_id           BIGINT        NOT NULL REFERENCES accounts (id) ON DELETE CASCADE,
    sof_number           VARCHAR(50)   NOT NULL,
    dof_number           VARCHAR(50)   NOT NULL,
    transaction_type     CHAR(1)       NOT NULL CHECK (transaction_type IN ('D', 'C')),
    amount               NUMERIC(18,2) NOT NULL,
    transaction_datetime TIMESTAMPTZ   NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_transactions_account_id ON transactions (account_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;
-- +goose StatementEnd
