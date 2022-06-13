-- 1 Gabungkan data transaksi dari user id 1 dan user id 2
SELECT * FROM transactions t INNER JOIN users u ON t.user_id = u.id WHERE u.id = 1 OR u.id = 2

-- 2 Tampilkan jumlah harga transaksi user id 1


-- 3 Tampilkan total transaksi dengan product type 2


-- 4 Tampilkan semua field prod dan field name table product type yg saling berhub


-- 5 Tampilkan semua field table transaction, field name table product dan field name table user


-- 6 Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan trx id yang dimaksud


-- 7 Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yg dihapus


-- 8 Tampilkan data products yang tidak pernah ada di tabel transaction_details dg sub-query
