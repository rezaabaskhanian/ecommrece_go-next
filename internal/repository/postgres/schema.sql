CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    phone_number TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    avatar_url TEXT,
    role TEXT DEFAULT 'user',
    created_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE products (
    id SERIAL PRIMARY KEY,

    shop_id INT,  -- اگر بعداً جدول shops داشتی می‌تونی FK کنی

    name VARCHAR(255) NOT NULL,

    description TEXT,

    price NUMERIC(12,2) NOT NULL,

    stock INT NOT NULL DEFAULT 0,

    category VARCHAR(100),

    image_url TEXT,

    created_at TIMESTAMP  DEFAULT NOW(),

    updated_at TIMESTAMP
);

