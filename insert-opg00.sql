DROP TABLE IF EXISTS opg00;

CREATE TABLE opg00 (
   id INT UNSIGNED UNIQUE PRIMARY KEY,
   name TINYINT,
   age TINYINT UNSIGNED
);

SELECT * FROM opg00;

INSERT INTO opg00 (id, name, age)
VALUES (1, "dietrich", 127);

INSERT INTO opg00 (id, name, age)
VALUES (2, "dieter", 255);

INSERT INTO opg00 (id, name, age)
VALUES (3, "dddd", 50);

INSERT INTO opg00 (id, name, age)
VALUES (4, "eeeee", 40);

INSERT INTO opg00 (id, name, age)
VALUES (5, "dimon", 40);

INSERT INTO opg00 (id, name, age)
VALUES (6, "dietz", 49);

INSERT INTO opg00 (id, name, age)
VALUES (7, "deeeee", 40);

INSERT INTO opg00 (id, name, age)
VALUES (8, "dennis", 254);

INSERT INTO opg00 (id, name, age)
VALUES (9, "dddd", 50);
