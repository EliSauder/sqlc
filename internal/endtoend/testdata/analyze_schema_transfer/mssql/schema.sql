CREATE TABLE orders (id INT NOT NULL, total DECIMAL(10,2) NOT NULL);
CREATE TABLE staging_items (id INT NOT NULL);
GO
CREATE SCHEMA sales;
GO
CREATE SCHEMA archive;
GO
ALTER SCHEMA sales TRANSFER dbo.orders;
ALTER SCHEMA archive TRANSFER OBJECT::staging_items;
ALTER SCHEMA dbo TRANSFER archive.staging_items;
