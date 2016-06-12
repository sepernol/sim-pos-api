package repositories

import (
	"database/sql"
	"errors"
	e "github.com/sepernol/sim-pos-api/entities"
	h "github.com/sepernol/sim-pos-api/helpers"
)

//InsertProductUnitPrice inserts to database
func InsertProductUnitPrice(data *e.ProductUnitPrice) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "INSERT INTO product_unit_prices (product_id, uom_id, unit_price) VALUES (?, ?, ?)", data.ProductID, data.UomID, data.UnitPrice)
	if err != nil {
		return
	}
	return
}

//UpdateProductUnitPrice updates record in database
func UpdateProductUnitPrice(data *e.ProductUnitPrice) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "UPDATE product_unit_prices SET unit_price = ? WHERE product_id = ? AND uom_id = ?", data.UnitPrice, data.ProductID, data.UomID)
	if err != nil {
		return
	}
	return
}

//GetProductUnitPrices gets list of unit price per product
func GetProductUnitPrices(productID int64) (result []e.ProductUnitPrice, err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	query := "select pu.product_id, pu.uom_id, u.description as uom_desc, pu.unit_price " +
		"from product_unit_prices pu " +
		"left join uoms u " +
		"on u.id = pu.uom_id " +
		"where product_id = ? "

	rows, err := db.Query(query, productID)
	if err != nil {
		return
	}
	defer rows.Close()

	result, err = fetchProductUnitPrices(rows)
	if err != nil {
		return
	}
	return
}

func fetchProductUnitPrices(rows *sql.Rows) (result []e.ProductUnitPrice, err error) {
	result = make([]e.ProductUnitPrice, 0)
	for rows.Next() {
		var obj e.ProductUnitPrice
		err = rows.Scan(&obj.ProductID, &obj.UomID, &obj.UomDesc, &obj.UnitPrice)
		if err != nil {
			return
		}
		result = append(result, obj)
	}
	return
}

//GetProductUnitPrice gets product unit price
func GetProductUnitPrice(productID int64, uomID int64) (result e.ProductUnitPrice, err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	query := "SELECT pu.product_id, pu.uom_id, u.description AS uom_desc, pu.unit_price " +
		"FROM product_unit_prices pu " +
		"LEFT join uoms u " +
		"ON u.id = pu.uom_id " +
		"WHERE product_id = ? " +
		"AND uom_id = ? "

	rows, err := db.Query(query, productID, uomID)
	if err != nil {
		return
	}
	defer rows.Close()

	list, err := fetchProductUnitPrices(rows)
	if err != nil {
		return
	}
	if len(list) < 1 {
		err = errors.New("Product Unit Price not found")
		return
	}
	err = nil
	result = list[0]
	return
}

//DeleteProductUnitPrice deletes from database
func DeleteProductUnitPrice(productID int64, uomID int64) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "DELETE FROM product_unit_prices where product_id = ? AND uom_id = ? ", productID, uomID)
	if err != nil {
		return
	}
	return
}
