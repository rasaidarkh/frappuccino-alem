CREATE TYPE ORDER_STATUS AS ENUM ('pending', 'processing', 'completed', 'cancelled');
CREATE TYPE PAYMENT_METHOD AS ENUM ('cash', 'card', 'online');
CREATE TYPE STAFF_ROLE AS ENUM ('barista', 'cashier', 'manager');

CREATE TABLE inventory (
    id SERIAL PRIMARY KEY,
    item_name TEXT NOT NULL,
    quantity DECIMAL(10,2) NOT NULL CHECK (quantity >= 0),
    unit TEXT NOT NULL,
    last_updated TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE menu_items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
    categories TEXT[] DEFAULT '{}',
    allergens TEXT[] DEFAULT '{}',
    metadata JSONB DEFAULT '{}'
);

CREATE TABLE menu_item_ingredients (
    menu_item_id INT REFERENCES menu_items(id)  ON DELETE CASCADE NOT NULL,
    ingredient_id INT REFERENCES inventory(id) ON DELETE CASCADE  NOT NULL,
    quantity_used DECIMAL(10,2) NOT NULL CHECK (quantity_used > 0),
    PRIMARY KEY (menu_item_id, ingredient_id)
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_name TEXT NOT NULL,
    status ORDER_STATUS NOT NULL DEFAULT 'pending',
    total_amount DECIMAL(10,2) NOT NULL CHECK (total_amount >= 0),
    payment_method PAYMENT_METHOD NOT NULL,
    special_instructions JSONB DEFAULT '{}',
    created_at TIMESTAMPTZ DEFAULT NOW(), 
    updated_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,
    menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity > 0),
    price_at_order DECIMAL(10,2) NOT NULL CHECK (price_at_order >= 0)
);

CREATE TABLE order_status_history (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE CASCADE  NOT NULL,
    status ORDER_STATUS NOT NULL,
    changed_at TIMESTAMPTZ DEFAULT NOW()
);


CREATE TABLE price_history (
    id SERIAL PRIMARY KEY,
    menu_item_id INT REFERENCES menu_items(id)  ON DELETE CASCADE NOT NULL,
    old_price DECIMAL(10,2) NOT NULL CHECK (old_price >= 0),
    new_price DECIMAL(10,2) NOT NULL CHECK (new_price >= 0),
    changed_at TIMESTAMPTZ DEFAULT NOW()
);


CREATE TABLE inventory_transactions (
    id SERIAL PRIMARY KEY,
    inventory_id INT REFERENCES inventory(id)ON DELETE CASCADE  NOT NULL,
    quantity_change DECIMAL(10,2) NOT NULL,
    reason TEXT NOT NULL,
    transaction_time TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE staff (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    role STAFF_ROLE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Insert mock staff
INSERT INTO staff (name, role) VALUES
    ('Alice Johnson', 'barista'),
    ('Bob Smith', 'cashier'),
    ('Charlie Davis', 'manager');

-- Insert mock inventory
INSERT INTO inventory (item_name, quantity, unit) VALUES
    ('Espresso Beans', 50, 'kg'),
    ('Milk', 100, 'liters'),
    ('Sugar', 30, 'kg'),
    ('Vanilla Syrup', 20, 'liters');

-- Insert mock menu items
INSERT INTO menu_items (name, description, price, categories, allergens, metadata) VALUES
    ('Espresso', 'Strong and bold coffee shot', 3.50, ARRAY['coffee'], ARRAY['caffeine'], '{}'),
    ('Cappuccino', 'Espresso with steamed milk and foam', 4.50, ARRAY['coffee', 'milk'], ARRAY['caffeine', 'dairy'], '{}'),
    ('Vanilla Latte', 'Espresso with vanilla syrup and milk', 5.00, ARRAY['coffee', 'milk'], ARRAY['caffeine', 'dairy'], '{"flavored": true}');

-- Insert mock orders
INSERT INTO orders (customer_name, status, total_amount, payment_method, special_instructions) VALUES
    ('John Doe', 'pending', 8.00, 'card', '{"extra_shot": true}'),
    ('Jane Smith', 'completed', 4.50, 'cash', '{}'),
    ('Michael Brown', 'processing', 5.00, 'online', '{"no_sugar": true}');

-- Insert mock order items
INSERT INTO order_items (order_id, menu_item_id, quantity, price_at_order) VALUES
    (1, 1, 1, 3.50),
    (1, 2, 1, 4.50),
    (2, 2, 1, 4.50),
    (3, 3, 1, 5.00);

-- Insert mock order status history
INSERT INTO order_status_history (order_id, status) VALUES
    (1, 'pending'),
    (2, 'completed'),
    (3, 'processing');

-- Insert mock price history
INSERT INTO price_history (menu_item_id, old_price, new_price) VALUES
    (1, 3.00, 3.50),
    (2, 4.00, 4.50),
    (3, 4.50, 5.00);

-- Insert mock inventory transactions
INSERT INTO inventory_transactions (inventory_id, quantity_change, reason) VALUES
    (1, -2, 'Used for espresso orders'),
    (2, -5, 'Used for cappuccino orders'),
    (3, -1, 'Used for vanilla latte orders');
