package orm

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
)

type ormTestSuite struct {
	suite.Suite

	orm *ShortORM
}

func (o *ormTestSuite) SetupSuite() {
	var (
		err error
		dsn = "test:test@tcp(127.0.0.1:3306)/shorten_url?"
	)
	o.orm, err = NewORM(dsn, "mysql")
	o.Require().Nil(err)
}

func (o ormTestSuite) Test_Create() {
	m := ShortURLDO{
		Source: "https://baidu.com/path/to/example/pages?query=foo",
	}
	err := o.orm.Create(&m)
	o.Nil(err)
	o.NotEmpty(m.ID)
	o.NotEmpty(m.Hash)
}

func (o ormTestSuite) Test_Update() {
	update := ShortURLDO{
		ID:      1,
		Shorted: "shorted",
	}
	err := o.orm.Update(&update)
	o.Nil(err)
	o.NotEmpty(update.Shorted)
	o.NotEmpty(update.ID)
}

func (o ormTestSuite) Test_Query() {
	m := ShortURLDO{
		ID: 1,
	}
	err := o.orm.Query(&m)
	o.Nil(err)
	o.NotEmpty(m.ID)
	o.NotEmpty(m.Source)
	o.NotEmpty(m.Shorted)
	o.NotEmpty(m.Hash)
}

func Test_orm(t *testing.T) {
	suite.Run(t, new(ormTestSuite))
}
