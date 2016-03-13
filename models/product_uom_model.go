package models

import (
	"database/sql"
	"errors"
	h "github.com/sepernol/sim-pos-api/helpers"
)

type ProductUom struct {
	ProductId int64  `json:"product_id"`
	UomId     int64  `json:"uom_id"`
	UomDesc   string `json:"uom_desc"`
}

func DeleteProductUom(productId int64, uomId int64) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "DELETE FROM product_uoms where product_id = ? AND uom_id = ? ", productId, uomId)
	if err != nil {
		return
	}
	return
}

func InsertProductUom(data *ProductUom) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "INSERT INTO product_uoms (product_id, uom_id) VALUES (?, ?)", data.ProductId, data.UomId)
	if err != nil {
		return
	}
	return
}

func fetchProductUoms(rows *sql.Rows) (result []ProductUom, err error) {
	result = make([]ProductUom, 0)
	for rows.Next() {
		var obj ProductUom
		err = rows.Scan(&obj.ProductId, &obj.UomId, &obj.UomDesc)
		if err != nil {
			return
		}
		result = append(result, obj)
	}
	return
}

func GetProductUom(productId int64, uomId int64) (result ProductUom, err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	query := "SELECT pu.product_id, pu.uom_id, u.description AS uom_desc " +
		"FROM product_uoms pu " +
		"LEFT join uoms u " +
		"ON u.id = pu.uom_id " +
		"WHERE product_id = ? " +
		"AND uom_id = ? "

	rows, err := db.Query(query, productId, uomId)
	if err != nil {
		return
	}
	defer rows.Close()

	list, err := fetchProductUoms(rows)
	if err != nil {
		return
	}
	if len(list) < 1 {
		err = errors.New("Product UOM not found")
		return
	}
	err = nil
	result = list[0]
	return
}

func GetProductUoms(productId int64) (result []ProductUom, err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	query := "select pu.product_id, pu.uom_id, u.description as uom_desc " +
		"from product_uoms pu " +
		"left join uoms u " +
		"on u.id = pu.uom_id " +
		"where product_id = ? "

	rows, err := db.Query(query, productId)
	if err != nil {
		return
	}
	defer rows.Close()

	result, err = fetchProductUoms(rows)
	if err != nil {
		return
	}
	return
}
