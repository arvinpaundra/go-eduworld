CREATE TABLE IF NOT EXISTS modules (
  id CHAR(36) PRIMARY KEY,
  course_id CHAR(36) NOT NULL,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (course_id) REFERENCES courses(id)
);