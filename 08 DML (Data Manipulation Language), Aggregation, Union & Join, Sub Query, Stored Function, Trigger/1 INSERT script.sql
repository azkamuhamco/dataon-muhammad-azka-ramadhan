-- Insert 5 operator
INSERT IGNORE INTO operators VALUES (1, 'Telkomsel', NOW(), NOW());
INSERT IGNORE INTO operators VALUES (2, 'Indosat Ooredoo', NOW(), NOW());
INSERT IGNORE INTO operators VALUES (3, 'XL Axiata', NOW(), NOW());
INSERT IGNORE INTO operators VALUES (4, 'Smartfren', NOW(), NOW());
INSERT IGNORE INTO operators VALUES (5, 'Three', NOW(), NOW());

-- Insert 3 product type
INSERT IGNORE INTO product_types VALUES (1, 'Pulsa', NOW(), NOW());
INSERT IGNORE INTO product_types VALUES (2, 'Paket Data', NOW(), NOW());
INSERT IGNORE INTO product_types VALUES (3, 'Paket Komunikasi', NOW(), NOW());

-- Insert 2 product dg product_types.id = 1 && operators.id = 3


-- Insert 3 product dg product_types.id = 2 && operators.id = 1


-- Insert 3 product dg product_types.id = 3 && operators.id = 4


-- Insert product description setiap product


-- Insert 3 payment methods


-- Insert 5 user


-- Insert 3 transaksi masing2 user


-- Insert 3 product di masing2 transaksi
