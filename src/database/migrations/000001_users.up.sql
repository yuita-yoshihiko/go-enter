CREATE TABLE users (
    id serial PRIMARY KEY,
    name varchar NOT NULL,
    password varchar NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);