CREATE TABLE sub_category_items
(
    id              bigserial primary key,
    category_id     bigint references categories (id) ON DELETE CASCADE,
    sub_category_id bigint references sub_categories (id) ON DELETE CASCADE,
    name            varchar(255) NOT NULL UNIQUE,
    size_guide      boolean      not null,
    created_at      TIMESTAMPTZ  NOT NULL,
    updated_at      TIMESTAMPTZ  NOT NULL,
    deleted_at      TIMESTAMPTZ
)