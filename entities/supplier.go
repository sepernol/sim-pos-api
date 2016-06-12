package entities

//Supplier defines supplier structure
type Supplier struct {
	ID      int64  `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	IsTaxed bool   `json:"is_taxed"`
}
