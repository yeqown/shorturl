package shorturl

import (
	"database/sql"
	"sync"
	"time"
)

var (
	db *sql.DB
	once sync.Once
)

// GetDB ...
func GetDB() (*sql.DB, error) {
	if db == nil {
		panic("DB is nil !")
	}

	return db, nil
}

// ConnectDB ...
// addr user:pwd@/dbname
func ConnectDB(addr string) (err error) {
	once.Do(func() {
		db, err = sql.Open("mysql", addr)
		if err != nil {
			return
		}

		if err = db.Ping(); err != nil {
			return
		}
		db.SetConnMaxLifetime(10 * time.Second)
	})

	if err != nil {
		// reinitialize once
		once = sync.Once{}
	}

	return
}

// CloseConnection ...
func CloseConnection() {
	db.Close()
	db = nil
}
