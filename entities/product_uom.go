package entities

//ProductUom defines struct of product uom
type ProductUom struct {
	ProductID int64  `json:"product_id"`
	UomID     int64  `json:"uom_id"`
	UomDesc   string `json:"uom_desc"`
}
