-- +goose Up
CREATE TABLE auth (
    id SERIAL PRIMARY KEY,
    title TEXT,
    tz TEXT
);

-- +goose Down
DROP TABLE auth;
