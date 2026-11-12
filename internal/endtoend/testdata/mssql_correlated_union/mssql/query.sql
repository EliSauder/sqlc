-- name: Related :many
SELECT a.id, r.title
FROM authors a
CROSS APPLY (
    SELECT title FROM books WHERE author_id = a.id
    UNION ALL
    SELECT name FROM authors WHERE id = a.id
) r;
