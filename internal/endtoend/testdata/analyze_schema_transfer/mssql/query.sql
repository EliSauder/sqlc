-- name: SalesOrders :many
SELECT id, total FROM sales.orders WHERE id = @id;

-- name: StagingItems :many
SELECT id FROM staging_items;
