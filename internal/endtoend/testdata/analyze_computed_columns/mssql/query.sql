-- name: GetComputed :one
SELECT total, doubled, doubled_nn, label, full_label, bio_label, cost, c_isnull, extra
FROM authors
WHERE id = @id;

-- name: ListAuthors :many
SELECT * FROM authors;

-- name: CreateAuthor :exec
INSERT INTO authors (name, bio, price) VALUES (@name, @bio, @price);
