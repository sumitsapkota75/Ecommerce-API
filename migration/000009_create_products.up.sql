CREATE TABLE IF NOT EXISTS `products`(
    `id` BINARY(16) NOT NUll ,
    `name` VARCHAR(255) NOT NULL,
    `category_id` BINARY(16) NULL,
    `vendor_id` VARCHAR(32) NOT NULL,
    `description` TEXT  NULL,
    `brand_id` BINARY(16) NULL,
    `cost_price` FLOAT  NULL,
    `price` FLOAT  NULL,
    `slug` VARCHAR(255) NOT NULL,
    `code` VARCHAR(255) NOT NULL,
    `quantity` FLOAT NOT NULL,
    `specification` TEXT NOT NULL,
    `top_selling` INT NOT NULL DEFAULT 0,
    `new_arrival` INT NOT NULL DEFAULT 0,
    `daily_deal` INT NOT NULL DEFAULT 0,
    `order_limit` INT NOT NULL DEFAULT 0,
    `stock_alert` INT NOT NULL DEFAULT 0,
    `refundable` INT NOT NULL DEFAULT 0,
    `is_active` INT NOT NULL DEFAULT 0,
    `sale_price` FLOAT NULL,
    `featured_collection` INT NOT NULL DEFAULT 0,
    `thumbnail` TEXT  NULL,
    `sale_from` DATETIME  NULL,
    `sale_to` DATETIME  NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`category_id`)
        REFERENCES `categories` (`id`)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (`brand_id`)
        REFERENCES `brands` (`id`)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (`vendor_id`)
        REFERENCES `vendors` (`id`)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;