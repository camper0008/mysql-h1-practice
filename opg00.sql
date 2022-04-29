-- all
SELECT * FROM opg00
ORDER BY Id;

-- select distinct ages where name starts with d, 
-- is either >= 50 or < 45
-- and is not a multiple of 127
SELECT DISTINCT Age FROM opg00 
WHERE Name LIKE "D%" 
AND (Age >= 50 OR Age < 45)
AND (Age % 127) <> 0;

-- select people with duplicate ages
SELECT Age, COUNT(Age) FROM opg00 
GROUP BY Age
HAVING COUNT(Age) > 1;

-- select * from people with duplicate ages
SELECT * FROM opg00
WHERE Age IN (
  SELECT t.Age FROM (
    SELECT Age, COUNT(Age) FROM opg00 
    GROUP BY Age
    HAVING COUNT(Age) > 1
  ) t
);

-- rank people by how long their names are
SELECT *, LENGTH(Name) FROM opg00
ORDER BY LENGTH(Name) DESC;
