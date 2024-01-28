CREATE TABLE IF NOT EXISTS product (
    id bigserial PRIMARY KEY,
    name varchar(30),
    price varchar(30),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);