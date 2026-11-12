-- name: Sales2023 :many
SELECT SUM(amount) AS [2023] FROM sales;

-- name: TotalForAuthor :one
SELECT [total amount] = SUM(amount) FROM sales WHERE author_id = @author_id;
