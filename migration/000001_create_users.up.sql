CREATE TABLE IF NOT EXISTS `users` (
    `id` VARCHAR(32) NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `email` VARCHAR(50) NOT NULL,
    `address` VARCHAR(200) NOT NULL,
    `phone` VARCHAR(10) NOT NULL,
    `user_type` VARCHAR(30) NOT NULL DEFAULT 'customer', /* Customer, admin */
    `is_verified` TINYINT UNSIGNED NOT NULL DEFAULT 0,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4; 