package models

import (
	"database/sql"
	"errors"
	h "github.com/sepernol/sim-pos-api/helpers"
)

type Supplier struct {
	Id      int64  `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	IsTaxed bool   `json:"is_taxed"`
}

func fetchSuppliers(rows *sql.Rows, paging h.PageParams) (result []Supplier, err error) {
	list := make([]Supplier, paging.Size)
	index := 0
	for rows.Next() {
		if index >= paging.Size {
			break
		}
		var obj Supplier
		err = rows.Scan(&obj.Id, &obj.Code, &obj.Name, &obj.IsTaxed)
		if err != nil {
			return
		}
		list[index] = obj
		index++
	}
	result = list[:index]
	return
}

func GetSuppliers(paging h.PageParams) (result []Supplier, err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	query := "SELECT * FROM suppliers" + h.GenerateLimitStatement(&paging)
	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	result, err = fetchSuppliers(rows, paging)
	if err != nil {
		return
	}
	return
}

func GetSupplier(id int64) (result Supplier, err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	query := "SELECT * FROM suppliers where id = ?"
	rows, err := db.Query(query, id)

	if err != nil {
		return
	}
	defer rows.Close()

	list, err := fetchSuppliers(rows, h.PageParams{Size: 1})
	if err != nil {
		return
	}
	if len(list) < 1 {
		err = errors.New("Supplier not found")
		return
	}
	err = nil
	result = list[0]
	return
}

func InsertSupplier(data *Supplier) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	result, err := h.ExecStatement(db, "INSERT INTO suppliers (code, name, is_taxed) VALUES (?, ?, ?)", data.Code, data.Name, data.IsTaxed)
	if err != nil {
		return
	}
	data.Id = result.LastInsertId
	return
}

func UpdateSupplier(data *Supplier) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "UPDATE suppliers SET code = ?, name = ?, is_taxed = ? where id = ?", data.Code, data.Name, data.IsTaxed, data.Id)
	if err != nil {
		return
	}

	return
}

func DeleteSupplier(id int64) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "DELETE FROM suppliers where id = ?", id)
	if err != nil {
		return
	}
	return
}
