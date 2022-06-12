
-- +migrate Up
CREATE TABLE IF NOT EXISTS users 
(
  id SERIAL NOT NULL PRIMARY KEY,
  username TEXT NOT NULL,
  email TEXT NOT NULL,
  password TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ
);


-- +migrate Down
DROP TABLE IF EXISTS `users`;
