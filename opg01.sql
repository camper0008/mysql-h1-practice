-- username and id joined with comment
-- ordered by comment id
SELECT opg01_comments.id, opg01_comments.content, opg01_users.id, opg01_users.name
FROM opg01_comments
JOIN opg01_users ON (opg01_comments.FK_user = opg01_users.id)
ORDER BY opg01_comments.id;

-- usernames of friends
SELECT users0.name, users1.name
FROM opg01_friends
JOIN opg01_users as users0 ON (opg01_friends.FK_user0 = users0.id)
JOIN opg01_users as users1 ON (opg01_friends.FK_user1 = users1.id);

-- drop view if exists
DROP VIEW IF EXISTS named_friend_relations;

-- create view that displays names of friends
CREATE VIEW named_friend_relations AS
SELECT users0.name as name0, users1.name as name1
FROM opg01_friends
JOIN opg01_users as users0 ON (opg01_friends.FK_user0 = users0.id)
JOIN opg01_users as users1 ON (opg01_friends.FK_user1 = users1.id);
