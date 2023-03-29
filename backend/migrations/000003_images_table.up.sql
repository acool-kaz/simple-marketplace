CREATE TABLE IF NOT EXISTS images(
    id serial primary key,
    product_id integer not null,
    url text not null,
    foreign key (product_id) references products (id) on delete cascade
);