CREATE TABLE product_variant_item_options
(
    product_variant_item_id   bigint references product_variant_items (id) on delete cascade,
    product_variant_option_id bigint references product_variant_options (id) on delete cascade
)