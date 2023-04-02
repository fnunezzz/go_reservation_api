create table if not exists users (
    id varchar primary key,
    first_name varchar not null,
    last_name varchar not null,
    age smallint not null,
    email varchar not null,
    password varchar(5000) not null
)