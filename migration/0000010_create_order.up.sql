CREATE TABLE  IF NOT EXISTS `orders`(
    `id` BINARY(16) NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `street_address` VARCHAR(255) NOT NULL,
    `total_amount` FLOAT NULL,
    `paid_amount` FLOAT NULL,
    `city` VARCHAR(255)  NULL,
    `state` VARCHAR(255)  NULL,
    `zip` VARCHAR(255)  NULL,
    `phone` VARCHAR(255)  NULL,
    `email` VARCHAR(255)  NULL,
    `notes` VARCHAR(255)  NULL,
    `order_status` VARCHAR(255) DEFAULT "ORDER PLACED",
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME  NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;