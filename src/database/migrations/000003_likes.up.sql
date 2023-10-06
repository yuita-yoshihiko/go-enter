CREATE TABLE likes (
    id serial PRIMARY KEY,
    user_id integer NOT NULL REFERENCES users(id),
    book_id integer NOT NULL REFERENCES books(id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);