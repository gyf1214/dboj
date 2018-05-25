package model

import (
	"database/sql"

	// sql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gyf1214/dboj/util"
)

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("mysql", util.Database)
	if err != nil {
		panic(err)
	}
}
