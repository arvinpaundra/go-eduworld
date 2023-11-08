CREATE TABLE IF NOT EXISTS categories (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  description TEXT,
  image VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);