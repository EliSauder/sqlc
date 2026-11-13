-- name: GetAmount :one
SELECT amount FROM balances WHERE amount = @amount;

-- name: GetBalance :one
SELECT * FROM balances WHERE token = @token;

-- name: SetOptional :exec
UPDATE balances SET optional_amount = @optional_amount WHERE token = @token;
