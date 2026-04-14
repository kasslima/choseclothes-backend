CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    category VARCHAR(100),
    image_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sources (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL -- ex: aliexpress, mercadolivre, shein
);

CREATE TABLE listings (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    source_id INT NOT NULL,

    external_id VARCHAR(100) NOT NULL,
    title VARCHAR(255) NOT NULL,
    seller_name VARCHAR(255),
    seller_rating DECIMAL(3,2),

    product_url TEXT NOT NULL,
    currency VARCHAR(10) DEFAULT 'USD',
    is_active BOOLEAN DEFAULT TRUE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    FOREIGN KEY (source_id) REFERENCES sources(id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX unique_listing_source 
ON listings (source_id, external_id);

CREATE TABLE price_history (
    id SERIAL PRIMARY KEY,
    listing_id INT NOT NULL,

    price DECIMAL(10,2) NOT NULL,
    discount_price DECIMAL(10,2),

    recorded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (listing_id) REFERENCES listings(id) ON DELETE CASCADE
);

CREATE TABLE channels (
    id SERIAL PRIMARY KEY,

    type VARCHAR(20) NOT NULL, -- telegram | whatsapp | email | web
    external_id VARCHAR(255) NOT NULL, 

    name VARCHAR(255),

    is_active BOOLEAN DEFAULT TRUE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(type, external_id)
);

CREATE TABLE price_alerts (
    id SERIAL PRIMARY KEY,

    channel_id INT NOT NULL,
    listing_id INT NOT NULL,

    target_price DECIMAL(10,2),

    is_active BOOLEAN DEFAULT TRUE,
    last_notified_at TIMESTAMP,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (group_id) REFERENCES telegram_groups(id) ON DELETE CASCADE,
    FOREIGN KEY (listing_id) REFERENCES listings(id) ON DELETE CASCADE
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,

    google_id VARCHAR(255) UNIQUE,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT,
    name VARCHAR(255),
    profile_picture TEXT,

    is_subscriber BOOLEAN DEFAULT FALSE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_listing_product ON listings(product_id);
CREATE INDEX idx_listing_source ON listings(source_id);

CREATE INDEX idx_price_listing ON price_history(listing_id);
CREATE INDEX idx_price_recorded ON price_history(recorded_at);

CREATE INDEX idx_alert_listing ON price_alerts(listing_id);
CREATE INDEX idx_alert_channel ON price_alerts(channel_id);