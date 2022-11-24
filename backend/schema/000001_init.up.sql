CREATE TABLE users
(
    id serial not null unique, 
    name varchar(255) not null, 
    username varchar(255) not null unique,
    city varchar(255) not null, 
    street varchar(255) not null, 
    card_nums varchar(255) not null,
    card_m_y integer[2] not null,
    password varchar(255) not null
);

CREATE TABLE admins
(
    id serial not null unique,
    username varchar(255) not null,
    password varchar(255) not null
);

INSERT INTO admins (username, password) VALUES
('acool', 'acool'),
('almas', 'almas'),
('ramazan', 'ramazan');

CREATE TABLE products
(
    id serial not null unique, 
    user_id integer not null,
    name varchar(255) not null, 
    description varchar(255) not null,  
    tags varchar(255)[] not null, 
    price integer not null,
    created_at varchar(255) not null
);

CREATE TABLE top_secret
(
    info varchar(255)
);

INSERT INTO top_secret (info) VALUES
('поздравляю ты нашел секретную информацию'),
('https://jut.su/oneepiece/episode-362.html'),
('https://www.youtube.com/watch?v=xm3YgoEiEDc'),
('http://127.0.0.1:5500/frontend/admin.html');