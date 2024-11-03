-- name: CreateRecipe :one
INSERT INTO recipies (id, created_at, updated_at, title, description, author_id) 
VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateRecipe :one
UPDATE recipies SET title = $1, description = $2 WHERE id = $3
RETURNING *;

-- name: DeleteRecipe :exec
DELETE FROM recipies WHERE (id=$1 AND author_id=$2);

-- name: GetUserRecipies :many
SELECT * FROM recipies WHERE (author_id = $1);

-- name: GetAllRecipies :many
SELECT recipies.*, users.name as author_name FROM recipies JOIN users ON recipies.author_id = users.id;
