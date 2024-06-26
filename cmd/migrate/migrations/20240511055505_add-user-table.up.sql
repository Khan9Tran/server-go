CREATE TABLE IF NOT EXISTS users(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT, 
    `firstName` NVARCHAR(255) NOT NULL,
    `lastName` NVARCHAR(255) NOT NULL,
    `email` NVARCHAR(255) NOT NULL,
    `password` NVARCHAR(255) NOT NULL,
    `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    UNIQUE KEY(`email`)
)