CREATE TYPE ORDER_STATUS AS ENUM ('pending', 'processing', 'completed', 'cancelled');
CREATE TYPE PAYMENT_METHOD AS ENUM ('cash', 'card', 'online');
CREATE TYPE STAFF_ROLE AS ENUM ('barista', 'cashier', 'manager');

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_name TEXT NOT NULL,
    status ORDER_STATUS NOT NULL DEFAULT 'pending',
    total_amount DECIMAL(10,2) NOT NULL CHECK (total_amount >= 0),
    payment_method PAYMENT_METHOD NOT NULL,
    special_instructions JSONB DEFAULT '{}',
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ 
);

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,
    menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity > 0),
    price_at_order DECIMAL(10,2) NOT NULL CHECK (price_at_order >= 0)
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
    menu_item_id INT REFERENCES menu_items(id) NOT NULL ON DELETE CASCADE,
    ingredient_id INT REFERENCES inventory(id) NOT NULL ON DELETE CASCADE,
    quantity_used DECIMAL(10,2) NOT NULL CHECK (quantity_used > 0),
    PRIMARY KEY (menu_item_id, ingredient_id)
);

CREATE TABLE inventory (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    quantity DECIMAL(10,2) NOT NULL CHECK (quantity >= 0),
    unit_type TEXT NOT NULL,
    last_updated TIMESTAMPTZ DEFAULT NOW()
);


CREATE TABLE order_status_history (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) NOT NULL ON DELETE CASCADE,
    status ORDER_STATUS NOT NULL,
    changed_at TIMESTAMPTZ DEFAULT NOW()
);


CREATE TABLE price_history (
    id SERIAL PRIMARY KEY,
    menu_item_id INT REFERENCES menu_items(id) NOT NULL ON DELETE CASCADE,
    old_price DECIMAL(10,2) NOT NULL CHECK (old_price >= 0),
    new_price DECIMAL(10,2) NOT NULL CHECK (new_price >= 0),
    changed_at TIMESTAMPTZ DEFAULT NOW()
);


CREATE TABLE inventory_transactions (
    id SERIAL PRIMARY KEY,
    inventory_id INT REFERENCES inventory(id) NOT NULL ON DELETE CASCADE,
    quantity_change DECIMAL(10,2) NOT NULL,
    reason TEXT NOT NULL,
    transaction_time TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE staff (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    role STAFF_ROLE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
)