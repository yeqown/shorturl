package orm

import (
	"github.com/go-sql-driver/mysql"
)

func isDuplicateIdx(err error) bool {
	if v, ok := err.(*mysql.MySQLError); ok && v.Number == 1062 {
		return true
	}

	return false
}
