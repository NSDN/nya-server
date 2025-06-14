CREATE TABLE IF NOT EXISTS topics(
  id BIGSERIAL PRIMARY KEY,
  author_id BIGINT NOT NULL REFERENCES users(id),
  plate_id VARCHAR(100) NOT NULL REFERENCES plates(id),
  title TEXT NOT NULL,
  topic_type VARCHAR(20) NOT NULL CHECK (topic_type IN ('comic', 'article')),
  thumbnail_link TEXT NOT NULL,
  tag VARCHAR(20)[],
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

INSERT INTO topics (
  author_id,
  plate_id,
  title,
  topic_type,
  thumbnail_link,
  tag,
  created_at,
  updated_at
) VALUES
  (
    1,
    'chat',
    '青年在选择职业时的考虑',
    'article',
    '',
    ARRAY['马克思'],
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  ),
  (
    1, 
    'chat', 
    '德意志意识形态', 
    'article',
    '',
    ARRAY['马克思', '恩格斯'], 
    CURRENT_TIMESTAMP, 
    CURRENT_TIMESTAMP
  ),
  (
    1, 
    'chat', 
    '共产党宣言', 
    'article',
    '',
    ARRAY['马克思', '陈望道'], 
    CURRENT_TIMESTAMP, 
    CURRENT_TIMESTAMP
  );

