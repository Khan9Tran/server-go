CREATE TABLE IF NOT EXISTS products(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` NVARCHAR(255) NOT NULL,
    `price` DECIMAL(10, 2) NOT NULL,
    `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `quantity` INT UNSIGNED NOT NULL,
    `description` TEXT NOT NULL,
    `image` NVARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
)