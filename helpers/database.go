package helpers

import (
	"database/sql"
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
