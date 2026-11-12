-- name: GetComputed :one
SELECT total, doubled, doubled_nn, label, bio_label, cost, extra
FROM authors
WHERE id = @id;

-- name: ListAuthors :many
SELECT * FROM authors;

-- name: CreateAuthor :exec
INSERT INTO authors (name, bio, price) VALUES (@name, @bio, @price);
