CREATE DATABASE IF NOT EXISTS items;
USE items;

CREATE TABLE IF NOT EXISTS items (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    price DOUBLE NOT NULL,
    category VARCHAR(255) NOT NULL,
    in_stock INT UNSIGNED NOT NULL,
    currency VARCHAR(10) DEFAULT 'KZT',
    status VARCHAR(50) DEFAULT 'in stock',
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);
