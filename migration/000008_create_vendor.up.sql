CREATE TABLE IF NOT EXISTS `vendors` (
    `id` VARCHAR(32) NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `email` VARCHAR(50) NOT NULL,
    `address` VARCHAR(200) NOT NULL,
    `store_name` VARCHAR(200) NOT NULL,
    `document_type` VARCHAR(200) NOT NULL,
    `document_id` VARCHAR(200) NOT NULL,
    `phone` VARCHAR(10) NOT NULL,
    `thumbnail` VARCHAR(300) NULL,
    `is_active` INT NOT NUll,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4; 