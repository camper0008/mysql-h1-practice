-- all
SELECT * FROM opg00
ORDER BY Id;

-- select distinct ages where name starts with d, 
-- is either >= 50 or < 45
-- and is not a multiple of 127
SELECT DISTINCT age FROM opg00 
WHERE name LIKE "D%" 
AND (age >= 50 OR age < 45)
AND (age % 127) <> 0;

-- select people with duplicate ages
SELECT age, COUNT(age) as occurances FROM opg00 
GROUP BY age
HAVING occurances > 1;

-- select * from people with duplicate ages
SELECT * FROM opg00
WHERE age IN (
  SELECT t.age FROM (
    SELECT age, COUNT(age) as occurances FROM opg00 
    GROUP BY age
    HAVING occurances > 1
  ) AS t
);

-- rank people by how long their names are
SELECT *, LENGTH(name) as name_length FROM opg00
ORDER BY name_length DESC;

-- average age of people with a certain name
SELECT name, AVG(age) as average_age FROM opg00
GROUP BY name
