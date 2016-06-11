package repositories

import (
	"database/sql"
	"errors"
	e "github.com/sepernol/sim-pos-api/entities"
	h "github.com/sepernol/sim-pos-api/helpers"
)

func fetchUoms(rows *sql.Rows, paging h.PageParams) (result []e.Uom, err error) {
	list := make([]e.Uom, paging.Size)
	index := 0
	for rows.Next() {
		if index >= paging.Size {
			break
		}
		var obj e.Uom
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

//GetUoms gets list of Uom with paging
func GetUoms(paging h.PageParams) (result []e.Uom, err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	query := "select * from uoms" + h.GenerateLimitStatement(&paging)
	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	result, err = fetchUoms(rows, paging)
	if err != nil {
		return
	}
	return
}

//GetUom gets uom by ID
func GetUom(id int64) (result e.Uom, err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	query := "select * from uoms where id = ?"
	rows, err := db.Query(query, id)

	if err != nil {
		return
	}
	defer rows.Close()

	list, err := fetchUoms(rows, h.PageParams{Size: 1})
	if err != nil {
		return
	}
	if len(list) < 1 {
		err = errors.New("Uom not found")
		return
	}
	err = nil
	result = list[0]
	return
}

//InsertUom inserts new record in database
func InsertUom(data *e.Uom) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	result, err := h.ExecStatement(db, "INSERT INTO uoms (code, description) VALUES (?, ?)", data.Code, data.Description)
	if err != nil {
		return
	}
	data.ID = result.LastInsertId
	return
}

//UpdateUom updates record in database
func UpdateUom(data *e.Uom) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "UPDATE uoms SET code = ?, description = ? where id = ?", data.Code, data.Description, data.ID)
	if err != nil {
		return
	}

	return
}

//DeleteUom deletes record from database
func DeleteUom(id int64) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "DELETE FROM uoms where id = ?", id)
	if err != nil {
		return
	}
	return
}
