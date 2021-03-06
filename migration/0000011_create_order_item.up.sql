CREATE TABLE  IF NOT EXISTS `order_items`(
    `id` BINARY(16) NOT NULL,
    `product_id` BINARY(16) NOT NULL,
    `order_id` BINARY(16) NOT NULL,
    `price`  FLOAT NOT NULL,
    `quantity` INT NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME  NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY(`product_id`)  
	    REFERENCES `products`(`id`) ON DELETE NO ACTION,
    FOREIGN KEY(`order_id`)  
	    REFERENCES `orders`(`id`) ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;