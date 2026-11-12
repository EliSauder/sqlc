-- name: AuthorNames :many
SELECT * FROM author_names WHERE id = @id;

-- name: AuthorStats :many
SELECT s.author_id, s.book_count, s.longest, n.name
FROM author_stats s
JOIN author_names n ON n.id = s.author_id;

-- name: ShortBooks :many
SELECT * FROM reporting.short_books;
