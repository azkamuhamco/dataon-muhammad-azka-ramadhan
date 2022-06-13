// a. Nama user gender M
db.users.find( { gender: 'M' });

// b. Tampilkan product id=3
db.products.find( { id: 3 });

// c. Hitung jumlah user gender Perempuan atau F (Female)
db.users.find( { gender: 'F' }).count();

// d. Tampilkan data pel dg urutan sesuai abjad
db.users.find().sort({ "name": 1 });

// e. Tampilkan 5 data pada product
db.products.find().limit(5);