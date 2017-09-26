-- noinspection SqlNoDataSourceInspectionForFile

-- Создание основных таблиц и их связей --

-- +migrate Up

-- [Table]  --
CREATE TABLE  (
  id           SERIAL       NOT NULL PRIMARY KEY,
  user_id      INTEGER      NOT NULL,
  chat_id      INTEGER      NOT NULL UNIQUE,
  secret_token VARCHAR(255) NOT NULL UNIQUE,
  created_at   TIMESTAMP
);

-- [Table] notification_types - references of notification types --
CREATE TABLE notification_types (
  id          SERIAL       NOT NULL PRIMARY KEY,
  title       VARCHAR(255) NOT NULL UNIQUE,
  description VARCHAR(255) NOT NULL,
  img_url     VARCHAR(255),
  moderated   BOOLEAN      NOT NULL DEFAULT FALSE
);

-- [INSERT] notification types reference --
INSERT INTO notification_types (title, description, img_url, moderated)
VALUES (
  'error', 'Error messages which stop app running',
  'https://cdn.pixabay.com/photo/2017/02/12/21/29/false-2061131_960_720.png', TRUE);

INSERT INTO notification_types (title, description, img_url, moderated)
VALUES (
  'info', 'Information messages like load progress, start and finish process, etc',
  'https://upload.wikimedia.org/wikipedia/commons/6/66/Info_groen.png', TRUE
);

INSERT INTO notification_types (title, description, img_url, moderated)
VALUES (
  'warn', 'Potentially harmful situations',
  'https://upload.wikimedia.org/wikipedia/commons/5/57/Circle-style-warning.svg', TRUE
);

INSERT INTO notification_types (title, description, img_url, moderated)
VALUES (
  'fatal', 'Critical application error that will presumably lead the application to abort.',
  'https://upload.wikimedia.org/wikipedia/commons/5/57/Circle-style-warning.svg', TRUE
);

-- [Table] notifications - registrate notifications --
CREATE TABLE notifications (
  id           SERIAL       NOT NULL PRIMARY KEY,
  subscribe_id INTEGER      NOT NULL PRIMARY KEY,
  type_id      INTEGER      NOT NULL,
  title        VARCHAR(255) NOT NULL,
  text         VARCHAR(255),
  tags         VARCHAR(255),
  FOREIGN KEY (subscribe_id) REFERENCES  (id),
  FOREIGN KEY (type_id) REFERENCES notification_types (id)
);

-- +migrate Down
-- DROP TABLE  CASCADE;
-- DROP TABLE notification_types CASCADE;
-- DROP TABLE notifications CASCADE;
