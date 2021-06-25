
-- +migrate Up
DROP TABLE IF EXISTS `users` ;
CREATE TABLE IF NOT EXISTS `users` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(100) NOT NULL,
    `password` VARCHAR(100) NOT NULL,
    `full_name` VARCHAR(100) NOT NULL,
    PRIMARY KEY (`id`)
    )ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
-- +migrate Down
DROP TABLE IF EXISTS `users` ;
