-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS login_log (
    id            BIGSERIAL        PRIMARY KEY,
    user_id       BIGINT           NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    is_authorized BOOLEAN          NOT NULL DEFAULT FALSE,
    ip_address    VARCHAR(45)      NOT NULL,
    timezone      VARCHAR(100)     NOT NULL DEFAULT '',
    lat           DOUBLE PRECISION NOT NULL DEFAULT 0,
    lon           DOUBLE PRECISION NOT NULL DEFAULT 0,
    access_time   TIMESTAMPTZ      NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_login_log_user_id_is_authorized ON login_log (user_id, is_authorized, id DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS login_log;
-- +goose StatementEnd
