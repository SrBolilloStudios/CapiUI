--name: categories.create_table
CREATE TABLE IF NOT EXISTS CategoryProduct(
id  INTEGER PRIMARY KEY AUTOINCREMENT ,
name TEXT NOT NULL
);

--name: categories.drop_table
DROP TABLE IF EXISTS CategoryProduct;

--name: categories.insert
INSERT INTO CategoryProduct (name) VALUES %s

--name: categories.get_by_ids
SELECT id, name FROM CategoryProduct WHERE id IN (%s)

--name: categories.list
SELECT id, name FROM CategoryProduct LIMIT ? OFFSET ?

--name: categories.delete_by_ids
DELETE FROM CategoryProduct WHERE id IN (%s)
