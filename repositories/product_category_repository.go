package repositories

import (
	"database/sql"
	"errors"
	e "github.com/sepernol/sim-pos-api/entities"
	h "github.com/sepernol/sim-pos-api/helpers"
)

func fetchProductCategories(rows *sql.Rows, paging h.PageParams) (result []e.ProductCategory, err error) {
	list := make([]e.ProductCategory, paging.Size)
	index := 0
	for rows.Next() {
		if index >= paging.Size {
			break
		}
		var obj e.ProductCategory
		err = rows.Scan(&obj.ID, &obj.Code, &obj.Description)
		if err != nil {
			return
		}
		list[index] = obj
		index++
	}
	result = list[:index]
	return
}

//GetProductCategories gets list product categories with paging
func GetProductCategories(paging h.PageParams) (result []e.ProductCategory, err error) {
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

	result, err = fetchProductCategories(rows, paging)
	if err != nil {
		return
	}
	return
}

//GetProductCategory gets product category by id
func GetProductCategory(id int64) (result e.ProductCategory, err error) {
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

	list, err := fetchProductCategories(rows, h.PageParams{Size: 1})
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

//InsertProductCategory inserts product category
func InsertProductCategory(data *e.ProductCategory) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	result, err := h.ExecStatement(db, "INSERT INTO product_categories (code, description) VALUES (?, ?)", data.Code, data.Description)
	if err != nil {
		return
	}
	data.ID = result.LastInsertId
	return
}

//UpdateProductCategory update product category
func UpdateProductCategory(data *e.ProductCategory) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "UPDATE product_categories SET code = ?, description = ? where id = ?", data.Code, data.Description, data.ID)
	if err != nil {
		return
	}

	return
}

//DeleteProductCategory deletes product category
func DeleteProductCategory(id int64) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "DELETE FROM product_categories where id = ?", id)
	if err != nil {
		return
	}
	return
}
