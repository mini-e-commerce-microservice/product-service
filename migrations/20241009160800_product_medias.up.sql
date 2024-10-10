CREATE TYPE media_type AS ENUM('image','video');

CREATE TABLE product_medias
(
    id               bigserial primary key,
    product_id       bigint references products (id) ON DELETE CASCADE,
    media            varchar(255) NOT NULL,
    media_type       media_type   not null,
    is_primary_media boolean
)