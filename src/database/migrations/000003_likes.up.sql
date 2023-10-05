CREATE TABLE likes (
    id varchar(36) PRIMARY KEY,
    user_id varchar NOT NULL REFERENCES users(id),
    book_id varchar NOT NULL REFERENCES books(id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);