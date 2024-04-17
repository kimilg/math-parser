-- +goose up
CREATE TABLE IF NOT EXISTS classifications
(
    id             UUID PRIMARY KEY,
    created_at     TIMESTAMP NOT NULL,
    updated_at     TIMESTAMP NOT NULL,
    deleted_at     TIMESTAMP NULL
);

-- +goose down
DROP TABLE IF EXISTS classifications;