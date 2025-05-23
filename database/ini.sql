CREATE DATABASE IF NOT EXISTS live;
USE live;

DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
  `id`          INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id`     VARCHAR(256) UNIQUE NOT NULL,
  `user_name`   VARCHAR(256) NOT NULL,
  `img`         VARCHAR(256) NULL,
  `introduce`   VARCHAR(256) NULL DEFAULT "",
  `password`    VARCHAR(256) NOT NULL,
  `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at`  TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `posts`;
CREATE TABLE IF NOT EXISTS `posts` (
  `id`          INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id`     VARCHAR(256) NOT NULL,
  `file`        VARCHAR(256) NOT NULL,              ### ex) dir名/file名
  `introduce`   VARCHAR(90) NOT NULL DEFAULT "",    ### 30文字以内->30×3=90
  `type`        ENUM('short', 'long'),
  `good`        INT UNSIGNED NOT NULL DEFAULT 0,
  `watch`       INT UNSIGNED NOT NULL DEFAULT 0,
  `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at`  TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX(`user_id` (`user_id`))
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `follows`;
CREATE TABLE IF NOT EXISTS `follows` (
  `id`                   INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `following_user_id`    VARCHAR(256) NOT NULL,
  `followed_user_id`     VARCHAR(256) NOT NULL,
  `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at`  TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX(`following` (`following_user_id`)),
  INDEX(`followed`  (`following_user_id`))
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
