-- name: WriterPeople :many
SELECT w.id, w.name FROM writers w JOIN app.people p ON p.id = w.id WHERE w.id = @id;

-- name: AllWriters :many
SELECT * FROM writers;
