-- +goose Up
CREATE TABLE recipies(
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    author_id UUID NOT NULL REFERENCES users(id)
);


-- +goose Down
DROP TABLE recipies;