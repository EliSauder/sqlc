-- name: RenamedColumn :many
SELECT id, w FROM t;

-- name: RenamedTable :many
SELECT id FROM renamed;

-- name: RenamedInSchema :many
SELECT b FROM s.y;
