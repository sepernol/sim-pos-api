package models

import (
	"database/sql"
	"errors"
	h "github.com/sepernol/sim-pos-api/helpers"
)

type Product struct {
	Id         int64  `json:"id"`
	SKU        string `json:"sku"`
	Name       string `json:"name"`
	ShortName  string `json:"short_name"`
	CategoryId int64  `json:"category_id"`
}

func fetchProducts(rows *sql.Rows, paging h.PageParams) (result []Product, err error) {
	list := make([]Product, paging.Size)
	index := 0
	for rows.Next() {
		if index >= paging.Size {
			break
		}
		var obj Product
		err = rows.Scan(&obj.Id, &obj.SKU, &obj.Name, &obj.ShortName, &obj.CategoryId)
		if err != nil {
			return
		}
		list[index] = obj
		index++
	}
	result = list[:index]
	return
}

func GetProducts(paging h.PageParams) (result []Product, err error) {
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

func GetProduct(id int64) (result Product, err error) {
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

func InsertProduct(data *Product) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	result, err := h.ExecStatement(db, "INSERT INTO products (sku, name, short_name, category_id) VALUES (?, ?, ?, ?)", data.SKU, data.Name, data.ShortName, data.CategoryId)
	if err != nil {
		return
	}
	data.Id = result.LastInsertId
	return
}

func UpdateProduct(data *Product) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "UPDATE products SET sku = ?, name = ?, short_name = ?, category_id = ? where id = ?", data.SKU, data.Name, data.ShortName, data.CategoryId, data.Id)
	if err != nil {
		return
	}

	return
}

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
