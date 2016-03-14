package models

import (
	"database/sql"
	"errors"
	h "github.com/sepernol/sim-pos-api/helpers"
)

type ProductUnitPrice struct {
	ProductId int64   `json:"product_id"`
	UomId     int64   `json:"uom_id"`
	UomDesc   string  `json:"uom_desc"`
	UnitPrice float32 `json:"unit_price"`
}

func InsertProductUnitPrice(data *ProductUnitPrice) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "INSERT INTO product_unit_prices (product_id, uom_id, unit_price) VALUES (?, ?, ?)", data.ProductId, data.UomId, data.UnitPrice)
	if err != nil {
		return
	}
	return
}

func UpdateProductUnitPrice(data *ProductUnitPrice) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "UPDATE product_unit_prices SET unit_price = ? WHERE product_id = ? AND uom_id = ?", data.UnitPrice, data.ProductId, data.UomId)
	if err != nil {
		return
	}
	return
}

func GetProductUnitPrices(productId int64) (result []ProductUnitPrice, err error) {
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

	rows, err := db.Query(query, productId)
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

func fetchProductUnitPrices(rows *sql.Rows) (result []ProductUnitPrice, err error) {
	result = make([]ProductUnitPrice, 0)
	for rows.Next() {
		var obj ProductUnitPrice
		err = rows.Scan(&obj.ProductId, &obj.UomId, &obj.UomDesc, &obj.UnitPrice)
		if err != nil {
			return
		}
		result = append(result, obj)
	}
	return
}

func GetProductUnitPrice(productId int64, uomId int64) (result ProductUnitPrice, err error) {
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

	rows, err := db.Query(query, productId, uomId)
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

func DeleteProductUnitPrice(productId int64, uomId int64) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "DELETE FROM product_unit_prices where product_id = ? AND uom_id = ? ", productId, uomId)
	if err != nil {
		return
	}
	return
}
