package models

import (
	"database/sql"
	"errors"
	h "github.com/sepernol/sim-pos-api/helpers"
)

type Uom struct {
	Id          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func fetchUoms(rows *sql.Rows, paging h.PageParams) (result []Uom, err error) {
	list := make([]Uom, paging.Size)
	index := 0
	for rows.Next() {
		if index >= paging.Size {
			break
		}
		var obj Uom
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

func GetUoms(paging h.PageParams) (result []Uom, err error) {
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

func GetUom(id int64) (result Uom, err error) {
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

func InsertUom(data *Uom) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	result, err := h.ExecStatement(db, "INSERT INTO uoms (code, description) VALUES (?, ?)", data.Code, data.Description)
	if err != nil {
		return
	}
	data.Id = result.LastInsertId
	return
}

func UpdateUom(data *Uom) (err error) {
	db, err := h.GetDBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = h.ExecStatement(db, "UPDATE uoms SET code = ?, description = ? where id = ?", data.Code, data.Description, data.Id)
	if err != nil {
		return
	}

	return
}

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
