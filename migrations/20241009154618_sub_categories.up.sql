CREATE TABLE sub_categories
(
    id          bigserial primary key,
    category_id bigint references categories (id) ON DELETE CASCADE,
    name        varchar(255) NOT NULL,
    created_at  TIMESTAMPTZ  NOT NULL,
    updated_at  TIMESTAMPTZ  NOT NULL,
    deleted_at  TIMESTAMPTZ
)