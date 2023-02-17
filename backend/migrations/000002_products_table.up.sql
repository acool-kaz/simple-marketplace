CREATE TABLE IF NOT EXISTS products(
    id serial primary key,
    user_id integer not null,
    name varchar(255) not null,
    description text not null,
    price numeric not null,
    foreign key (user_id) references users (id) on delete cascade
);