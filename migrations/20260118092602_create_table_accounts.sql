-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS accounts (
    email TEXT PRIMARY KEY,
    name TEXT,
    age int,
    balance float
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE accounts;
-- +goose StatementEnd
