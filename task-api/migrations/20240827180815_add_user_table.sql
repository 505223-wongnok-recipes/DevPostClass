-- +goose Up
CREATE TABLE
    users (
        id bigserial NOT NULL,
        username varchar(50) NOT NULL UNIQUE,
        password varchar(100) NOT NULL,
        PRIMARY KEY (id)
    );
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS users;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd