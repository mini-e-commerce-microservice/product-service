CREATE TYPE product_condition AS ENUM ('new', 'second-hand');

CREATE TABLE products
(
    id                bigserial primary key,
    name              varchar(255)      NOT NULL,
    description       text,
    product_condition product_condition NOT NULL,
    is_used_variant   boolean           NOT NULL,
    minimum_purchase  integer           not null,
    size_guide_image  varchar(255),
    created_at        TIMESTAMPTZ       NOT NULL,
    updated_at        TIMESTAMPTZ       NOT NULL,
    deleted_at        TIMESTAMPTZ
)