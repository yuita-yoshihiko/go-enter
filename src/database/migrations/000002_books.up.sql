CREATE TABLE books (
    id serial PRIMARY KEY,
    title varchar NOT NULL,
    genre varchar NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);