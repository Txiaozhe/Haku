CREATE DATABASE IF NOT EXISTS core;
SET DATABASE = core;

CREATE TABLE IF NOT EXISTS admin (
  id          SMALLSERIAL   PRIMARY KEY,
  name        STRING(32)    UNIQUE NOT NULL,
  pass        STRING(256)   NOT NULL,
  github      STRING(32)    UNIQUE NOT NULL,
  email       STRING(64)    UNIQUE,
  created     TIMESTAMP     DEFAULT now()
);
