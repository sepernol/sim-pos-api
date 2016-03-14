package helpers

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
)

const (
	DEFAULT_DB_USER    = "root"
	DEFAULT_DB_PASS    = "1234"
	DEFAULT_DB_NAME    = "sim_pos"
	DEFAULT_DB_ADDRESS = "localhost"
	DEFAULT_DB_PORT    = "3306"
)

func GetDBConnection() (db *sql.DB, err error) {
	dbUser := os.Getenv("SIM_POS_DB_USER")
	if dbUser == "" {
		dbUser = DEFAULT_DB_USER
	}
	dbPass := os.Getenv("SIM_POS_DB_PASS")
	if dbPass == "" {
		dbPass = DEFAULT_DB_PASS
	}
	dbName := os.Getenv("SIM_POS_DB_NAME")
	if dbName == "" {
		dbName = DEFAULT_DB_NAME
	}
	dbAddress := os.Getenv("SIM_POS_DB_ADDRESS")
	if dbAddress == "" {
		dbAddress = DEFAULT_DB_ADDRESS
	}
	dbPort := os.Getenv("SIM_POS_DB_PORT")
	if dbPort == "" {
		dbPort = DEFAULT_DB_PORT
	}
	connStr := dbUser + ":" + dbPass + "@tcp(" + dbAddress + ":" + dbPort + ")/" + dbName
	db, err = sql.Open("mysql", connStr)
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
