CREATE TABLE t (id INT NOT NULL, v NVARCHAR(10));
CREATE TABLE [old name] (id INT NOT NULL);
GO
EXEC sp_rename 't.v', 'w', 'COLUMN';
GO
EXEC sp_rename @objname = N'dbo.[old name]', @newname = 'renamed';
GO
CREATE SCHEMA s;
GO
CREATE TABLE s.x (a INT);
GO
EXEC sp_rename 's.x.a', 'b', 'COLUMN';
GO
EXEC sp_rename 's.x', 'y';
GO
