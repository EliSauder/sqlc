-- name: Styled :many
SELECT CONVERT(nvarchar(30), at, @style) AS styled, CONVERT(varchar(10), at, 23) AS fixed FROM t;
