CREATE TABLE IF NOT EXISTS plates(
  id VARCHAR (100) PRIMARY KEY,
  name VARCHAR (100) NOT NULL,
  description TEXT,
  background TEXT,
  page_type VARCHAR (20) NOT NULL
);
