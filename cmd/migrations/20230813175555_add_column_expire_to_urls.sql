-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE urls ADD COLUMN expire INTEGER;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE urls DROP COLUMN expire;
-- +goose StatementEnd
