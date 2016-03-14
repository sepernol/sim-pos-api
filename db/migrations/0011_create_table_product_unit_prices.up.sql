CREATE TABLE product_unit_prices (
    `product_id` int(10) NOT NULL,
    `uom_id` int(10) NOT NULL,
    `unit_price` float(10,2) NOT NULL,

    PRIMARY KEY(`product_id`, `uom_id`),
    FOREIGN KEY(`product_id`) REFERENCES products(`id`),
    FOREIGN KEY(`uom_id`) REFERENCES uoms(`id`)
)
