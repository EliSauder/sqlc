-- name: Sales2023 :many
SELECT pv.[2023] FROM sales PIVOT (SUM(amount) FOR year IN ([2023])) pv;

-- name: TotalForAuthor :one
SELECT [total amount] = SUM(amount) FROM sales WHERE author_id = @author_id;

-- name: LeadingUnicodeDigit :many
SELECT amount AS [١total] FROM sales;

-- name: NoIdentifierCharacters :many
SELECT amount AS [!] FROM sales;
