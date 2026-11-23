-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS templates (
    code  VARCHAR(100) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    body  TEXT         NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS templates;
-- +goose StatementEnd
