CREATE TABLE authors (
    id BIGINT IDENTITY(1,1) PRIMARY KEY,
    name NVARCHAR(100) NOT NULL,
    bio NVARCHAR(MAX)
);

CREATE TABLE books (
    id BIGINT PRIMARY KEY,
    author_id BIGINT NOT NULL,
    title NVARCHAR(200) NOT NULL,
    pages INT
);
GO

CREATE VIEW author_names AS SELECT id, name FROM authors;
GO

CREATE VIEW dbo.author_stats (author_id, book_count, longest) AS
SELECT a.id, COUNT(b.id), MAX(b.pages)
FROM authors a
LEFT JOIN books b ON b.author_id = a.id
GROUP BY a.id;
GO

CREATE SCHEMA reporting;
GO

CREATE OR ALTER VIEW reporting.short_books AS
SELECT id, title, pages FROM books WHERE pages < 100;
GO

CREATE VIEW retired AS SELECT id FROM books;
GO

DROP VIEW IF EXISTS retired;
GO

ALTER VIEW author_names AS SELECT id, name, bio FROM authors;
