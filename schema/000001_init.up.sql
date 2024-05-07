CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    surname VARCHAR(64),
    image TEXT,
    password VARCHAR(256) NOT NULL,
    is_premium BOOL DEFAULT false
);