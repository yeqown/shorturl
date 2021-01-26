package orm

import "database/sql"

const (
	_insertSQL = "INSERT INTO shorted_url(source) VALUES(?)"
	_updateSQL = "UPDATE shorted_url SET shorted = ? WHERE id = ?"
	_querySQL  = "SELECT id, source, shorted FROM shorted_url WHERE id = ?"
)

// ShortURLDO ...
type ShortURLDO struct {
	*sql.DB

	ID      int64
	Source  string
	Shorted string
}

func (m *ShortURLDO) TableName() string {
	return "shorted_url"
}
