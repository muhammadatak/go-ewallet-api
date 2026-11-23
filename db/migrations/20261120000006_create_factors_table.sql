-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS factors (
    id      BIGSERIAL    PRIMARY KEY,
    user_id BIGINT       NOT NULL UNIQUE REFERENCES users (id) ON DELETE CASCADE,
    pin     VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS factors;
-- +goose StatementEnd
