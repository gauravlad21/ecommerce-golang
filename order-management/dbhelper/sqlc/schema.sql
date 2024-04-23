
CREATE TABLE sub_order (
    id SERIAL PRIMARY KEY,
    unique_order_id VARCHAR(100) NOT NULL,
    product_id int NOT NULL,
    quantity int NOT NULL DEFAULT 1,
    status VARCHAR(50) NOT NULL,  -- created | shipped | out for delivery | delivered | cancelled
    version int UNIQUE NOT NULL DEFAULT 1,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE user_order (
    id VARCHAR(100) PRIMARY KEY,
    user_id int NOT NULL,
    unique_order_id VARCHAR(100) NOT NULL,
    status VARCHAR(50) NOT NULL default 'CREATED', -- created | completed | cancelled
    total_price float UNIQUE NOT NULL,
    version int UNIQUE NOT NULL DEFAULT 1,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);
