package shorturl

import (
	"database/sql"
)

// URLModel ...
type URLModel struct {
	// DB       *sql.DB
	DB *sql.DB

	// ID       int64
	ID int64

	// LongURL  string
	LongURL string

	// ShortURL string
	ShortURL string
}

// clear ...
func (u *URLModel) clear() {
	u.LongURL = ""
	u.ShortURL = ""
}

// insert ...
func (u *URLModel) insert() (int64, error) {
	stmt, err := u.DB.Prepare("INSERT longurl SET id=?, long_url=?")

	if err != nil {
		return -1, err
	}
	var ret sql.Result
	if ret, err = stmt.Exec(nil, u.LongURL); err != nil {
		return -1, err
	}
	u.ID, _ = ret.LastInsertId()
	return ret.LastInsertId()
}

func (u *URLModel) query() error {
	stmt, err := u.DB.Prepare("SELECT id, long_url, short_url FROM longurl WHERE id=?")
	if err != nil {
		return err
	}

	if row := stmt.QueryRow(u.ID); true {
		row.Scan(&u.ID, &u.LongURL, &u.ShortURL)
		return nil
	}

	return err
}

func (u *URLModel) update() error {
	stmt, err := u.DB.Prepare("update longurl set long_url=?, short_url=? where id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.LongURL, u.ShortURL, u.ID)
	if err != nil {
		return err
	}
	return nil
}
