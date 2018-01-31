package shorturl

import (
	"database/sql"
)

type UrlModel struct {
	DB       *sql.DB
	Id       int64
	LongUrl  string
	ShortUrl string
}

func (u *UrlModel) Clear() {
	u.LongUrl = ""
	u.ShortUrl = ""
}

func (u *UrlModel) Insert() (int64, error) {
	stmt, err := u.DB.Prepare("INSERT longurl SET id=?, long_url=?")

	if err != nil {
		return -1, err
	}
	var ret sql.Result
	if ret, err = stmt.Exec(nil, u.LongUrl); err != nil {
		return -1, err
	}
	u.Id, _ = ret.LastInsertId()
	return ret.LastInsertId()
}

func (u *UrlModel) Query() error {
	stmt, err := u.DB.Prepare("SELECT id, long_url, short_url FROM longurl WHERE id=?")
	if err != nil {
		return err
	}
	if row := stmt.QueryRow(u.Id); err == nil {
		row.Scan(&u.Id, &u.LongUrl, &u.ShortUrl)
		return nil
	} else {
		return err
	}
}

func (u *UrlModel) Update() error {
	stmt, err := u.DB.Prepare("update longurl set long_url=?, short_url=? where id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.LongUrl, u.ShortUrl, u.Id)
	if err != nil {
		return err
	}
	return nil
}
