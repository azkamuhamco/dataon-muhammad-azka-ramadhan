-- 1 Gabungkan data transaksi dari user id 1 dan user id 2
SELECT * FROM transactions t INNER JOIN users u ON t.user_id = u.id WHERE u.id = 1 OR u.id = 2

-- 2 Tampilkan jumlah harga transaksi user id 1
SELECT t.total_price FROM transactions t INNER JOIN users u ON t.user_id = u.id WHERE u.id = 1

-- 3 Tampilkan total transaksi dengan product type 2
SELECT COUNT(*) FROM transaction_details td INNER JOIN products p ON td.product_id = p.id WHERE p.id = 2

-- 4 Tampilkan semua field prod dan field name table product type yg saling berhub
SELECT * FROM products p INNER JOIN product_types pt ON p.product_type_id = pt.id

-- 5 Tampilkan semua field table transaction, field name table product dan field name table user
SELECT * FROM transaction_details td 
    INNER JOIN transactions t ON td.transaction_id = t.id
    INNER JOIN products p ON td.product_id = p.id
    INNER JOIN users u ON t.user_id = u.id;

-- 6 Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan trx id yang dimaksud
DELETE FROM transaction_details WHERE transaction_id IN (SELECT id FROM transactions WHERE id=@Id);
DELETE FROM transactions WHERE id=@Id;

-- 7 Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yg dihapus
DELETE FROM transaction_details WHERE transaction_id IN
    (SELECT id FROM transactions WHERE total_qty IN 
        (SELECT total_qty FROM transactions WHERE updated_at IN 
            (SELECT MAX(updated_at) FROM transactions)
        )
    );

-- 8 Tampilkan data products yang tidak pernah ada di tabel transaction_details dg sub-query
SELECT * FROM products WHERE id NOT IN 
    ();