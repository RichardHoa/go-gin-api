INSERT INTO products (name, description, image, price, quantity, createdAt) VALUES
  ('Wireless Headphones', 'Noise-canceling over-ear headphones with Bluetooth connectivity.', 'headphones.jpg', 89.99, 50, NOW()),
  ('Smartphone 12', 'Latest smartphone with a 6.5-inch OLED display and a 64MP camera.', 'smartphone12.jpg', 699.99, 150, DATE_ADD(NOW(), INTERVAL 10 MINUTE)),
  ('Gaming Laptop', 'High-performance gaming laptop with 16GB RAM and an RTX 3060 GPU.', 'gaming_laptop.jpg', 1299.99, 30, DATE_ADD(NOW(), INTERVAL 20 MINUTE)),
  ('4K UHD TV', '55-inch 4K UHD TV with HDR and smart functionality.', '4k_tv.jpg', 499.99, 20, DATE_ADD(NOW(), INTERVAL 30 MINUTE)),
  ('Smartwatch', 'Fitness tracking smartwatch with heart rate monitor and sleep tracking.', 'smartwatch.jpg', 199.99, 100, DATE_ADD(NOW(), INTERVAL 40 MINUTE));
