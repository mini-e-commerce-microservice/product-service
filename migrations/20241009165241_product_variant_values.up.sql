CREATE TABLE product_variant_values
(
    id                 bigserial primary key,
    product_variant_id bigint references product_variants (id) ON DELETE CASCADE,
    value              varchar(255)
)