create table if not exists tokens (
    id varchar(255) primary key,
    image varchar(1000) not null,
    description varchar (1000) not null,
    name varchar(255) unique not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp
);
