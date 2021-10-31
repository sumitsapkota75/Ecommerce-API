CREATE TABLE IF NOT EXISTS `coupons` (
    `id` BINARY(16) NOT NULL,
    `code` VARCHAR(255) NOT NULL,
    `device` VARCHAR(255) NULL,
    `discount` VARCHAR(255) NOT NULL,
    `discount_type` VARCHAR(255) NOT NULL,
    `valid_from` DATETIME NOT NULL,
    `valid_to` DATETIME NOT NULL,
    `min_checkout_value` FLOAT NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;