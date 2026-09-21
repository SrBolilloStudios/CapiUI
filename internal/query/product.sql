--name: products.create_table
CREATE TABLE IF NOT EXISTS Products(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sku TEXT NOT NULL,
    name TEXT NOT NULL,
    price REAL NOT NULL,
    description TEXT,
    img TEXT,
    stock INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    FOREIGN KEY (category_id) REFERENCES CategoryProduct(id)
);

--name: products.drop_table
DROP TABLE IF EXISTS Products;

--name: products.insert
INSERT INTO Products (sku, name, price, description, img, stock, category_id) VALUES %s

--name: products.get_by_ids
SELECT p.id, p.sku, p.name, p.price, p.description, p.img, p.stock, p.category_id, c.name
FROM Products p
         LEFT JOIN CategoryProduct c ON p.category_id = c.id
WHERE p.id IN (%s)

--name: products.list
SELECT p.id, p.sku, p.name, p.price, p.description, p.img, p.stock, p.category_id, c.name
FROM Products p
         LEFT JOIN CategoryProduct c ON p.category_id = c.id
LIMIT ? OFFSET ?

--name: products.delete_by_ids
DELETE FROM Products WHERE id IN (%s)
