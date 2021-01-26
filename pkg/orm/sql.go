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
	// TODO(@yeqown) what's sql.Stmt, concurrent safe?
}

func NewORM(dsn, driver string) (*ShortORM, error) {
	var err error

	if _orm == nil {
		_once.Do(func() {
			_orm = &ShortORM{
				dsn:       dsn,
				driver:    driver,
				_mu:       sync.RWMutex{},
				stmtCache: make(map[string]*sql.Stmt),
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
	m.hash()

	_key := o.stmtCacheKey(m, "insert")
	o._mu.RLock()
	stmt, ok := o.stmtCache[_key]
	o._mu.RUnlock()
	if !ok {
		if stmt, err = o.Prepare(_insertSQL); err != nil {
			return
		}

		o._mu.Lock()
		o.stmtCache[_key] = stmt
		o._mu.Unlock()
	}

	ret, err := stmt.Exec(m.Source, m.Hash)
	if err != nil {
		if isDuplicateIdx(err) {
			return o.QueryByHash(m)
		}

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
	o._mu.RLock()
	stmt, ok := o.stmtCache[_key]
	o._mu.RUnlock()
	if !ok {
		if stmt, err = o.Prepare(_updateSQL); err != nil {
			return
		}

		o._mu.Lock()
		o.stmtCache[_key] = stmt
		o._mu.Unlock()
	}

	_, err = stmt.Exec(m.Shorted, m.ID)
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
	o._mu.RLock()
	stmt, ok := o.stmtCache[_key]
	o._mu.RUnlock()
	if !ok {
		if stmt, err = o.Prepare(_querySQL); err != nil {
			return
		}

		o._mu.Lock()
		o.stmtCache[_key] = stmt
		o._mu.Unlock()
	}

	if row := stmt.QueryRow(m.ID); true {
		return row.Scan(&m.ID, &m.Source, &m.Hash, &m.Shorted)
	}

	return err
}

func (o *ShortORM) QueryByHash(m *ShortURLDO) (err error) {
	if m.Hash <= 0 {
		return errors.New("invalid hash")
	}

	_key := o.stmtCacheKey(m, "query_by_hash")
	o._mu.RLock()
	stmt, ok := o.stmtCache[_key]
	o._mu.RUnlock()
	if !ok {
		if stmt, err = o.Prepare(_queryByHashSQL); err != nil {
			return
		}

		o._mu.Lock()
		o.stmtCache[_key] = stmt
		o._mu.Unlock()
	}

	if row := stmt.QueryRow(m.Hash); true {
		return row.Scan(&m.ID, &m.Source, &m.Hash, &m.Shorted)
	}

	return err
}

func (o *ShortORM) stmtCacheKey(t ITabler, operation string) string {
	return t.TableName() + "." + operation
}
