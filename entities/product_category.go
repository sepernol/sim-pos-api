package entities

//ProductCategory defines struct for product category
type ProductCategory struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
