CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    user_name varchar(30),
    full_name varchar(50),
    password varchar(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);  