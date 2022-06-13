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
INSERT IGNORE INTO products VALUES (1, 1, 3, 'XL-PULSA-50', 50000, 'XL Pulsa 50000', 1, NOW(), NOW());
INSERT IGNORE INTO products VALUES (2, 1, 3, 'XL-PULSA-100', 100000, 'XL Pulsa 100000', 1, NOW(), NOW());

-- Insert 3 product dg product_types.id = 2 && operators.id = 1
INSERT IGNORE INTO products VALUES (3, 2, 1, 'TSEL-DATA-5GB', 40000, 'Telkomsel Data Reguler 5 GB', 1, NOW(), NOW());
INSERT IGNORE INTO products VALUES (4, 2, 1, 'TSEL-DATA-10GB', 75000, 'Telkomsel Data Reguler 10 GB', 1, NOW(), NOW());
INSERT IGNORE INTO products VALUES (5, 2, 1, 'TSEL-DATA-15GB', 120000, 'Telkomsel Data Reguler 15 GB', 1, NOW(), NOW());

-- Insert 3 product dg product_types.id = 3 && operators.id = 4
INSERT IGNORE INTO products VALUES (6, 3, 4, 'SMART-KOM-TEL20', 20000, 'Smartfren Bebas Bicara 20 Menit', 1, NOW(), NOW());
INSERT IGNORE INTO products VALUES (7, 3, 4, 'SMART-KOM-TEL50', 45000, 'Smartfren Bebas Bicara 50 Menit', 1, NOW(), NOW());
INSERT IGNORE INTO products VALUES (8, 3, 4, 'SMART-KOM-TEL90', 80000, 'Smartfren Bebas Bicara 90 Menit', 1, NOW(), NOW());

-- Insert product description setiap product
INSERT IGNORE INTO product_descriptions VALUES (1, 'XL Pulsa 50000 masa aktif 60 hari', NOW(), NOW());
INSERT IGNORE INTO product_descriptions VALUES (2, 'XL Pulsa 100000 masa aktif 90 hari', NOW(), NOW());
INSERT IGNORE INTO product_descriptions VALUES (3, 'Telkomsel Data Reguler 5 GB masa aktif 10 hari', NOW(), NOW());
INSERT IGNORE INTO product_descriptions VALUES (4, 'Telkomsel Data Reguler 10 GB masa aktif 20 hari', NOW(), NOW());
INSERT IGNORE INTO product_descriptions VALUES (5, 'Telkomsel Data Reguler 15 GB masa aktif 30 hari', NOW(), NOW());
INSERT IGNORE INTO product_descriptions VALUES (6, 'Smartfren Bebas Bicara 20 Menit masa aktif 1 hari', NOW(), NOW());
INSERT IGNORE INTO product_descriptions VALUES (7, 'Smartfren Bebas Bicara 50 Menit masa aktif 3 hari', NOW(), NOW());
INSERT IGNORE INTO product_descriptions VALUES (8, 'Smartfren Bebas Bicara 90 Menit masa aktif 6 hari', NOW(), NOW());

-- Insert 3 payment methods
INSERT IGNORE INTO payment_methods VALUES (1, 'Tunai', 1, NOW(), NOW());
INSERT IGNORE INTO payment_methods VALUES (2, 'Debit', 1, NOW(), NOW());
INSERT IGNORE INTO payment_methods VALUES (3, 'E-pay', 1, NOW(), NOW());

-- Insert 5 user
INSERT IGNORE INTO users VALUES (1, 'Karim', 1, '1994-07-05', 'M', NOW(), NOW());
INSERT IGNORE INTO users VALUES (2, 'Karina', 1, '1992-07-15', 'F', NOW(), NOW());
INSERT IGNORE INTO users VALUES (3, 'Yoda', 1, '2020-04-04', 'M', NOW(), NOW());
INSERT IGNORE INTO users VALUES (4, 'Rahma', 1, '1989-10-09', 'F', NOW(), NOW());
INSERT IGNORE INTO users VALUES (5, 'Azka', 1, '1991-01-07', 'M', NOW(), NOW());

-- Insert 3 transaksi masing2 user
INSERT IGNORE INTO transactions VALUES (1, 5, 2, 1, 3, 145000, NOW(), NOW());
INSERT IGNORE INTO transactions VALUES (2, 4, 3, 1, 3, 185000, NOW(), NOW());
INSERT IGNORE INTO transactions VALUES (3, 1, 1, 1, 3, 300000, NOW(), NOW());

-- Insert 3 product di masing2 transaksi
INSERT IGNORE INTO transaction_details VALUES (1, 1, 'Terbayar', 1, 50000, NOW(), NOW());
INSERT IGNORE INTO transaction_details VALUES (1, 4, 'Terbayar', 1, 75000, NOW(), NOW());
INSERT IGNORE INTO transaction_details VALUES (1, 6, 'Terbayar', 1, 20000, NOW(), NOW());

INSERT IGNORE INTO transaction_details VALUES (2, 2, 'Terbayar', 1, 100000, NOW(), NOW());
INSERT IGNORE INTO transaction_details VALUES (2, 3, 'Terbayar', 1,  40000, NOW(), NOW());
INSERT IGNORE INTO transaction_details VALUES (2, 7, 'Terbayar', 1,  45000, NOW(), NOW());

INSERT IGNORE INTO transaction_details VALUES (3, 2, 'Terbayar', 1, 100000, NOW(), NOW());
INSERT IGNORE INTO transaction_details VALUES (3, 5, 'Terbayar', 1, 120000, NOW(), NOW());
INSERT IGNORE INTO transaction_details VALUES (3, 8, 'Terbayar', 1,  80000, NOW(), NOW());