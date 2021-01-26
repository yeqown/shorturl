package orm

import (
	"database/sql"
	"hash/crc32"
)

const (
	_insertSQL      = "INSERT INTO shorted_url(source, hash) VALUES(?, ?)"
	_updateSQL      = "UPDATE shorted_url SET shorted = ? WHERE id = ?"
	_querySQL       = "SELECT id, source, hash, shorted FROM shorted_url WHERE id = ?"
	_queryByHashSQL = "SELECT id, source, hash, shorted FROM shorted_url WHERE hash = ?"
)

// ShortURLDO ...
type ShortURLDO struct {
	*sql.DB

	ID      int64
	Hash    int64 // md5 hash of source URL
	Source  string
	Shorted string
}

func (m *ShortURLDO) TableName() string {
	return "shorted_url"
}

func (m *ShortURLDO) hash() {
	v := int64(crc32.ChecksumIEEE([]byte(m.Source)))
	if v >= 0 {
		m.Hash = v
		return
	}

	m.Hash = 0 - v
	return
}
