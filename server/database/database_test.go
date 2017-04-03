package database

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/murphybytes/journal/server/journal"
	"github.com/stretchr/testify/suite"
)

type DatabaseTestSuite struct {
	suite.Suite
	db   journal.Database
	skip bool
}

func (d *DatabaseTestSuite) SetupSuite() {
	var err error
	d.db, err = New("Server=127.0.0.1;Uid=root;Pwd=toor;Database=journal;")
	if err != nil {
		d.skip = true
		d.T().Log("could not connect to db, ", err)
	}
}

func (d *DatabaseTestSuite) TearDownSuite() {
	if d.db != nil {
		d.db.Close()
	}
}

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}
