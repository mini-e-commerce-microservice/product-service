INSERT INTO categories (name, created_at, updated_at)
VALUES ('Electronics', NOW(), NOW()),
       ('Clothing', NOW(), NOW());

INSERT INTO sub_categories (category_id, name, created_at, updated_at)
VALUES (1, 'Laptops', NOW(), NOW()),
       (1, 'Smartphones', NOW(), NOW()),
       (2, 'Mens Clothing', NOW(), NOW());

INSERT INTO sub_category_items (category_id, sub_category_id, name, size_guide, created_at, updated_at) VALUES
    (1, 1, 'Dell XPS 13', TRUE, NOW(), NOW()),
    (1, 2, 'iPhone 13', FALSE, NOW(), NOW()),
    (2, 3, 'Men''s T-shirt', TRUE, NOW(), NOW());
