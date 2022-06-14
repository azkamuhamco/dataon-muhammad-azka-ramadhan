// a. Delete data pada tabel product dg id=1
db.products.deleteOne( { id: 1 } );

// b. Delete data pada tabel product dg prod type id=1
db.products.deleteOne( { product_type_id: 1 } );
