CREATE TABLE product_variant_items
(
    id                   bigserial primary key,
    product_id           bigint references products (id) ON DELETE CASCADE,
    is_primary_product   boolean        not null,
    price                decimal(10, 2) not null,
    stock                integer        not null,
    sku                  varchar(255),
    weight               integer        not null,
    package_length       decimal(10, 2) not null,
    package_width        decimal(10, 2) not null,
    package_height       decimal(10, 2) not null,
    dimensional_weight   decimal(10, 2) not null,
    is_active            boolean        not null,

    created_at           TIMESTAMPTZ    NOT NULL,
    updated_at           TIMESTAMPTZ    NOT NULL,
    deleted_at           TIMESTAMPTZ
)