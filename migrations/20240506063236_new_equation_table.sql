-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS equations
(
    id             BIGSERIAL PRIMARY KEY,
    value          text NOT NULL,
    created_at     TIMESTAMP NOT NULL,
    updated_at     TIMESTAMP NOT NULL,
    deleted_at     TIMESTAMP NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS equations;
-- +goose StatementEnd
