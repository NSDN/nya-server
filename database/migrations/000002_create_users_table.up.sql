CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS users(
  uid         UUID PRIMARY KEY DEFAULT gen_random_uuid(),   -- 用户编号
  username    VARCHAR(64) NOT NULL UNIQUE, -- 用户名，唯一
  password    TEXT NOT NULL,             -- 密码（明文哈希化后Base64编码）
  salt        TEXT NOT NULL,             -- 盐值（Base64编码）
  nickname    VARCHAR(128) NOT NULL,              -- 昵称
  user_group  VARCHAR(64) NOT NULL,               -- 用户组
  icon        TEXT                       -- 头像（可以存储URL或者Base64）
);

INSERT INTO users (
  username,
  password,
  salt,
  nickname,
  user_group
)
-- reimu:password
VALUES (
  'reimu',
  'JDJhJDEwJEZSek9BTnE5Q1dYaENuaFZDbjliSk9Vblh1dGNsSkVkNm5GTWZCZlQzMUlWWk5oM0UvTW55',
  'z6oaLurTZY8y3Jo8jS4zEQ==',
  '霊夢',
  '結界管理者'
)
RETURNING uid;

