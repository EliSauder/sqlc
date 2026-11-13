-- name: Ids :many
SELECT id FROM t WHERE name = '@not_a_param' /* @nor_this */ AND id > @min;

-- name: Sample :many
SELECT id FROM t WITH (INDEX(@ix));
