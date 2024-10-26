-- name: CreateRecipe :one
INSERT INTO recipies (id, created_at, updated_at, title, description, author_id) 
VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;