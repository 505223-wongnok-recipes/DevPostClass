-- +goose Up
ALTER TABLE items ADD quantity integer NOT NULL;

-- +goose StatementBegin
SELECT
    'up SQL query';

-- +goose StatementEnd
-- +goose Down
ALTER TABLE items
DROP COLUMN quantity;

-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd