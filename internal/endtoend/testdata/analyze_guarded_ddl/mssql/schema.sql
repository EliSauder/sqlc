IF OBJECT_ID('dbo.authors', 'U') IS NULL
    CREATE TABLE authors (id INT NOT NULL PRIMARY KEY, name NVARCHAR(100) NOT NULL);
GO
IF NOT EXISTS (SELECT 1 FROM sys.tables WHERE name = 'books')
BEGIN
    CREATE TABLE books (id INT NOT NULL, author_id INT NOT NULL);
    INSERT INTO books (id, author_id) VALUES (1, 1);
END
GO
IF COL_LENGTH('dbo.books', 'title') IS NULL
    ALTER TABLE books ADD title NVARCHAR(200) NULL;
GO
IF OBJECT_ID('dbo.book_titles', 'V') IS NOT NULL
    DROP VIEW book_titles;
GO
CREATE VIEW book_titles AS SELECT id, title FROM books;
