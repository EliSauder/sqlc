-- name: AuthorBooks :many
SELECT a.name, b.title FROM authors a JOIN books b ON b.author_id = a.id;

-- name: BookTitles :many
SELECT * FROM book_titles;
