CREATE DATABASE IF NOT EXISTS zeniba;
SET DATABASE = zeniba;

CREATE TABLE IF NOT EXISTS blog (
  id        SMALLSERIAL    PRIMARY KEY,
  title     STRING(32)     NOT NULL,
  category  STRING(16)     NOT NULL,
  content   STRING(256)    NOT NULL,
  created   TIMESTAMP      DEFAULT now()
);
