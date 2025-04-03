INSERT INTO inventory (item_name, quantity, unit, price, created_at, updated_at) VALUES
    -- Coffee & Tea
    ('Arabica Coffee Beans', 25.0, 'kg', 15.99, NOW(), NOW()),
    ('Robusta Coffee Beans', 18.5, 'kg', 12.50, NOW(), NOW()),
    ('Earl Grey Tea Leaves', 8.2, 'kg', 22.75, NOW(), NOW()),
    ('Green Tea Leaves', 6.7, 'kg', 18.99, NOW(), NOW()),
    ('Decaf Coffee Beans', 12.3, 'kg', 17.25, NOW(), NOW()),
    
    -- Dairy & Alternatives
    ('Whole Milk', 45.0, 'liters', 3.20, NOW(), NOW()),
    ('Almond Milk', 22.5, 'liters', 4.50, NOW(), NOW()),
    ('Oat Milk', 30.0, 'liters', 4.25, NOW(), NOW()),
    ('Heavy Cream', 15.0, 'liters', 5.99, NOW(), NOW()),
    
    -- Sweeteners & Flavorings
    ('White Sugar', 32.0, 'kg', 2.10, NOW(), NOW()),
    ('Vanilla Syrup', 18.0, 'liters', 8.50, NOW(), NOW()),
    ('Caramel Syrup', 12.5, 'liters', 9.25, NOW(), NOW()),
    ('Hazelnut Syrup', 10.0, 'liters', 9.75, NOW(), NOW()),
    
    -- Bakery Ingredients
    ('Croissant Dough', 120, 'units', 1.20, NOW(), NOW()),
    ('Chocolate Chips', 14.0, 'kg', 6.99, NOW(), NOW()),
    ('Coconut Milk', 25.0, 'liters', 4.75, NOW(), NOW()),
    ('Pumpkin Spice Syrup', 8.0, 'liters', 10.99, NOW(), NOW()),
    ('Bagel Dough', 80, 'units', 1.50, NOW(), NOW()),
    ('Dark Chocolate Chunks', 12.0, 'kg', 8.25, NOW(), NOW()),
    ('Lavender Tea Leaves', 5.5, 'kg', 20.50, NOW(), NOW());


-- Inventory transactions (usage and restocking)
INSERT INTO inventory_transactions (inventory_id, quantity_change, reason, created_at) VALUES
    -- Arabica Coffee Beans (ID 1)
    (1, -0.85, 'Daily coffee orders', '2023-01-05 08:00'),
    (1, 10.0, 'Monthly restock', '2023-01-28 14:00'),
    (1, -1.20, 'Weekend rush orders', '2023-02-12 09:30'),
    (1, 15.0, 'Bulk purchase', '2023-03-15 11:00'),
    (1, -0.95, 'Special event orders', '2023-04-20 16:45'),
    (1, 12.0, 'Supplier delivery', '2023-05-10 10:15'),
    (1, -1.10, 'Holiday season demand', '2023-06-08 12:30'),

    -- Whole Milk (ID 6)
    (6, -8.5, 'Daily beverage prep', '2023-01-07 07:30'),
    (6, 20.0, 'Weekly dairy delivery', '2023-01-14 09:00'),
    (6, -7.2, 'Valentine''s Day specials', '2023-02-14 10:15'),
    (6, 25.0, 'Milk tanker refill', '2023-03-01 08:45'),
    (6, -6.8, 'Spring break demand', '2023-04-05 11:20'),
    (6, 30.0, 'Summer stock preparation', '2023-05-20 10:00'),
    (6, -9.1, 'Weekend iced drinks', '2023-06-17 13:15'),

    -- Oat Milk (ID 8)
    (8, -5.5, 'Vegan week promotion', '2023-01-20 09:30'),
    (8, 15.0, 'Plant-based milk order', '2023-02-10 11:00'),
    (8, -4.2, 'Lactose-free demand', '2023-03-12 10:45'),
    (8, 20.0, 'Bulk oat milk delivery', '2023-04-18 14:30'),
    (8, -6.8, 'Summer smoothies', '2023-05-25 12:15'),
    (8, 25.0, 'Pre-holiday stock', '2023-06-20 09:00'),

    -- Vanilla Syrup (ID 11)
    (11, -0.8, 'Daily latte flavoring', '2023-01-10 08:15'),
    (11, 5.0, 'Syrup restock', '2023-02-05 10:30'),
    (11, -1.2, 'Special dessert menu', '2023-03-18 14:20'),
    (11, 8.0, 'Seasonal flavor prep', '2023-04-22 11:45'),
    (11, -0.9, 'Iced coffee season', '2023-05-15 13:00'),
    (11, 6.5, 'Monthly syrup order', '2023-06-12 09:30'),

    -- Croissant Dough (ID 14)
    (14, -18, 'Morning rush pastries', '2023-01-08 06:45'),
    (14, 50, 'Weekly bakery delivery', '2023-01-15 07:00'),
    (14, -25, 'Weekend brunch service', '2023-02-19 08:30'),
    (14, 60, 'Valentine''s prep', '2023-02-10 06:15'),
    (14, -32, 'Holiday special orders', '2023-04-07 07:45'),
    (14, 70, 'Summer stock increase', '2023-05-05 06:30'),
    (14, -28, 'Father''s Day demand', '2023-06-18 08:00'),

    -- Chocolate Chips (ID 15)
    (15, -2.5, 'Cookie baking batch', '2023-01-12 10:00'),
    (15, 10.0, 'Bulk chocolate order', '2023-02-08 11:15'),
    (15, -1.8, 'Dessert specials', '2023-03-14 14:30'),
    (15, 8.0, 'Seasonal restock', '2023-04-19 09:45'),
    (15, -3.2, 'Summer ice cream toppings', '2023-05-22 12:00'),
    (15, 12.0, 'Pre-summer stock', '2023-06-14 10:30'),

    -- Green Tea Leaves (ID 4)
    (4, -0.35, 'Daily matcha drinks', '2023-01-09 08:45'),
    (4, 3.0, 'Monthly tea restock', '2023-02-01 10:00'),
    (4, -0.45, 'Spring tea promotion', '2023-03-20 11:15'),
    (4, 5.0, 'Bulk tea purchase', '2023-04-25 14:20'),
    (4, -0.6, 'Iced tea season start', '2023-05-12 09:30'),
    (4, 4.5, 'Summer stock', '2023-06-10 08:00'),

    -- Heavy Cream (ID 9)
    (9, -2.1, 'Dessert preparations', '2023-01-11 09:15'),
    (9, 8.0, 'Dairy delivery', '2023-02-03 10:45'),
    (9, -3.0, 'Specialty drinks', '2023-03-17 12:30'),
    (9, 10.0, 'Holiday season stock', '2023-04-21 11:00'),
    (9, -2.8, 'Summer parfaits', '2023-05-18 13:15'),
    (9, 12.0, 'Monthly restock', '2023-06-22 09:45');

-- Add 30 more transactions with smaller adjustments
INSERT INTO inventory_transactions (inventory_id, quantity_change, reason, created_at) VALUES
    (2, -0.5, 'Espresso blend maintenance', '2023-01-16 08:00'),
    (3, -0.2, 'Tea service spillage', '2023-02-09 10:30'),
    (5, -0.3, 'Decaf afternoon orders', '2023-03-22 14:15'),
    (7, -1.5, 'Almond milk latte surge', '2023-04-06 11:00'),
    (10, -0.4, 'Sugar bowl refills', '2023-05-19 09:45'),
    (12, -0.6, 'Car drizzle maintenance', '2023-06-11 12:30'),
    (13, -0.3, 'Hazelnut flavor week', '2023-01-25 08:45'),
    (6, -2.0, 'Milk frother calibration loss', '2023-02-14 10:00'),
    (8, -1.2, 'Oat milk expiration', '2023-03-28 14:30'),
    (11, -0.2, 'Vanilla pump adjustment', '2023-04-15 11:15'),
    (14, -3, 'Croissant shaping loss', '2023-05-08 07:00'),
    (15, -0.5, 'Chocolate tasting samples', '2023-06-01 09:30'),
    (4, -0.1, 'Tea ceremony event', '2023-01-31 10:45'),
    (9, -0.8, 'Whipped cream practice', '2023-02-18 12:00'),
    (5, -0.4, 'Evening decaf orders', '2023-03-29 15:30'),
    (7, -0.7, 'Almond milk trial', '2023-04-12 10:15'),
    (10, -0.3, 'Sugar jar refill', '2023-05-23 08:45'),
    (12, -0.4, 'Caramel pump maintenance', '2023-06-16 11:00'),
    (13, -0.2, 'Hazelnut test batch', '2023-01-19 09:30'),
    (6, -1.5, 'Latte art class', '2023-02-22 14:00');