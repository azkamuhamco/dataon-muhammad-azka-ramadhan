-- UPDATE data product 1 dengan nama 'product dummy'
UPDATE products SET name = 'product dummy' WHERE id = 1

-- update qty = 3 pada transaction detail dg product id = 1
UPDATE transaction_details SET qty = 3 WHERE product_id = 1