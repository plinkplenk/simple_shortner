-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE urls ADD CONSTRAINT unique_urls_id UNIQUE (id);

ALTER TABLE users ADD CONSTRAINT unique_users_email UNIQUE (email);
ALTER TABLE users ADD CONSTRAINT unique_users_username UNIQUE (username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE urls DROP CONSTRAINT unique_urls_id;
ALTER TABLE users DROP CONSTRAINT unique_users_email;
ALTER TABLE users DROP CONSTRAINT unique_users_username;
-- +goose StatementEnd
