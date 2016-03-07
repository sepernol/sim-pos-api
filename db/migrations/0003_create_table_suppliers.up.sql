CREATE TABLE suppliers (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `code` varchar(10) NOT NULL,
    `name` varchar(100) NOT NULL,
    `is_taxed` boolean NOT NULL,
    PRIMARY KEY (`id`)
)
