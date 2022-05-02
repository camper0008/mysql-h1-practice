-- Adminer 4.8.1 MySQL 8.0.29 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

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

-- 2022-05-02 12:54:57
