-- name: TopPercent :many
SELECT TOP (@pct) PERCENT id, name FROM authors ORDER BY id;
