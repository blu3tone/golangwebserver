CREATE TABLE IF NOT EXISTS user_books(
  id VARCHAR(100) PRIMARY KEY UNIQUE,
  user_id VARCHAR(100),
  book_id VARCHAR(100)
);