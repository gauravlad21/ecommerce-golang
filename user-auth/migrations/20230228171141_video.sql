-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(500) NOT NULL,
    user_type VARCHAR(500) NOT NULL,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
-- DROP TABLE video;
-- +goose StatementEnd
