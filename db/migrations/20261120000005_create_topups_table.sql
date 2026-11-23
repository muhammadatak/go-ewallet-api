-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS topups (
    id       VARCHAR(36)   PRIMARY KEY,
    user_id  BIGINT        NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    status   SMALLINT      NOT NULL DEFAULT 0,
    amount   NUMERIC(18,2) NOT NULL,
    snap_url VARCHAR(255)  NOT NULL DEFAULT ''
);

CREATE INDEX IF NOT EXISTS idx_topups_user_id ON topups (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS topups;
-- +goose StatementEnd
