CREATE DATABASE IF NOT EXISTS grade;

SET DATABASE = grade;

CREATE TABLE IF NOT EXISTS badge (
  id       SMALLSERIAL   PRIMARY KEY,
  blogid   INT64         NOT NULL,
  name     STRING(32)    DEFAULT 'Anonymous',
  avatar   STRING(16)    NOT NULL,
  content  STRING(64)    NOT NULL,
  created     TIMESTAMP  DEFAULT now()
);
