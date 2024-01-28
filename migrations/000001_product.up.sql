CREATE TABLE IF NOT EXISTS products (
    id bigserial PRIMARY KEY,
    name varchar(30),
    price INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);  