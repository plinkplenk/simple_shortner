-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE users ALTER COLUMN id SET DEFAULT uuid_generate_v4();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE users ALTER COLUMN id DROP DEFAULT;
-- +goose StatementEnd
