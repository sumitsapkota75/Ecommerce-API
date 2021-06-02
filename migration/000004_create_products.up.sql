CREATE TABLE IF NOT EXISTS `products`(
    `id` INT NOT NUll AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `category_id`INT NULL,
    `description` TEXT  NULL,
    `brand_id` INT NULL,
    `cost_price` INT UNSIGNED  NULL,
    `price` INT UNSIGNED  NULL,
    `slug` VARCHAR(255) NOT NULL,
    `code` VARCHAR(255) NOT NULL,
    `quantity` INT UNSIGNED NOT NULL,
    `specification` TEXT NOT NULL,
    `top_selling` INT NOT NULL DEFAULT 0,
    `new_arrival` INT NOT NULL DEFAULT 0,
    `daily_deal` INT NOT NULL DEFAULT 0,
    `order_limit` INT NOT NULL DEFAULT 0,
    `stock_alert` INT NOT NULL DEFAULT 0,
    `refundable` INT NOT NULL DEFAULT 0,
    `is_active` INT NOT NULL DEFAULT 0,
    `featured_collection` INT NOT NULL DEFAULT 0,
    `thumbnail` TEXT  NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`category_id`)
        REFERENCES `categories` (`id`)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (`brand_id`)
        REFERENCES `brands` (`id`)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;