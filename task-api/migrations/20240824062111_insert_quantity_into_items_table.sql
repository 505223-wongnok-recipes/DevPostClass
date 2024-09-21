-- +goose Up
UPDATE items
SET
    quantity = 10;

-- +goose StatementBegin
SELECT
    'up SQL query';

-- +goose StatementEnd
-- +goose Down
UPDATE items
SET
    quantity = 0;

-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd