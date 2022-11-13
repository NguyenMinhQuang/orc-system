CREATE TABLE IF NOT EXISTS `user_info`
(
    `id`        INT         NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `user_name` varchar(20) NOT NULL,
    `sex`       varchar(50) NOT NULL,
    `addr`      varchar(50) NOT NULL,
    `create_at` TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE = INNODB;