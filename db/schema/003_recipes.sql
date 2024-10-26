-- +goose Up
CREATE TABLE recipies(
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    author_id UUID REFERENCES users(id)
);


-- +goose Down
DROP TABLE recipies;