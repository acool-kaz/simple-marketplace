CREATE TABLE IF NOT EXISTS users(
    id serial primary key,
    first_name varchar(50) not null, 
    second_name varchar(50) not null, 
    email varchar(255) not null unique, 
    username varchar(50) not null unique, 
    password varchar(50) not null
);