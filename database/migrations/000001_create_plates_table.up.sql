CREATE TABLE IF NOT EXISTS plates(
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  description TEXT NOT NULL,
  background TEXT NOT NULL,
  page_type VARCHAR(20) NOT NULL,
  sort_order SMALLINT UNIQUE NOT NULL CHECK (sort_order >= 0)
);

INSERT INTO plates (
  id,
  name,
  description,
  background,
  page_type,
  sort_order
)
VALUES
  (
    'localization',
    '喵玉汉化馆',
    '',
    'https://i.imgur.com/ohQuzivl.jpg',
    'comic',
    0
  ),
  (
    'music',
    '喵玉咏唱组',
    '',
    'https://i.imgur.com/IHo7tTyl.jpg',
    'article',
    1
  ),
  (
    'chat',
    '魔女的茶会',
    '',
    'https://i.imgur.com/JsWkJ4jl.jpg',
    'article',
    2
  );
