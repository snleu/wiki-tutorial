CREATE TABLE pages (
  title    char(14) NOT NULL,
  body   TEXT
);

INSERT INTO pages (title, body) VALUES
('test', 'Hello world!');

ALTER TABLE pages ADD PRIMARY KEY (title);