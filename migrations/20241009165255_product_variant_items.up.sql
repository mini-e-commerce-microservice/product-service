CREATE TABLE product_variant_items
(
    id                         bigserial primary key,
    product_id                 bigint references products (id) ON DELETE CASCADE,
    product_variant_value_1_id bigint references product_variant_values (id) on delete cascade,
    product_variant_value_2_id bigint references product_variant_values (id) on delete cascade,
    is_primary_product         boolean        not null,
    price                      decimal(10, 2) not null,
    stock                      integer        not null,
    sku                        varchar(255),
    weight                     integer        not null,
    package_length             integer not null,
    package_width              integer not null,
    package_height             integer not null,
    dimensional_weight         numeric(10, 2) not null,
    is_active                  boolean        not null,
    image                      varchar(255),
    created_at                 TIMESTAMPTZ    NOT NULL,
    updated_at                 TIMESTAMPTZ    NOT NULL,
    deleted_at                 TIMESTAMPTZ
)