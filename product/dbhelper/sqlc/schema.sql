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
