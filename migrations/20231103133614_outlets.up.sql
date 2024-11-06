CREATE TABLE outlets
(
    id          bigserial primary key,
    user_id     bigint       not null,
    logo        varchar(100) not null default '',
    name        varchar(100) not null,
    slogan      varchar(255) not null default '',
    description text not null default ''
)