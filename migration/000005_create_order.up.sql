CREATE TABLE  IF NOT EXISTS `orders`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `company_name` VARCHAR(255) NULL,
    `street_address` VARCHAR(255) NOT NULL,
    `total_amount` VARCHAR(255) NULL,
    `paid_amount` VARCHAR(255) NULL,
    `city` VARCHAR(255)  NULL,
    `state` VARCHAR(255)  NULL,
    `zip` VARCHAR(255)  NULL,
    `phone` VARCHAR(255)  NULL,
    `email` VARCHAR(255)  NULL,
    `notes` VARCHAR(255)  NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME  NULL,
    PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;