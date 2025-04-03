-- Insert 30 orders with different statuses
INSERT INTO orders (customer_name, status, total_amount, payment_method, created_at) VALUES
    ('John Smith', 'completed', 8.74, 'card', '2023-01-20 08:30'),
    ('Emma Johnson', 'processing', 12.50, 'online', '2023-02-15 09:45'),
    ('Michael Brown', 'cancelled', 5.25, 'cash', '2023-03-10 11:15'),
    ('Sarah Wilson', 'completed', 10.99, 'card', '2023-04-05 14:20'),
    ('David Miller', 'pending', 7.50, 'cash', '2023-05-18 16:00'),
    ('Olivia Davis', 'completed', 15.75, 'online', '2023-06-12 10:10'),
    ('James Wilson', 'processing', 9.25, 'card', '2023-01-25 12:30'),
    ('Sophia Moore', 'completed', 6.99, 'cash', '2023-02-28 13:45'),
    ('William Taylor', 'cancelled', 18.50, 'online', '2023-03-15 15:00'),
    ('Emily Anderson', 'pending', 4.75, 'card', '2023-04-20 17:20'),
    ('Daniel Thomas', 'completed', 11.25, 'cash', '2023-05-05 08:45'),
    ('Ava Jackson', 'processing', 14.00, 'online', '2023-06-20 09:30'),
    ('Lucas White', 'completed', 8.50, 'card', '2023-01-30 10:15'),
    ('Mia Harris', 'cancelled', 7.25, 'cash', '2023-02-10 11:40'),
    ('Benjamin Martin', 'pending', 16.75, 'online', '2023-03-25 14:50'),
    ('Charlotte Thompson', 'completed', 5.99, 'card', '2023-04-15 16:10'),
    ('Henry Garcia', 'processing', 13.25, 'cash', '2023-05-10 17:30'),
    ('Amelia Martinez', 'completed', 9.75, 'online', '2023-06-05 08:20'),
    ('Ethan Robinson', 'cancelled', 10.50, 'card', '2023-01-15 09:45'),
    ('Isabella Clark', 'pending', 7.00, 'cash', '2023-02-20 12:00'),
    ('Alexander Rodriguez', 'completed', 14.99, 'online', '2023-03-05 13:15'),
    ('Sophia Lewis', 'processing', 6.25, 'card', '2023-04-10 14:30'),
    ('Mason Lee', 'completed', 11.75, 'cash', '2023-05-25 15:45'),
    ('Charlotte Walker', 'cancelled', 8.25, 'online', '2023-06-15 16:50'),
    ('Elijah Hall', 'pending', 12.50, 'card', '2023-01-10 17:05'),
    ('Harper Young', 'completed', 7.99, 'cash', '2023-02-05 08:20'),
    ('Daniel Hernandez', 'processing', 15.25, 'online', '2023-03-20 09:35'),
    ('Evelyn King', 'completed', 10.75, 'card', '2023-04-25 10:50'),
    ('Logan Wright', 'cancelled', 5.50, 'cash', '2023-05-15 12:05'),
    ('Abigail Lopez', 'pending', 13.99, 'online', '2023-06-25 13:20');

-- Insert order items
INSERT INTO order_items (order_id, menu_item_id, quantity, price_at_order) VALUES
    -- Order 1 (Completed - Jan 20)
    (1, 1, 2, 3.25), (1, 7, 1, 3.75),
    
    -- Order 2 (Processing - Feb 15)
    (2, 2, 2, 4.75), (2, 9, 1, 5.75),
    
    -- Order 3 (Cancelled - Mar 10)
    (3, 3, 1, 5.25),
    
    -- Order 4 (Completed - Apr 5)
    (4, 4, 2, 5.50), (4, 8, 1, 2.99),
    
    -- Order 5 (Pending - May 18)
    (5, 5, 2, 3.25), (5, 10, 1, 2.50),
    
    -- Order 6 (Completed - Jun 12)
    (6, 6, 3, 5.00),
    
    -- Order 7 (Processing - Jan 25)
    (7, 2, 1, 4.50), (7, 7, 2, 3.75),
    
    -- Order 8 (Completed - Feb 28)
    (8, 9, 1, 5.50), (8, 10, 2, 2.25),
    
    -- Order 9 (Cancelled - Mar 15)
    (9, 1, 3, 3.25), (9, 3, 2, 5.00),
    
    -- Order 10 (Pending - Apr 20)
    (10, 4, 1, 5.50), (10, 8, 1, 2.99),
    
    -- Order 11 (Completed - May 5)
    (11, 5, 1, 3.25), (11, 6, 1, 4.75),
    
    -- Order 12 (Processing - Jun 20)
    (12, 2, 2, 4.75), (12, 9, 1, 6.00),
    
    -- Order 13 (Completed - Jan 30)
    (13, 7, 3, 3.75), (13, 10, 2, 2.25),
    
    -- Order 14 (Cancelled - Feb 10)
    (14, 3, 2, 5.00), (14, 8, 1, 2.75),
    
    -- Order 15 (Pending - Mar 25)
    (15, 1, 1, 3.25), (15, 2, 1, 4.75),
    
    -- Order 16 (Completed - Apr 15)
    (16, 4, 1, 5.50), (16, 7, 1, 3.99),
    
    -- Order 17 (Processing - May 10)
    (17, 5, 2, 3.25), (17, 6, 1, 5.00),
    
    -- Order 18 (Completed - Jun 5)
    (18, 9, 2, 5.75), (18, 10, 1, 2.50),
    
    -- Order 19 (Cancelled - Jan 15)
    (19, 2, 1, 4.50), (19, 8, 2, 2.50),
    
    -- Order 20 (Pending - Feb 20)
    (20, 3, 1, 5.00), (20, 7, 1, 3.75),
    
    -- Order 21 (Completed - Mar 5)
    (21, 1, 2, 3.25), (21, 9, 1, 5.50),
    
    -- Order 22 (Processing - Apr 10)
    (22, 4, 1, 5.50), (22, 6, 1, 5.00),
    
    -- Order 23 (Completed - May 25)
    (23, 5, 3, 3.25), (23, 10, 2, 2.50),
    
    -- Order 24 (Cancelled - Jun 15)
    (24, 2, 2, 4.75), (24, 7, 1, 3.99),
    
    -- Order 25 (Pending - Jan 10)
    (25, 3, 1, 5.00), (25, 8, 1, 2.75),
    
    -- Order 26 (Completed - Feb 5)
    (26, 9, 1, 5.50), (26, 4, 1, 5.25),
    
    -- Order 27 (Processing - Mar 20)
    (27, 1, 2, 3.25), (27, 6, 1, 4.75),
    
    -- Order 28 (Completed - Apr 25)
    (28, 7, 2, 3.99), (28, 10, 1, 2.50),
    
    -- Order 29 (Cancelled - May 15)
    (29, 2, 1, 4.75), (29, 9, 1, 5.75),
    
    -- Order 30 (Pending - Jun 25)
    (30, 5, 2, 3.50), (30, 8, 1, 2.99);

INSERT INTO order_status_history (order_id, status, changed_at) VALUES
    -- Order 1 (Completed)
    (1, 'pending', '2023-01-20 08:30:00'),
    (1, 'processing', '2023-01-20 08:42:00'),
    (1, 'completed', '2023-01-20 09:05:00'),

    -- Order 2 (Processing)
    (2, 'pending', '2023-02-15 09:45:00'),
    (2, 'processing', '2023-02-15 09:57:00'),

    -- Order 3 (Cancelled)
    (3, 'pending', '2023-03-10 11:15:00'),
    (3, 'processing', '2023-03-10 11:25:00'),
    (3, 'cancelled', '2023-03-10 11:40:00'),

    -- Order 4 (Completed)
    (4, 'pending', '2023-04-05 14:20:00'),
    (4, 'processing', '2023-04-05 14:35:00'),
    (4, 'completed', '2023-04-05 14:55:00'),

    -- Order 5 (Pending) - No changes
    (5, 'pending', '2023-05-18 16:00:00'),

    -- Order 6 (Completed)
    (6, 'pending', '2023-06-12 10:10:00'),
    (6, 'processing', '2023-06-12 10:22:00'),
    (6, 'completed', '2023-06-12 10:45:00'),

    -- Order 7 (Processing)
    (7, 'pending', '2023-01-25 12:30:00'),
    (7, 'processing', '2023-01-25 12:45:00'),

    -- Order 8 (Completed)
    (8, 'pending', '2023-02-28 13:45:00'),
    (8, 'processing', '2023-02-28 13:57:00'),
    (8, 'completed', '2023-02-28 14:15:00'),

    -- Order 9 (Cancelled)
    (9, 'pending', '2023-03-15 15:00:00'),
    (9, 'cancelled', '2023-03-15 15:10:00'),

    -- Order 10 (Pending) - No changes
    (10, 'pending', '2023-04-20 17:20:00'),

    -- Order 11 (Completed)
    (11, 'pending', '2023-05-05 08:45:00'),
    (11, 'processing', '2023-05-05 08:57:00'),
    (11, 'completed', '2023-05-05 09:15:00'),

    -- Order 12 (Processing)
    (12, 'pending', '2023-06-20 09:30:00'),
    (12, 'processing', '2023-06-20 09:45:00'),

    -- Order 13 (Completed)
    (13, 'pending', '2023-01-30 10:15:00'),
    (13, 'processing', '2023-01-30 10:27:00'),
    (13, 'completed', '2023-01-30 10:50:00'),

    -- Order 14 (Cancelled)
    (14, 'pending', '2023-02-10 11:40:00'),
    (14, 'processing', '2023-02-10 11:52:00'),
    (14, 'cancelled', '2023-02-10 12:05:00'),

    -- Order 15 (Pending) - No changes
    (15, 'pending', '2023-03-25 14:50:00'),

    -- Order 16 (Completed)
    (16, 'pending', '2023-04-15 16:10:00'),
    (16, 'processing', '2023-04-15 16:25:00'),
    (16, 'completed', '2023-04-15 16:45:00'),

    -- Order 17 (Processing)
    (17, 'pending', '2023-05-10 17:30:00'),
    (17, 'processing', '2023-05-10 17:45:00'),

    -- Order 18 (Completed)
    (18, 'pending', '2023-06-05 08:20:00'),
    (18, 'processing', '2023-06-05 08:35:00'),
    (18, 'completed', '2023-06-05 08:55:00'),

    -- Order 19 (Cancelled)
    (19, 'pending', '2023-01-15 09:45:00'),
    (19, 'processing', '2023-01-15 09:57:00'),
    (19, 'cancelled', '2023-01-15 10:10:00'),

    -- Order 20 (Pending) - No changes
    (20, 'pending', '2023-02-20 12:00:00'),

    -- Order 21 (Completed)
    (21, 'pending', '2023-03-05 13:15:00'),
    (21, 'processing', '2023-03-05 13:27:00'),
    (21, 'completed', '2023-03-05 13:50:00'),

    -- Order 22 (Processing)
    (22, 'pending', '2023-04-10 14:30:00'),
    (22, 'processing', '2023-04-10 14:45:00'),

    -- Order 23 (Completed)
    (23, 'pending', '2023-05-25 15:45:00'),
    (23, 'processing', '2023-05-25 15:57:00'),
    (23, 'completed', '2023-05-25 16:15:00'),

    -- Order 24 (Cancelled)
    (24, 'pending', '2023-06-15 16:50:00'),
    (24, 'cancelled', '2023-06-15 17:05:00'),

    -- Order 25 (Pending) - No changes
    (25, 'pending', '2023-01-10 17:05:00'),

    -- Order 26 (Completed)
    (26, 'pending', '2023-02-05 08:20:00'),
    (26, 'processing', '2023-02-05 08:35:00'),
    (26, 'completed', '2023-02-05 08:55:00'),

    -- Order 27 (Processing)
    (27, 'pending', '2023-03-20 09:35:00'),
    (27, 'processing', '2023-03-20 09:50:00'),

    -- Order 28 (Completed)
    (28, 'pending', '2023-04-25 10:50:00'),
    (28, 'processing', '2023-04-25 11:05:00'),
    (28, 'completed', '2023-04-25 11:25:00'),

    -- Order 29 (Cancelled)
    (29, 'pending', '2023-05-15 12:05:00'),
    (29, 'processing', '2023-05-15 12:17:00'),
    (29, 'cancelled', '2023-05-15 12:30:00'),

    -- Order 30 (Pending) - No changes
    (30, 'pending', '2023-06-25 13:20:00');