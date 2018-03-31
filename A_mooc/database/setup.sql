DROP TABLE courses CASCADE;
DROP TABLE users CASCADE;

CREATE TABLE courses (
  ID        INTEGER PRIMARY KEY,
  title     TEXT,
  content   TEXT,
  host      TEXT,
  url       TEXT
);

CREATE TABLE users (
  ID        INTEGER PRIMARY KEY,
  email     VARCHAR(64) NOT NULL UNIQUE,
  password  TEXT
);

--INSERT INTO courses (ID, title, content) VALUES (1, '1title', '1text');
--INSERT INTO courses (ID, title, content) VALUES (2, '2title', '2text');

--DELETE FROM courses WHERE ID = 2