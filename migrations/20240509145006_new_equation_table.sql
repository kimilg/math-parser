-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS equation
(
    id             BIGSERIAL PRIMARY KEY,
    value          text NOT NULL,
    category       text NOT NULL,
    cause          text NULL,
    effect         text NULL,
    created_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at     TIMESTAMP NULL
);
CREATE TABLE IF NOT EXISTS variable
(
    id             BIGSERIAL PRIMARY KEY,
    name           text NOT NULL,
    vcategory      text NOT NULL,
    UNIQUE         (name, vcategory),
    created_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at     TIMESTAMP NULL
);
CREATE TABLE IF NOT EXISTS equation_variable
(
    equation_id     BIGSERIAL NOT NULL REFERENCES equation(id) ON UPDATE CASCADE ON DELETE CASCADE,
    variable_id     BIGSERIAL NOT NULL REFERENCES variable(id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT unique_equation_variable UNIQUE (equation_id, variable_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS equation_variable;
DROP TABLE IF EXISTS equation;
DROP TABLE IF EXISTS variable;
-- +goose StatementEnd
