CREATE TABLE product_variants
(
    id         bigserial primary key,
    product_id bigint references products (id) ON DELETE CASCADE,
    name       varchar(255),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
)