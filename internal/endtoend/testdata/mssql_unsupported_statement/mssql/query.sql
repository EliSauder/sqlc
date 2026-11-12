-- name: ListAuthors :many
SELECT id, name FROM authors;

-- name: UpsertAuthor :exec
MERGE authors AS t
USING (SELECT @id AS id, @name AS name) s ON t.id = s.id
WHEN MATCHED THEN UPDATE SET name = s.name
WHEN NOT MATCHED THEN INSERT (name) VALUES (s.name);
