package repositories

import (
	"database/sql"
	"errors"
	e "github.com/sepernol/sim-pos-api/entities"
	h "github.com/sepernol/sim-pos-api/helpers"
)

func fetchProducts(rows *sql.Rows, paging h.PageParams) (result []e.Product, err error) {
	list := make([]e.Product, paging.Size)
	index := 0
	for rows.Next() {
		if index >= paging.Size {
			break
		}
		var obj e.Product
		err = rows.Scan(&obj.ID, &obj.SKU, &obj.Name, &obj.ShortName, &obj.CategoryID)
		if err != nil {
			return
		}
		list[index] = obj
		index++
	}
	result = list[:index]
	return
}

//GetProducts gets product list with paging
func GetProducts(paging h.PageParams) (result []e.Product, err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	query := "select * from products " + h.GenerateLimitStatement(&paging)
	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	result, err = fetchProducts(rows, paging)
	if err != nil {
		return
	}
	return
}

//GetProduct gets product by id
func GetProduct(id int64) (result e.Product, err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	query := "select * from products where id = ?"
	rows, err := db.Query(query, id)

	if err != nil {
		return
	}
	defer rows.Close()

	list, err := fetchProducts(rows, h.PageParams{Size: 1})
	if err != nil {
		return
	}
	if len(list) < 1 {
		err = errors.New("Product not found")
		return
	}
	err = nil
	result = list[0]
	return
}

//InsertProduct inserts new record in products
func InsertProduct(data *e.Product) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	result, err := h.ExecStatement(db, "INSERT INTO products (sku, name, short_name, category_id) VALUES (?, ?, ?, ?)", data.SKU, data.Name, data.ShortName, data.CategoryID)
	if err != nil {
		return
	}
	data.ID = result.LastInsertId
	return
}

//UpdateProduct updates existing data in products
func UpdateProduct(data *e.Product) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "UPDATE products SET sku = ?, name = ?, short_name = ?, category_id = ? where id = ?", data.SKU, data.Name, data.ShortName, data.CategoryID, data.ID)
	if err != nil {
		return
	}

	return
}

//DeleteProduct deletes product record from database
func DeleteProduct(id int64) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "DELETE FROM products where id = ?", id)
	if err != nil {
		return
	}
	return
}
