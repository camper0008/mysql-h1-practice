-- Adminer 4.8.1 MySQL 8.0.29 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `opg00`;
CREATE TABLE `opg00` (
  `id` int unsigned NOT NULL,
  `name` tinytext NOT NULL,
  `age` tinyint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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

-- 2022-05-02 12:55:56
