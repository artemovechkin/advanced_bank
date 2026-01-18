-- +goose Up
-- +goose StatementBegin
ALTER TABLE accounts ADD COLUMN is_active BOOLEAN DEFAULT true;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE accounts DROP COLUMN is_active;
-- +goose StatementEnd
