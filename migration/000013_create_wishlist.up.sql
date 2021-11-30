CREATE TABLE IF NOT EXISTS `wishlists`(
    `id` BINARY(16) NOT NULL,
    `product_id` BINARY(16) NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`product_id`) 
        REFERENCES `products` (`id`)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (`user_id`) 
        REFERENCES `users` (`id`)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;