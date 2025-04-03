-- Insert 10 menu items
INSERT INTO menu_items (name, description, price, categories, allergens, metadata) VALUES
    ('Classic Espresso', 'Single shot of Arabica coffee', 3.50, '{coffee}', '{none}', '{"strength": "strong"}'),
    ('Cappuccino', 'Espresso with steamed whole milk and foam', 4.75, '{coffee,milk}', '{dairy}', '{"foam_thickness": "1cm"}'),
    ('Vanilla Latte', 'Espresso with oat milk and vanilla syrup', 5.25, '{coffee,milk}', '{none}', '{"vegan": true}'),
    ('Iced Green Tea Latte', 'Matcha green tea with almond milk over ice', 5.50, '{tea,iced}', '{none}', '{"serving": "cold"}'),
    ('Decaf Americano', 'Smooth decaffeinated coffee', 3.25, '{coffee}', '{none}', '{"caffeine_level": "0mg"}'),
    ('Hazelnut Cappuccino', 'Rich espresso with hazelnut syrup', 5.00, '{coffee,milk}', '{dairy}', '{"flavor_intensity": "medium"}'),
    ('Butter Croissant', 'Freshly baked French-style croissant', 3.99, '{bakery}', '{gluten,dairy}', '{"heating": "optional"}'),
    ('Double Chocolate Cookie', 'Chocolate chip cookie with extra chunks', 2.99, '{bakery}', '{gluten,dairy}', '{"size": "large"}'),
    ('Caramel Macchiato', 'Layered espresso with vanilla and caramel', 5.75, '{coffee,milk}', '{dairy}', '{"layers": 3}'),
    ('Earl Grey Tea', 'Traditional bergamot-flavored black tea', 2.50, '{tea}', '{none}', '{"serving": "hot"}');

-- Link menu items to inventory ingredients
INSERT INTO menu_item_ingredients (menu_item_id, ingredient_id, quantity_used) VALUES
    -- Classic Espresso (ID 1)
    (1, 1, 0.02),  -- Arabica Coffee Beans
    
    -- Cappuccino (ID 2)
    (2, 1, 0.02),   -- Arabica Coffee Beans
    (2, 6, 0.20),   -- Whole Milk
    (2, 9, 0.05),   -- Heavy Cream
    
    -- Vanilla Latte (ID 3)
    (3, 1, 0.02),   -- Arabica Coffee Beans
    (3, 8, 0.25),   -- Oat Milk
    (3, 11, 0.03),  -- Vanilla Syrup
    
    -- Iced Green Tea Latte (ID 4)
    (4, 4, 0.015),  -- Green Tea Leaves
    (4, 7, 0.30),   -- Almond Milk
    
    -- Decaf Americano (ID 5)
    (5, 5, 0.025),  -- Decaf Coffee Beans
    
    -- Hazelnut Cappuccino (ID 6)
    (6, 1, 0.02),   -- Arabica Coffee Beans
    (6, 6, 0.20),   -- Whole Milk
    (6, 13, 0.04),  -- Hazelnut Syrup
    
    -- Butter Croissant (ID 7)
    (7, 14, 1.00),  -- Croissant Dough
    
    -- Double Chocolate Cookie (ID 8)
    (8, 15, 0.10),  -- Chocolate Chips
    (8, 10, 0.05),  -- White Sugar
    
    -- Caramel Macchiato (ID 9)
    (9, 1, 0.02),   -- Arabica Coffee Beans
    (9, 6, 0.20),   -- Whole Milk
    (9, 12, 0.05),  -- Caramel Syrup
    
    -- Earl Grey Tea (ID 10)
    (10, 3, 0.01);  -- Earl Grey Tea Leaves

-- Price history spanning January to June 2023
INSERT INTO price_history (menu_item_id, old_price, new_price, changed_at) VALUES
    -- Classic Espresso (ID 1)
    (1, 3.00, 3.25, '2023-01-15 00:00:00'),
    (1, 3.25, 3.50, '2023-04-01 00:00:00'),
    
    -- Cappuccino (ID 2)
    (2, 4.50, 4.75, '2023-02-10 00:00:00'),
    (2, 4.75, 5.00, '2023-05-15 00:00:00'),
    (2, 5.00, 4.75, '2023-06-01 00:00:00'),  -- Price drop
    
    -- Vanilla Latte (ID 3)
    (3, 5.00, 5.25, '2023-03-01 00:00:00'),
    (3, 5.25, 5.50, '2023-05-20 00:00:00'),
    
    -- Iced Green Tea Latte (ID 4)
    (4, 5.25, 5.50, '2023-02-28 00:00:00'),
    (4, 5.50, 5.75, '2023-06-10 00:00:00'),
    
    -- Decaf Americano (ID 5)
    (5, 3.00, 3.25, '2023-01-20 00:00:00'),
    (5, 3.25, 3.50, '2023-04-15 00:00:00'),
    
    -- Hazelnut Cappuccino (ID 6)
    (6, 4.75, 5.00, '2023-03-15 00:00:00'),
    (6, 5.00, 5.25, '2023-06-01 00:00:00'),
    
    -- Butter Croissant (ID 7)
    (7, 3.75, 3.99, '2023-02-01 00:00:00'),
    (7, 3.99, 4.25, '2023-05-05 00:00:00'),
    
    -- Double Chocolate Cookie (ID 8)
    (8, 2.50, 2.75, '2023-01-25 00:00:00'),
    (8, 2.75, 2.99, '2023-04-10 00:00:00'),
    
    -- Caramel Macchiato (ID 9)
    (9, 5.50, 5.75, '2023-03-10 00:00:00'),
    (9, 5.75, 6.00, '2023-06-15 00:00:00'),
    
    -- Earl Grey Tea (ID 10)
    (10, 2.25, 2.50, '2023-02-15 00:00:00'),
    (10, 2.50, 2.75, '2023-05-25 00:00:00');