DROP TABLE IF EXISTS opg00;

CREATE TABLE opg00 (
   id INT UNSIGNED UNIQUE PRIMARY KEY NOT NULL,
   name TINYTEXT NOT NULL,
   age TINYINT UNSIGNED NOT NULL
);

INSERT INTO opg00 (id, name, age)
VALUES
(1,  "dietrich", 127),
(2,  "dieter",   255),
(3,  "donald",   50),
(4,  "erik",     40),
(4,  "erik",     50),
(5,  "dimon",    40),
(6,  "dietz",    49),
(7,  "deeeee",   40),
(8,  "dennis",   254);
