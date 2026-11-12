-- name: CopyBooks :copyfrom
INSERT INTO books (id, title) VALUES (@id, @title);
