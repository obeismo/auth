-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_role AS ENUM ('UNKNOWN', 'ADMIN', 'USER');
create table if not exists users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role user_role NOT NULL DEFAULT 'UNKNOWN',
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type if exists user_role
drop table if exists users;
-- +goose StatementEnd
