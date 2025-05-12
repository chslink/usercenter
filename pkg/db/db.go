package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func New(conf *Config) (*bun.DB, error) {
	dbConf := conf.Data.Database
	sqlDb, err := sql.Open(dbConf.Driver, dbConf.Source)
	if err != nil {
		return nil, err
	}
	db := bun.NewDB(sqlDb, mysqldialect.New())

	return db, nil
}
