// 1. Menampilkan data buku authorid 1 dan 2
db.books.find( { authorID: { $in: [1,2] } } );

// 2. Tampilkan daftar buku dan harga author id 1
db.books.find( { authorID: 1 }, { title: 1, price: 1 } );

// 3. Tampilkan total jumlah halaman buku author id 2
db.books.aggregate([
    { $match: { authorID: 2 } },
    { $unwind: '$stats' },
    { $group: {
        _id: "$authorID",
        totalPages: { $sum: "$stats.page" }
    }}
]);

// 4. Tampilkan semua field books dan authors terkait
// Output 1
db.authors.aggregate([{
    $lookup: {
      from: "books",
      localField: "_id",
      foreignField: "authorID",
      as: "books"
    }
}]);

// Output 2
db.books.aggregate([{
    $lookup: {
      from: "authors",
      localField: "authorID",
      foreignField: "_id",
      as: "authors"
    }
}]);

// 5. Tampilkan semua field books, authors, publishers terkait

// 6. Tampilkan summary data authors, books, dan publishers sesuai Output

// 7. digital_outlet ingin memberikan diskon untuk setiap buku, diskon
//    ditentukan melihat harga buku tersebut dg pembagian seperti ini:
//    Price < 60000                 1 %
//    60000 < Price < 90000         2 %
//    90000 < Price                 3 %

// 8. Tampilkan semua nama buku, harga, nama author dan nama publisher, 
//    urutkan dari harga termahal ke harga termurah

// 9. Tampilkan data nama buku harga dan publisher, kemudian tampilkan
//    hanya data ke 3 dan ke 4
