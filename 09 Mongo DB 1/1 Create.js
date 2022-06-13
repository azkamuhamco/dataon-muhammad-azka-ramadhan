db.operators.insert([
    {id: 1, name: 'Telkomsel'}, {id: 2, name: 'Indosat Ooredoo'}, {id: 3, name: 'XL Axiata'}, 
    {id: 4, name: 'Smartfren'}, {id: 5, name: 'Three'}
]);

db.product_types.insert([
    {id: 1, name: 'Pulsa'}, {id: 2, name: 'Paket Data'}, {id: 3, name: 'Paket Komunikasi'}
]);

db.products.insert([
    {id: 1, product_type_id: 1, operator_id: 3, code: 'XL-PULSA-50',     price:  50000, name: 'XL Pulsa 50000',                  status: 1},
    {id: 2, product_type_id: 1, operator_id: 3, code: 'XL-PULSA-100',    price: 100000, name: 'XL Pulsa 100000',                 status: 1},
    {id: 3, product_type_id: 2, operator_id: 1, code: 'TSEL-DATA-5GB',   price:  45000, name: 'Telkomsel Data Reguler 5 GB',     status: 1}, 
    {id: 4, product_type_id: 2, operator_id: 1, code: 'TSEL-DATA-10GB',  price:  75000, name: 'Telkomsel Data Reguler 10 GB',    status: 1},
    {id: 5, product_type_id: 2, operator_id: 1, code: 'TSEL-DATA-15GB',  price: 120000, name: 'Telkomsel Data Reguler 15 GB',    status: 1},
    {id: 6, product_type_id: 3, operator_id: 4, code: 'SMART-KOM-TEL20', price:  20000, name: 'Smartfren Bebas Bicara 20 Menit', status: 1},
    {id: 7, product_type_id: 3, operator_id: 4, code: 'SMART-KOM-TEL50', price:  45000, name: 'Smartfren Bebas Bicara 50 Menit', status: 1},
    {id: 8, product_type_id: 3, operator_id: 4, code: 'SMART-KOM-TEL90', price:  20000, name: 'Smartfren Bebas Bicara 90 Menit', status: 1}
]);

db.product_descriptions.insert([
    {id: 1, description: 'XL Pulsa 50000 masa aktif 60 hari'                 },
    {id: 2, description: 'XL Pulsa 100000 masa aktif 90 hari'                },
    {id: 3, description: 'Telkomsel Data Reguler 5 GB masa aktif 10 hari'    }, 
    {id: 4, description: 'Telkomsel Data Reguler 10 GB masa aktif 20 hari'   },
    {id: 5, description: 'Telkomsel Data Reguler 15 GB masa aktif 30 hari'   },
    {id: 6, description: 'Smartfren Bebas Bicara 20 Menit masa aktif 1 hari' },
    {id: 7, description: 'Smartfren Bebas Bicara 50 Menit masa aktif 3 hari' },
    {id: 8, description: 'Smartfren Bebas Bicara 90 Menit masa aktif 6 hari' }
]);

db.payment_methods.insert([
    {id: 1, name: 'Tunai', status: 1}, {id: 2, name: 'Debit', status: 1}, {id: 3, name: 'E-pay', status: 1}
]);

db.users.insert([
    {id: 1, name: 'Karim', status: 1, dob: '1994-07-05', gender: 'M'},
    {id: 2, name: 'Karina', status: 1, dob: '1992-07-15', gender: 'F'},
    {id: 3, name: 'Yoda', status: 1, dob: '2020-04-04', gender: 'M'},
    {id: 4, name: 'Rahma', status: 1, dob: '1989-10-09', gender: 'F'},
    {id: 5, name: 'Azka', status: 1, dob: '1991-01-07', gender: 'M'}
]);

db.transactions.insert([
    {id: 1, user_id: 2, payment_methods_id: 2, status: 1, total_qty: 3, totalprice: 145000},
    {id: 2, user_id: 4, payment_methods_id: 3, status: 1, total_qty: 3, totalprice: 185000},
    {id: 1, user_id: 1, payment_methods_id: 1, status: 1, total_qty: 3, totalprice: 145000}
]);

db.transaction_details.insert([
    {transaction_id: 1, product_id: 1, status: 'Terbayar', qty: 1, price:  50000},
    {transaction_id: 1, product_id: 4, status: 'Terbayar', qty: 1, price:  75000},
    {transaction_id: 1, product_id: 6, status: 'Terbayar', qty: 1, price:  20000},

    {transaction_id: 2, product_id: 2, status: 'Terbayar', qty: 1, price: 100000},
    {transaction_id: 2, product_id: 3, status: 'Terbayar', qty: 1, price:  40000},
    {transaction_id: 2, product_id: 7, status: 'Terbayar', qty: 1, price:  45000},

    {transaction_id: 3, product_id: 2, status: 'Terbayar', qty: 1, price: 100000},
    {transaction_id: 3, product_id: 5, status: 'Terbayar', qty: 1, price: 120000},
    {transaction_id: 3, product_id: 8, status: 'Terbayar', qty: 1, price:  80000}
]);

