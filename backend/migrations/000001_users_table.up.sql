CREATE TABLE IF NOT EXISTS users(
    id serial primary key,
    role varchar(50) not null default 'user',
    first_name varchar(50) not null, 
    second_name varchar(50) not null, 
    phone_number varchar(25) not null,
    email varchar(255) not null unique, 
    username varchar(50) not null unique, 
    password varchar(50) not null
);

INSERT INTO users
    (role, first_name, second_name, phone_number, email, username, password)
VALUES
    ('admin', 'Akylzhan', 'Eleusizov', '87081435719', 'pastapappie23@gmail.com', 'akylzhan', 'akilzhan2001');