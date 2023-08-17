-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE EXTENSION if not exists "uuid-ossp";
CREATE TABLE users (
    id uuid CONSTRAINT "users_id_pk" PRIMARY KEY,
    email TEXT NOT NULL,
    username TEXT NOT NULL,
    password TEXT NOT NULL
);
CREATE TABLE urls (
  id text constraint "urls_id_pk" primary key,
  original text,
  user_id uuid,
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE urls;
DROP TABLE users;
-- +goose StatementEnd
