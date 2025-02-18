CREATE TABLE users
(
    id serial not null unique,
    username varchar(256) not null unique,
    password_hash varchar(256) not null,
    coins int default 0
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
    item varchar(256),
    quantity int
);