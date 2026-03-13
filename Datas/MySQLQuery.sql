CREATE DATABASE IF NOT EXISTS server02_golang_db;

USE server02_golang_db;

CREATE TABLE
    orders (
        id INT AUTO_INCREMENT PRIMARY KEY,
        user_id VARCHAR(24) NOT NULL,
        product_name VARCHAR(255) NOT NULL,
        quantity INT NOT NULL,
        price DECIMAL(10, 2) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );