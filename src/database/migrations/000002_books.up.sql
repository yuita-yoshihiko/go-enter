CREATE TABLE books (
    id varchar(36) PRIMARY KEY,
    title varchar NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);