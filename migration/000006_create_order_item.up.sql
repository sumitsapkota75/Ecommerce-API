CREATE TABLE  IF NOT EXISTS `order_items`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `product_id` INT NOT NULL,
    `order_id` INT NOT NULL,
    `price`  FLOAT NOT NULL,
    `quantity` INT NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME  NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT fk_product
    FOREIGN KEY(product_id)  
	    REFERENCES products(id) ON DELETE NO ACTION,
    CONSTRAINT fk_order
    FOREIGN KEY(order_id)  
	    REFERENCES orders(id) ON DELETE NO ACTION
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;