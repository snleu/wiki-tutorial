DROP TABLE pages;

CREATE TABLE pages (
  title    char(20) NOT NULL,
  body     text
);

INSERT INTO pages (title, body) VALUES
('test', 'Hello world!'),
('ABlankPage', '');

ALTER TABLE pages ADD PRIMARY KEY (title);