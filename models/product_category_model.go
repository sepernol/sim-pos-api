package models

import (
	"database/sql"
	"errors"
	h "github.com/sepernol/sim-pos-api/helpers"
)

type ProductCategory struct {
	Id          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func fetchRows(rows *sql.Rows, paging h.PageParams) (result []ProductCategory, err error) {
	list := make([]ProductCategory, paging.Size)
	index := 0
	for rows.Next() {
		if index >= paging.Size {
			break
		}
		var obj ProductCategory
		err = rows.Scan(&obj.Id, &obj.Code, &obj.Description)
		if err != nil {
			return
		}
		list[index] = obj
		index++
	}
	result = list[:index]
	return
}

func GetProductCategories(paging h.PageParams) (result []ProductCategory, err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	query := "select * from product_categories" + h.GenerateLimitStatement(&paging)
	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	result, err = fetchRows(rows, paging)
	if err != nil {
		return
	}
	return
}

func GetProductCategory(id int) (result ProductCategory, err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	query := "select * from product_categories where id = ?"
	rows, err := db.Query(query, id)

	if err != nil {
		return
	}
	defer rows.Close()

	list, err := fetchRows(rows, h.PageParams{Size: 1})
	if err != nil {
		return
	}
	if len(list) < 1 {
		err = errors.New("Product Category not found")
		return
	}
	err = nil
	result = list[0]
	return
}

func InsertProductCategory(data *ProductCategory) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO product_categories (code, description) VALUES (?, ?)", data.Code, data.Description)
	if err != nil {
		return
	}
	data.Id, err = result.LastInsertId()
	if err != nil {
		data.Id = 0
		return
	}
	return

}
