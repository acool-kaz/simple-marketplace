CREATE TABLE IF NOT EXISTS products(
    id serial primary key,
    name varchar(255) not null unique
);