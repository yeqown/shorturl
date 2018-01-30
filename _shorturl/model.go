package shorturl

import (
	"database/sql"
)

type UrlModel struct {
	DB       *sql.DB
	LongUrl  string
	Id       int64
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
	stmt, err := u.DB.Prepare("Select * from longurl where id=? limit 1")
	if err != nil {
		return err
	}
	rows, _ := stmt.Query(u.Id)
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&u.Id, &u.LongUrl, &u.ShortUrl); err == nil {
			break
		}
	}
	return nil
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
