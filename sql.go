package shorturl

import (
	"database/sql"
	"time"
)

var db *sql.DB

// GetDB ...
func GetDB() (*sql.DB, error) {
	if db == nil {
		// ins := GetInstance()
		// if err := ConnectDB(ins.MySql); err != nil {
		// 	return nil, err
		// }
		panic("DB is disconnected !")
	}
	return db, nil
}

// ConnectDB ...
// yeqown:yeqown@/shorturl
func ConnectDB(connStr string) (err error) {
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	db.SetConnMaxLifetime(10 * time.Second)
	return nil
}

// CloseConnection ...
func CloseConnection() {
	db.Close()
	db = nil
}
