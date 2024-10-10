CREATE TYPE product_condition AS ENUM ('new', 'second-hand');

CREATE TABLE products
(
    id                 bigserial primary key,
    name               varchar(255)      NOT NULL,
    description        text,
    product_condition  product_condition NOT NULL,
    is_used_variant    boolean           NOT NULL,
    minimum_purchase   integer           not null,
    size_guide_image   varchar(255),
    price              decimal(10, 2), -- Default price (could be used if no variants)
    stock              integer,        -- Default stock (could be used if no variants)
    sku                varchar(255),   -- Default sku (could be used if no variants)
    weight             integer,-- Default weight (could be used if no variants)
    package_length     decimal(10, 2), --Default package_length (could be used if no variants)
    package_width      decimal(10, 2), --Default package_width (could be used if no variants)
    package_height     decimal(10, 2), --Default package_height (could be used if no variants)
    dimensional_weight decimal(10, 2), -- Default dimensional_weight (could be used if no variants)
    is_active          boolean,-- Default is_active (could be used if no variants)
    trace_parent       varchar(255),
    created_at         TIMESTAMPTZ       NOT NULL,
    updated_at         TIMESTAMPTZ       NOT NULL,
    deleted_at         TIMESTAMPTZ
)