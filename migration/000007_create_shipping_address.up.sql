CREATE TABLE IF NOT EXISTS `shipping_addresses`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` VARCHAR(32) NOT NULL,
    `district`  VARCHAR(255)  NOT NULL,
    `address` TEXT NULL,
    `phone` VARCHAR(15) NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME  NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY(`user_id`)  
	    REFERENCES `users`(`id`) ON DELETE NO ACTION
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;