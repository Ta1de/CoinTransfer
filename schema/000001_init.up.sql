CREATE TABLE users
(
    id serial not null unique,
    username varchar(256) not null unique,
    password_hash varchar(256) not null,
    coins int default 0
)