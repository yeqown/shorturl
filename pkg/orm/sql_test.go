package orm

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ormTestSuite struct {
	suite.Suite

	orm *ShortORM
}

func (o *ormTestSuite) SetupSuite() {
	var (
		err error
		dsn = ""
	)
	o.orm, err = NewORM(dsn, "mysql")
	o.Require().Nil(err)
}

func (o ormTestSuite) Test_Create(t *testing.T) {
	m := ShortURLDO{
		Source: "https://baidu.com/path/to/example/pages?query=foo",
	}
	err := o.orm.Create(&m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.ID)
}

func (o ormTestSuite) Test_Update(t *testing.T) {
	update := ShortURLDO{
		ID:      1,
		Shorted: "shorted",
	}
	err := o.orm.Update(&update)
	assert.Nil(t, err)
	assert.NotEmpty(t, update.Source)
}

func (o ormTestSuite) Test_Query(t *testing.T) {
	m := ShortURLDO{
		ID: 1,
	}
	err := o.orm.Query(&m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.ID)
	assert.NotEmpty(t, m.Source)
	assert.NotEmpty(t, m.Shorted)
}
