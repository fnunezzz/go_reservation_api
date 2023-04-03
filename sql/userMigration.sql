create table if not exists users (
    id varchar primary key,
    first_name varchar(50) not null,
    last_name varchar(50) not null,
    age smallint not null,
    cpf varchar(11) not null,
    email varchar(100) not null,
    password varchar(5000) not null,
    created_at timestamp default now()
);