CREATE TABLE users
(
    id serial not null unique,
    username varchar(256) not null unique,
    password_hash varchar(256) not null,
    coins int default 1000
);

CREATE TABLE coinHistory
(
    id serial primary key,
    from_user int not null,
    to_user int not null,
    amount int not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE inventory
(
    id serial primary key,
    user_id int not null,
    item varchar(256) not null,
    quantity int not null,
    UNIQUE (user_id, item)
);

CREATE TABLE items
(
    id serial primary key,
    item varchar(256),
    price int
);

INSERT INTO items (item, price) VALUES
    ('t-shirt', 80),
    ('cup', 20),
    ('book', 50),
    ('pen', 10),
    ('powerbank', 200),
    ('hoody', 300),
    ('umbrella', 200),
    ('socks', 10),
    ('wallet', 50),
    ('pink-hoody', 500);
