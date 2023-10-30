CREATE TABLE IF NOT EXISTS users (
  id CHAR(36) PRIMARY KEY,
  interest_id CHAR(36) NOT NULL,
  email VARCHAR(50) UNIQUE,
  username VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  fullname VARCHAR(255) NOT NULL,
  status VARCHAR(10) NOT NULL,
  role CHAR(6) NOT NULL,
  bio VARCHAR(255),
  phone CHAR(15),
  birth_date DATE,
  profile_picture VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (interest_id) REFERENCES interests(id)
);