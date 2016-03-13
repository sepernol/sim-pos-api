CREATE TABLE product_uoms (
    `product_id` int(10) NOT NULL,
    `uom_id` int(10) NOT NULL,

    PRIMARY KEY(`product_id`, `uom_id`),
    FOREIGN KEY(`product_id`) REFERENCES products(`id`),
    FOREIGN KEY(`uom_id`) REFERENCES uoms(`id`)
)
