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
INSERT IGNORE INTO products VALUES (1, 1, 3, 'XL-PULSA-50', 'XL Pulsa 50000', 1, NOW(), NOW());
INSERT IGNORE INTO products VALUES (2, 1, 3, 'XL-PULSA-100', 'XL Pulsa 100000', 1, NOW(), NOW());

-- Insert 3 product dg product_types.id = 2 && operators.id = 1
INSERT IGNORE INTO products VALUES (3, 2, 1, 'TSEL-DATA-5GB', 'Telkomsel Data Reguler 5 GB', 1, NOW(), NOW());
INSERT IGNORE INTO products VALUES (4, 2, 1, 'TSEL-DATA-10GB', 'Telkomsel Data Reguler 10 GB', 1, NOW(), NOW());
INSERT IGNORE INTO products VALUES (5, 2, 1, 'TSEL-DATA-15GB', 'Telkomsel Data Reguler 15 GB', 1, NOW(), NOW());

-- Insert 3 product dg product_types.id = 3 && operators.id = 4
INSERT IGNORE INTO products VALUES (6, 3, 4, 'SMART-KOM-TEL20', 'Smartfren Bebas Bicara 20 Menit', 1, NOW(), NOW());
INSERT IGNORE INTO products VALUES (7, 3, 4, 'SMART-KOM-TEL50', 'Smartfren Bebas Bicara 50 Menit', 1, NOW(), NOW());
INSERT IGNORE INTO products VALUES (8, 3, 4, 'SMART-KOM-TEL90', 'Smartfren Bebas Bicara 90 Menit', 1, NOW(), NOW());

-- Insert product description setiap product
INSERT IGNORE INTO product_descriptions VALUES (1, 'XL Pulsa 50000 masa aktif 60 hari', NOW(), NOW())
INSERT IGNORE INTO product_descriptions VALUES (2, 'XL Pulsa 100000 masa aktif 90 hari', NOW(), NOW())
INSERT IGNORE INTO product_descriptions VALUES (3, 'Telkomsel Data Reguler 5 GB masa aktif 10 hari', NOW(), NOW())
INSERT IGNORE INTO product_descriptions VALUES (4, 'Telkomsel Data Reguler 10 GB masa aktif 20 hari', NOW(), NOW())
INSERT IGNORE INTO product_descriptions VALUES (5, 'Telkomsel Data Reguler 15 GB masa aktif 30 hari', NOW(), NOW())
INSERT IGNORE INTO product_descriptions VALUES (6, 'Smartfren Bebas Bicara 20 Menit masa aktif 1 hari', NOW(), NOW())
INSERT IGNORE INTO product_descriptions VALUES (7, 'Smartfren Bebas Bicara 50 Menit masa aktif 3 hari', NOW(), NOW())
INSERT IGNORE INTO product_descriptions VALUES (8, 'Smartfren Bebas Bicara 90 Menit masa aktif 6 hari', NOW(), NOW())

-- Insert 3 payment methods


-- Insert 5 user


-- Insert 3 transaksi masing2 user


-- Insert 3 product di masing2 transaksi
