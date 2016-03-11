CREATE TABLE products (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `sku` varchar(30) NOT NULL,
    `name` varchar(100) NOT NULL,
    `short_name` varchar(20) NOT NULL,
    `category_id` int(10) NOT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`category_id`) REFERENCES product_categories(`id`)
)
