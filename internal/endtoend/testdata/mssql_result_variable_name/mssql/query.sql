-- name: Sales2023 :many
SELECT SUM(amount) AS [2023] FROM sales;

-- name: TotalForAuthor :one
SELECT [total amount] = SUM(amount) FROM sales WHERE author_id = @author_id;

-- name: LeadingUnicodeDigit :many
SELECT amount AS [١total] FROM sales;

-- name: NoIdentifierCharacters :many
SELECT amount AS [!] FROM sales;
