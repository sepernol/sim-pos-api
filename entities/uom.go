package entities

//Uom defines entity for uom object
type Uom struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
