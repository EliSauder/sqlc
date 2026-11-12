CREATE TABLE authors (id INT NOT NULL, name NVARCHAR(50) NULL);
GO
CREATE SCHEMA app;
GO
CREATE SYNONYM writers FOR dbo.authors;
CREATE SYNONYM app.people FOR authors;
CREATE SYNONYM later FOR dbo.not_yet_created;
CREATE SYNONYM remote FOR otherdb.dbo.authors;
