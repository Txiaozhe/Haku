CREATE DATABASE IF NOT EXISTS content;
SET DATABASE = content;

CREATE TABLE IF NOT EXISTS blog (
  id        SMALLSERIAL    PRIMARY KEY,
  title     STRING(32)     NOT NULL,
  category  INT8           NOT NULL,
  abstract  STRING(64)     NOT NULL,
  tag       STRING(64)     NOT NULL,
  contentid INT64          NOT NULL,
  star      INT16          DEFAULT 0,
  created   TIMESTAMP      DEFAULT now()
);
