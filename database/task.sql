CREATE TABLE `task`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `userid`     bigint unsigned NOT NULL  COMMENT 'User id',
    `adress`     varchar(128) NOT NULL DEFAULT '' COMMENT 'wallet address',
    `token`      varchar(128) NOT NULL DEFAULT '' COMMENT 'Token',
    `recipients` text  COMMENT 'recipients list',
    `amounts`    text  COMMENT 'amounts list',
    `status`     tinyint unsigned NOT NULL DEFAULT 0 COMMENT 'Task status',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'task create time',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'task update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'task delete time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_userid` (`userid`) COMMENT 'Userid  index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='task table'