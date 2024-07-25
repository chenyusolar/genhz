CREATE TABLE `address`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `address`    varchar(128) NOT NULL DEFAULT '' COMMENT 'Address',
    `address_type` varchar(128) NOT NULL DEFAULT '' COMMENT 'Address type',
    `address_owner`   bigint NOT NULL  COMMENT 'Address belongs to',
    `private`    varchar(128) NOT NULL DEFAULT '' COMMENT 'Address private key',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Address create time',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Address update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Address delete time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_address` (`address`) COMMENT 'address index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Address information table'