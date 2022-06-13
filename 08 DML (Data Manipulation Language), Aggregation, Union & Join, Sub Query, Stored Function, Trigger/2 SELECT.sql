-- a. Nama user gender M
SELECT * FROM users WHERE gender = 'M'

-- b. Tampilkan product id=3
SELECT * FROM products WHERE id=3

-- c. Tampilkan data pel create_at 7 hari ke belakang dan nama like %a%
SELECT * FROM users WHERE name LIKE '%a%' AND created_at > CURDATE() -INTERVAL 7 DAY AND NOW()

-- d. Hitung jumlah user gender Perempuan atau F (Female)
SELECT COUNT FROM users WHERE gender = 'F'

-- e. Tampilkan data pel dg urutan sesuai abjad
SELECT * FROM users ORDER BY name ASC

-- f. Tampilkan 5 data pada product
SELECT * FROM products LIMIT 5