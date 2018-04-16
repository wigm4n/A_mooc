DROP TABLE users CASCADE;
DROP TABLE subscriptions CASCADE;
DROP TABLE sessions CASCADE;

CREATE TABLE users (
  ID        INTEGER PRIMARY KEY,
  email     VARCHAR(64) NOT NULL UNIQUE,
  password  TEXT
);

CREATE TABLE sessions (
  userID    INTEGER REFERENCES users(ID)
);

CREATE TABLE subscriptions (
  id              INTEGER PRIMARY KEY,
  userID          INTEGER REFERENCES users(ID),
  teg             TEXT,
  platforms       TEXT,
  languages       TEXT,
  levels          TEXT,
  availabilities  TEXT,
  date            TIMESTAMP,
  frequency       INTEGER
);

--INSERT INTO courses (ID, title, content) VALUES (1, '1title', '1text');
--INSERT INTO courses (ID, title, content) VALUES (2, '2title', '2text');

DELETE FROM subscriptions WHERE ID = 2