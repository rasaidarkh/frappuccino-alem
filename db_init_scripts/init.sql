CREATE TYPE ORDER_STATUS AS ENUM ('pending', 'processing', 'completed', 'cancelled');
CREATE TYPE PAYMENT_METHOD AS ENUM ('cash', 'card', 'online');
CREATE TYPE STAFF_ROLE AS ENUM ('barista', 'cashier', 'manager');

CREATE TABLE inventory (
    id SERIAL PRIMARY KEY,
    item_name TEXT NOT NULL,
    quantity DECIMAL(10,2) NOT NULL CHECK (quantity >= 0),
    unit TEXT NOT NULL,
    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),  -- NEW COLUMN
    created_at TIMESTAMPTZ DEFAULT NOW(), 
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE menu_items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    price DECIMAL(15,2) NOT NULL CHECK (price >= 0),
    categories TEXT[] DEFAULT '{}',
    allergens TEXT[] DEFAULT '{}',
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMPTZ DEFAULT NOW(), 
    updated_at TIMESTAMPTZ DEFAULT NOW()
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
    total_amount DECIMAL(15,2) NOT NULL CHECK (total_amount >= 0),
    payment_method PAYMENT_METHOD NOT NULL,
    special_instructions JSONB DEFAULT '{}',
    created_at TIMESTAMPTZ DEFAULT NOW(), 
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    search_vector tsvector
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
    inventory_id INT REFERENCES inventory(id) ON DELETE CASCADE  NOT NULL,
    quantity_change DECIMAL(10,2) NOT NULL,
    reason TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE staff (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    role STAFF_ROLE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(), 
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

UPDATE orders 
SET search_vector = 
    setweight(to_tsvector('english', COALESCE(customer_name, '')), 'A') ||
    setweight((
        SELECT to_tsvector('english', COALESCE(string_agg(mi.name, ' '), ''))
        FROM order_items oi
        JOIN menu_items mi ON oi.menu_item_id = mi.id
        WHERE oi.order_id = orders.id
    ), 'B');

ALTER TABLE menu_items 
ADD COLUMN search_vector tsvector GENERATED ALWAYS AS (
    setweight(to_tsvector('english', name), 'A') ||
    setweight(to_tsvector('english', description), 'B')
) STORED;

-- Order search vector function
CREATE OR REPLACE FUNCTION update_order_search_vector()
RETURNS TRIGGER AS $$
BEGIN
    NEW.search_vector := 
        setweight(to_tsvector('english', COALESCE(NEW.customer_name, '')), 'A') ||
        setweight((
            SELECT to_tsvector('english', 
                COALESCE(string_agg(mi.name || ' ' || COALESCE(mi.description, ''), ' '), '')
            )
            FROM order_items oi
            JOIN menu_items mi ON oi.menu_item_id = mi.id
            WHERE oi.order_id = NEW.id
        ), 'B');
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- Triggers
CREATE TRIGGER order_search_update
BEFORE INSERT OR UPDATE ON orders
FOR EACH ROW EXECUTE FUNCTION update_order_search_vector();

CREATE OR REPLACE FUNCTION refresh_order_search_vector()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE orders SET updated_at = NOW()
    WHERE id = COALESCE(NEW.order_id, OLD.order_id);
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER order_items_search_update
AFTER INSERT OR UPDATE OR DELETE ON order_items
FOR EACH ROW EXECUTE FUNCTION refresh_order_search_vector();

-- Indexes
CREATE INDEX idx_orders_search ON orders USING GIN(search_vector);
CREATE INDEX idx_menu_items_search ON menu_items USING GIN(search_vector);

-- Initialize existing data
UPDATE orders 
SET search_vector = 
    setweight(to_tsvector('english', COALESCE(customer_name, '')), 'A') ||
    setweight((
        SELECT to_tsvector('english', 
            COALESCE(string_agg(mi.name || ' ' || COALESCE(mi.description, ''), ' '), '')
        )
        FROM order_items oi
        JOIN menu_items mi ON oi.menu_item_id = mi.id
        WHERE oi.order_id = orders.id
    ), 'B');
