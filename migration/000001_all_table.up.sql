

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(500) NOT NULL,
    user_type VARCHAR(500) NOT NULL,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    weight int UNIQUE NOT NULL,
    unit VARCHAR(100) UNIQUE NOT NULL,
    quantity int UNIQUE NOT NULL DEFAULT 1,
    price_per_product float NOT NULL,
    version int UNIQUE NOT NULL DEFAULT 1,
    is_deleted bool DEFAULT false,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS sub_order (
    id SERIAL PRIMARY KEY,
    unique_order_id VARCHAR(100) NOT NULL,
    product_id int NOT NULL,
    quantity int NOT NULL DEFAULT 1,
    status VARCHAR(50) NOT NULL,  -- created | shipped | out for delivery | delivered | cancelled
    version int UNIQUE NOT NULL DEFAULT 1,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS user_order (
    id VARCHAR(100) PRIMARY KEY,
    user_id int NOT NULL,
    unique_order_id VARCHAR(100) NOT NULL,
    status VARCHAR(50) NOT NULL default 'CREATED', -- created | completed | cancelled
    total_price float UNIQUE NOT NULL,
    version int UNIQUE NOT NULL DEFAULT 1,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);