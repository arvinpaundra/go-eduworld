CREATE TYPE level_course AS ENUM ('beginner', 'intermediate', 'advanced');
CREATE TYPE type_course AS ENUM ('premium', 'free');

CREATE TABLE IF NOT EXISTS courses (
  id CHAR(36) PRIMARY KEY,
  category_id CHAR(36) NOT NULL,
  user_id CHAR(36) NOT NULL,
  interest_id CHAR(36) NOT NULL,
  title VARCHAR(255) NOT NULL,
  type type_course NOT NULL,
  level level_course NOT NULL,
  description TEXT,
  thumbnail VARCHAR(255),
  is_published BOOLEAN,
  price BIGINT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (category_id) REFERENCES categories(id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (interest_id) REFERENCES interests(id)
);