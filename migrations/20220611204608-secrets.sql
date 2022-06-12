
-- +migrate Up
CREATE TABLE IF NOT EXISTS secrets 
(
  id SERIAL NOT NULL PRIMARY KEY,
  user_id integer references users(id),
  description TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL
);


-- +migrate Down
DROP TABLE IF EXISTS `secrets`;
