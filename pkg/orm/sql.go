package orm

import (
	"database/sql"
	"errors"
	"sync"
	"time"
)

var (
	_once sync.Once
	_orm  *ShortORM
)

type ShortORM struct {
	*sql.DB

	dsn    string
	driver string

	_mu       sync.RWMutex
	stmtCache map[string]*sql.Stmt
}

func NewORM(dsn, driver string) (*ShortORM, error) {
	var err error

	if _orm == nil {
		_once.Do(func() {
			_orm = &ShortORM{
				dsn:    dsn,
				driver: driver,
			}
			err = _orm.initConn()
		})
	}

	return _orm, err
}

func (o *ShortORM) initConn() (err error) {
	o.DB, err = sql.Open(o.driver, o.dsn)
	if err != nil {
		return
	}

	if err = o.DB.Ping(); err != nil {
		return
	}

	// TODO: make these as options
	o.DB.SetConnMaxLifetime(10 * time.Second)
	o.DB.SetMaxOpenConns(100)
	o.DB.SetMaxIdleConns(20)

	return nil
}

// Close ...
func (o *ShortORM) Close() error {
	if o.DB != nil {
		v := o.DB
		o.DB = nil
		return v.Close()
	}

	return nil
}

func (o *ShortORM) Create(m *ShortURLDO) (err error) {
	_key := o.stmtCacheKey(m, "insert")
	stmt, ok := o.stmtCache[_key]
	if !ok {
		if stmt, err = o.Prepare(_insertSQL); err != nil {
			return
		}

		o.stmtCache[_key] = stmt
	}

	ret, err := stmt.Exec(m.Source)
	if err != nil {
		return err
	}

	id, err := ret.LastInsertId()
	if err != nil {
		return err
	}

	m.ID = id
	return nil
}

func (o *ShortORM) Update(m *ShortURLDO) (err error) {
	if m.ID <= 0 {
		return errors.New("invalid ID")
	}

	_key := o.stmtCacheKey(m, "update")
	stmt, ok := o.stmtCache[_key]
	if !ok {
		if stmt, err = o.Prepare(_updateSQL); err != nil {
			return
		}

		o.stmtCache[_key] = stmt
	}

	_, err = stmt.Exec(m.Source, m.Shorted, m.ID)
	if err != nil {
		return err
	}
	return nil
}

func (o *ShortORM) Query(m *ShortURLDO) (err error) {
	if m.ID <= 0 {
		return errors.New("invalid ID")
	}

	_key := o.stmtCacheKey(m, "query")
	stmt, ok := o.stmtCache[_key]
	if !ok {
		if stmt, err = o.Prepare(_querySQL); err != nil {
			return
		}

		o.stmtCache[_key] = stmt
	}

	if row := stmt.QueryRow(m.ID); true {
		return row.Scan(&m.ID, &m.Source, &m.Shorted)
	}

	return err
}

func (o *ShortORM) stmtCacheKey(t ITabler, operation string) string {
	return t.TableName() + "." + operation
}
