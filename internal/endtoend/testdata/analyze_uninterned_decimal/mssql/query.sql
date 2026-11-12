-- name: Negated :many
SELECT -CAST(i AS DECIMAL) AS nd, i - CAST(f AS DECIMAL(5,1)) AS diff, d * CAST(i AS DECIMAL(12,3)) AS product FROM p2;
