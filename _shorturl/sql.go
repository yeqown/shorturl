package shorturl

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "log"
)

var db *sql.DB

func GetDB() (*sql.DB, error) {
	if db == nil {
		ins := GetInstance()
		if err := ConnectDB(ins.MySql); err != nil {
			return nil, err
		}
	}
	return db, nil
}

func ConnectDB(connStr string) error {
	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func CloseConnection() {
	db.Close()
	db = nil
}
