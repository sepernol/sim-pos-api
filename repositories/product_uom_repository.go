package repositories

import (
	"database/sql"
	"errors"
	e "github.com/sepernol/sim-pos-api/entities"
	h "github.com/sepernol/sim-pos-api/helpers"
)

//DeleteProductUom deletes from database
func DeleteProductUom(productID int64, uomID int64) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "DELETE FROM product_uoms where product_id = ? AND uom_id = ? ", productID, uomID)
	if err != nil {
		return
	}
	return
}

//InsertProductUom inserts to database
func InsertProductUom(data *e.ProductUom) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "INSERT INTO product_uoms (product_id, uom_id) VALUES (?, ?)", data.ProductID, data.UomID)
	if err != nil {
		return
	}
	return
}

func fetchProductUoms(rows *sql.Rows) (result []e.ProductUom, err error) {
	result = make([]e.ProductUom, 0)
	for rows.Next() {
		var obj e.ProductUom
		err = rows.Scan(&obj.ProductID, &obj.UomID, &obj.UomDesc)
		if err != nil {
			return
		}
		result = append(result, obj)
	}
	return
}

//GetProductUom gets product uom by uom id and product id
func GetProductUom(productID int64, uomID int64) (result e.ProductUom, err error) {
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

	rows, err := db.Query(query, productID, uomID)
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

//GetProductUoms gets all product uoms per product
func GetProductUoms(productID int64) (result []e.ProductUom, err error) {
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

	rows, err := db.Query(query, productID)
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
