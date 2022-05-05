-- Adminer 4.8.1 MySQL 8.0.29 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `opg01_comments`;
CREATE TABLE `opg01_comments` (
  `id` int unsigned NOT NULL,
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


DROP TABLE IF EXISTS `opg01_friends`;
CREATE TABLE `opg01_friends` (
  `id` int unsigned NOT NULL,
  `FK_user0` int unsigned NOT NULL,
  `FK_user1` int unsigned NOT NULL,
  UNIQUE KEY `id` (`id`),
  KEY `user0` (`FK_user0`),
  KEY `user1` (`FK_user1`),
  CONSTRAINT `opg01_friends_ibfk_1` FOREIGN KEY (`FK_user0`) REFERENCES `opg01_users` (`id`),
  CONSTRAINT `opg01_friends_ibfk_2` FOREIGN KEY (`FK_user1`) REFERENCES `opg01_users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `opg01_friends` (`id`, `FK_user0`, `FK_user1`) VALUES
(0,	0,	1),
(1,	1,	2);

DROP TABLE IF EXISTS `opg01_posts`;
CREATE TABLE `opg01_posts` (
  `id` int unsigned NOT NULL,
  `title` text NOT NULL,
  `content` text NOT NULL,
  `FK_user` int unsigned NOT NULL,
  UNIQUE KEY `id` (`id`),
  KEY `post_id` (`id`),
  KEY `FK_user` (`FK_user`),
  CONSTRAINT `opg01_posts_ibfk_1` FOREIGN KEY (`FK_user`) REFERENCES `opg01_users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


DROP TABLE IF EXISTS `opg01_users`;
CREATE TABLE `opg01_users` (
  `id` int unsigned NOT NULL,
  `name` text NOT NULL,
  `permission` int NOT NULL,
  UNIQUE KEY `id` (`id`),
  KEY `user_id` (`id`),
  CONSTRAINT `opg01_users_ibfk_1` FOREIGN KEY (`id`) REFERENCES `opg01_users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- 2022-05-04 12:44:42
