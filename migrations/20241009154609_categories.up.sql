CREATE TABLE categories
(
    id         bigserial primary key,
    name       varchar(255) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL,
    updated_at TIMESTAMPTZ  NOT NULL,
    deleted_at TIMESTAMPTZ
)