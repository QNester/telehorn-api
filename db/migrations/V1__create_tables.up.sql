-- Создание основных таблиц и их связей --

-- +migrate Up

-- [Table] Users (project managers) --
CREATE TABLE subscribers(
  id                 SERIAL NOT NULL PRIMARY KEY,
  user_id            INTEGER NOT NULL,
  chat_id            INTEGER NOT NULL UNIQUE,
  scret_token        VARCHAR(255) NOT NULL UNIQUE,
  created_at         TIMESTAMP
);


-- +migrate Down
-- DROP TABLE subscribers CASCADE;
