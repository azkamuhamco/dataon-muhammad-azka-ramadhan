CREATE DATABASE IF NOT EXISTS alta_online_shop;

USE alta_online_shop;

CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status SMALLINT NOT NULL,
    dob DATE NOT NULL,
    gender CHAR(1) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS product_types (
    id INT PRIMARY KEY, 
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL    
);

CREATE TABLE IF NOT EXISTS operators (
    id INT PRIMARY KEY, 
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL    
);

CREATE TABLE IF NOT EXISTS product_descriptions (
    id INT PRIMARY KEY,
    description TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL    
);

CREATE TABLE IF NOT EXISTS payment_methods (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL      
);

CREATE TABLE IF NOT EXISTS products (
    id INT PRIMARY KEY,
    product_type_id INT,
    operator_id INT,
    code VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (product_type_id) REFERENCES product_types(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (operator_id) REFERENCES operators(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS transactions (
    id INT PRIMARY KEY,
    user_id INT,
    payment_method_id INT,
    status VARCHAR(10) NOT NULL,
    total_qty INT NOT NULL,
    total_price NUMERIC(25,2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS transaction_details (
    transaction_id INT,
    product_id INT,
    status VARCHAR(10) NOT NULL,
    total_qty INT NOT NULL,
    total_price NUMERIC(25,2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (transaction_id, product_id),
    FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS kurir (
    id INT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

ALTER TABLE kurir ADD ongkos_dasar NUMERIC(25,2) NOT NULL;

ALTER TABLE kurir RENAME TO shipping;

DROP TABLE shipping;
