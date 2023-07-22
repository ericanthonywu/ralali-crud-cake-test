create table if not exists cakes
(
    id          bigint unsigned auto_increment primary key unique,
    title       varchar(200)                       not null,
    description text                               not null,
    rating      float                              not null,
    image       varchar(200)                       not null,
    created_at  datetime default current_timestamp not null,
    updated_at  datetime                           null,
    deleted_at  datetime                           null
);

