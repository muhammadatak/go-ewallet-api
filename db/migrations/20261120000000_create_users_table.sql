-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id                BIGSERIAL    PRIMARY KEY,
    email             VARCHAR(255) NOT NULL,
    fullname          VARCHAR(255) NOT NULL,
    username          VARCHAR(100) NOT NULL UNIQUE,
    password          VARCHAR(255) NOT NULL,
    phone             VARCHAR(20)  NOT NULL,
    email_verified_at TIMESTAMPTZ  NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
