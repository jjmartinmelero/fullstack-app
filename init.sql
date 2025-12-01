-- Create Categories table
CREATE TABLE IF NOT EXISTS categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(100)
);

-- Create Recipes table
CREATE TABLE IF NOT EXISTS recipes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    category_id INT,
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(100),
    time VARCHAR(100) NOT NULL,
    photo VARCHAR(100) NOT NULL,
    description TEXT,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
);

-- Create Contacts table
CREATE TABLE IF NOT EXISTS contacts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(100) NOT NULL,
    message TEXT NOT NULL,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert sample categories
INSERT INTO categories (name, slug) VALUES
('Desserts', 'desserts'),
('Main Course', 'main-course'),
('Appetizers', 'appetizers'),
('Salads', 'salads');

-- Insert sample recipes
INSERT INTO recipes (category_id, name, slug, time, photo, description, date) VALUES
(2, 'Pasta Carbonara', 'pasta-carbonara', '30 min', 'carbonara.jpg', 'Classic Italian pasta with eggs, pancetta, and parmesan cheese', NOW()),
(2, 'Chicken Curry', 'chicken-curry', '45 min', 'curry.jpg', 'Spicy Indian curry with coconut milk and aromatic spices', NOW()),
(4, 'Caesar Salad', 'caesar-salad', '15 min', 'salad.jpg', 'Fresh romaine lettuce with caesar dressing and croutons', NOW()),
(1, 'Chocolate Cake', 'chocolate-cake', '60 min', 'cake.jpg', 'Rich and moist chocolate cake with ganache frosting', NOW()),
(3, 'Bruschetta', 'bruschetta', '20 min', 'bruschetta.jpg', 'Toasted bread topped with fresh tomatoes and basil', NOW());

-- Insert sample contacts
INSERT INTO contacts (name, email, phone, message, date) VALUES
('John Doe', 'john@example.com', '+1234567890', 'Great recipes! Would love to see more vegetarian options.', NOW()),
('Jane Smith', 'jane@example.com', '+0987654321', 'The carbonara recipe was amazing. Thanks for sharing!', NOW());
