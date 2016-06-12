package repositories

import (
	"database/sql"
	"errors"
	e "github.com/sepernol/sim-pos-api/entities"
	h "github.com/sepernol/sim-pos-api/helpers"
)

func fetchSuppliers(rows *sql.Rows, paging h.PageParams) (result []e.Supplier, err error) {
	list := make([]e.Supplier, paging.Size)
	index := 0
	for rows.Next() {
		if index >= paging.Size {
			break
		}
		var obj e.Supplier
		err = rows.Scan(&obj.ID, &obj.Code, &obj.Name, &obj.IsTaxed)
		if err != nil {
			return
		}
		list[index] = obj
		index++
	}
	result = list[:index]
	return
}

//GetSuppliers gets list of suppliers paged
func GetSuppliers(paging h.PageParams) (result []e.Supplier, err error) {
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

//GetSupplier gets supplier by id
func GetSupplier(id int64) (result e.Supplier, err error) {
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

//InsertSupplier inserts supplier data
func InsertSupplier(data *e.Supplier) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	result, err := h.ExecStatement(db, "INSERT INTO suppliers (code, name, is_taxed) VALUES (?, ?, ?)", data.Code, data.Name, data.IsTaxed)
	if err != nil {
		return
	}
	data.ID = result.LastInsertId
	return
}

//UpdateSupplier updates supplier data
func UpdateSupplier(data *e.Supplier) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "UPDATE suppliers SET code = ?, name = ?, is_taxed = ? where id = ?", data.Code, data.Name, data.IsTaxed, data.ID)
	if err != nil {
		return
	}

	return
}

//DeleteSupplier deletes supplier data
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
