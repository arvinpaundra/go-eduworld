CREATE TABLE IF NOT EXISTS interests (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  description TEXT,
  icon VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);