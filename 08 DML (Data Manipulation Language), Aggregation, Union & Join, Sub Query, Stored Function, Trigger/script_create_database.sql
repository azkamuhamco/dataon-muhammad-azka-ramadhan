CREATE DATABASE alta_online_shop;

CREATE TABLE users (
    id INT(11) PRIMARY KEY,
    dob DATE NOT NULL,
    gender CHAR(1) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE product_types (
    id INT(11) PRIMARY KEY, 
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL    
);

CREATE TABLE operators (
    id INT(11), 
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL    
);

CREATE TABLE product_descriptions (
    id INT(11) PRIMARY KEY,
    description TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL    
);

CREATE TABLE payment_methods (
    id INT(11) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL      
);

CREATE TABLE products (
    id INT(11) PRIMARY KEY,
    product_type_id INT(11),
    operator_id INT(11),
    code VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (product_type_id) REFERENCES product_types(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (operator_id) REFERENCES operators(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE transactions (
    id INT(11) PRIMARY KEY,
    user_id INT(11),
    payment_method_id INT(11),
    status VARCHAR(10) NOT NULL,
    total_qty INT(11) NOT NULL,
    total_price NUMERIC(25,2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE transaction_details (
    transaction_id INT(11) PRIMARY KEY,
    product_id INT(11) PRIMARY KEY,
    status VARCHAR(10) NOT NULL,
    total_qty INT(11) NOT NULL,
    total_price NUMERIC(25,2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE kurir (
    id INT(11),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

ALTER TABLE kurir ADD ongkos_dasar NUMERIC(25,2) NOT NULL;

ALTER TABLE kurir RENAME TO shipping;

DROP TABLE shipping;
