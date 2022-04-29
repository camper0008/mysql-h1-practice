DROP TABLE IF EXISTS opg00;

CREATE TABLE opg00 (
   id INT UNSIGNED UNIQUE PRIMARY KEY,
   name TINYTEXT,
   age TINYINT UNSIGNED
);

INSERT INTO opg00 (id, name, age)
VALUES
(1,  "dietrich", 127),
(2,  "dieter",   255),
(3,  "dddd",     50),
(4,  "eeeee",    40),
(5,  "dimon",    40),
(6,  "dietz",    49),
(7,  "deeeee",   40),
(8,  "dennis",   254),
(9,  "dddd",     50);
