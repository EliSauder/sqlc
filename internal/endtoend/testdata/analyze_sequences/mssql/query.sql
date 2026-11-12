-- name: NextIDs :one
SELECT NEXT VALUE FOR order_seq AS a, NEXT VALUE FOR dbo.small_seq AS b, NEXT VALUE FOR billing.invoice_seq AS c;

-- name: InsertOrder :exec
INSERT INTO orders (id, n) VALUES (NEXT VALUE FOR order_seq, @n);
