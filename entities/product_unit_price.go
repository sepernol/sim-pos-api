package entities

//ProductUnitPrice defines structure of product unit price
type ProductUnitPrice struct {
	ProductID int64   `json:"product_id"`
	UomID     int64   `json:"uom_id"`
	UomDesc   string  `json:"uom_desc"`
	UnitPrice float32 `json:"unit_price"`
}
