-- Adminer 4.8.1 MySQL 8.0.29 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

INSERT INTO `opg00` (`id`, `name`, `age`) VALUES
(1,	'dietrich',	127),
(2,	'dieter',	255),
(3,	'dddd',	50),
(4,	'eeeee',	40),
(5,	'dimon',	40),
(6,	'dietz',	49),
(7,	'deeeee',	40),
(8,	'dennis',	254),
(9,	'dddd',	50),
(10,	'dennis',	233),
(11,	'donald',	18);

DROP TABLE IF EXISTS `opg01_comments`;
CREATE TABLE `opg01_comments` (
  `id` int NOT NULL,
  `content` text NOT NULL,
  `FK_post` int NOT NULL,
  `FK_user` int NOT NULL,
  UNIQUE KEY `id` (`id`),
  KEY `comment_id` (`id`),
  KEY `post` (`FK_post`),
  KEY `author` (`FK_user`),
  CONSTRAINT `opg01_comments_ibfk_1` FOREIGN KEY (`FK_post`) REFERENCES `opg01_posts` (`id`),
  CONSTRAINT `opg01_comments_ibfk_2` FOREIGN KEY (`FK_user`) REFERENCES `opg01_users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `opg01_comments` (`id`, `content`, `FK_post`, `FK_user`) VALUES
(0,	'har du fået lov af helledie til at lave den post?',	0,	1),
(1,	'@soelberg\r\nshjøøøør\r\nMvh Dietz',	0,	0);

DROP TABLE IF EXISTS `opg01_friends`;
CREATE TABLE `opg01_friends` (
  `id` int unsigned NOT NULL,
  `FK_user0` int NOT NULL,
  `FK_user1` int NOT NULL,
  KEY `user0` (`FK_user0`),
  KEY `user1` (`FK_user1`),
  CONSTRAINT `opg01_friends_ibfk_1` FOREIGN KEY (`FK_user0`) REFERENCES `opg01_users` (`id`),
  CONSTRAINT `opg01_friends_ibfk_2` FOREIGN KEY (`FK_user1`) REFERENCES `opg01_users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `opg01_friends` (`id`, `FK_user0`, `FK_user1`) VALUES
(0,	1,	2),
(1,	1,	0);

DROP TABLE IF EXISTS `opg01_posts`;
CREATE TABLE `opg01_posts` (
  `id` int NOT NULL,
  `title` text NOT NULL,
  `content` text NOT NULL,
  `FK_user` int NOT NULL,
  UNIQUE KEY `id` (`id`),
  KEY `post_id` (`id`),
  KEY `FK_user` (`FK_user`),
  CONSTRAINT `opg01_posts_ibfk_1` FOREIGN KEY (`FK_user`) REFERENCES `opg01_users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `opg01_posts` (`id`, `title`, `content`, `FK_user`) VALUES
(0,	'post #0',	'hello and welcome to my programming tutorial\r\n\r\ntoday we will learn how to use rust',	0);

DROP TABLE IF EXISTS `opg01_users`;
CREATE TABLE `opg01_users` (
  `id` int NOT NULL,
  `name` text NOT NULL,
  `permission` int NOT NULL,
  UNIQUE KEY `id` (`id`),
  KEY `user_id` (`id`),
  CONSTRAINT `opg01_users_ibfk_1` FOREIGN KEY (`id`) REFERENCES `opg01_users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `opg01_users` (`id`, `name`, `permission`) VALUES
(0,	'dietz dieter',	1),
(1,	'ole soelberg',	2),
(2,	'ole helledie',	2);

-- 2022-05-04 12:36:17
