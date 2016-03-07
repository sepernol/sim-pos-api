package helpers

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func GetDBConnection() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:1234@tcp/sim_pos")
	return
}

func GenerateLimitStatement(p *PageParams) string {
	return " limit " + strconv.Itoa(p.Offset) + "," + strconv.Itoa(p.Size)
}

type QueryResult struct {
	RowsAffected int64
	LastInsertId int64
}

func ExecStatement(db *sql.DB, statement string, args ...interface{}) (queryResult QueryResult, err error) {
	result, err := db.Exec(statement, args...)
	if err != nil {
		return
	}

	queryResult.RowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}
	queryResult.LastInsertId, err = result.LastInsertId()
	if err != nil {
		return
	}

	if queryResult.RowsAffected == 0 {
		err = errors.New("No data affected")
		return
	}
	return

}
