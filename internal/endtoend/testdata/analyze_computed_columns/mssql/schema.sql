CREATE TABLE authors (
  id BIGINT IDENTITY(1,1) PRIMARY KEY,
  name NVARCHAR(100) NOT NULL,
  bio NVARCHAR(200) NULL,
  price INT NOT NULL,
  total AS (id * 2),
  doubled AS (price * 2) PERSISTED,
  doubled_nn AS (price * 2) PERSISTED NOT NULL,
  label AS (name + N'!'),
  full_label AS (CONCAT(name, bio)),
  bio_label AS (bio + N'?'),
  cost AS (CAST(price AS decimal(10,2))),
  c_isnull AS (ISNULL(bio, N'none'))
);

ALTER TABLE authors ADD extra AS (price + 1);
