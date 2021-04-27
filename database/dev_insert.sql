### this sql is for development or local.

USE live IF EXISTS live;

TRUNCATE TABLE `users`;
INSERT INTO `users`
(`user_id`, `user_name`, `img`, `introduce`, `password`) VALUES
("maaaaakoto", "makoto", "", "introduce, introduce, introduce", "pass"),
("nanashi01", "nanashi1", "", "introduce, introduce, introduce", "pass"),
("nanashi02", "nanashi2", "", "", "pass");

TRUNCATE TABLE`posts`;
INSERT INTO `posts`
(`user_id`, `file`, `introduce`) VALUES
("maaaaakoto", "sample/sample.mp3", "introduce, introduce, introduce");

TRUNCATE TABLE `follows`;
INSERT INTO `follows`
(`following_user_id`, `followed_user_id`) VALUES
("maaaaakoto", "nanashi01"),
("nanashi02", "maaaaakoto");