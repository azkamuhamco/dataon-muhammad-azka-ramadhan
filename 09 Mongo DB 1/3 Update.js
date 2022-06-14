// a. Ubah data product id 1 jadi "product dummy"
db.products.updateOne( { id: 1 }, { $set: { name: 'product dummy' } } );

// b. Ubah qty=3 pada transaction_details dg product id 1
db.transaction_details.updateMany( { product_id: 1 }, { $set: { qty: 3 } } );
