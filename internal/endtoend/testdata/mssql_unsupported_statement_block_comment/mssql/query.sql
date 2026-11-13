/* name: ListAuthors :many */
SELECT id, name FROM authors;

/* name: DeleteAll :exec */
MERGE authors AS t
USING (SELECT @id AS id) s ON t.id = s.id
WHEN MATCHED THEN DELETE;
