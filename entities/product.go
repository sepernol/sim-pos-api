package entities

//Product defines structure for product entity
type Product struct {
	ID         int64  `json:"id"`
	SKU        string `json:"sku"`
	Name       string `json:"name"`
	ShortName  string `json:"short_name"`
	CategoryID int64  `json:"category_id"`
}
