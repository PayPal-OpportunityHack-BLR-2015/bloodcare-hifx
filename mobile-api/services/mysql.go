package services

import (
	"database/sql"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/app"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(conString string, maxIdleCon int) *MySQL {
	myDB, dbError := sql.Open("mysql", conString)
	app.Chk(dbError)
	/*
	   test connection
	*/
	app.Chk(myDB.Ping())
	myDB.SetMaxIdleConns(maxIdleCon)

	pg := MySQL{db: myDB}
	return &pg
}

func (m *MySQL) Query(query string, values ...interface{}) (*sql.Rows, error) {
	return m.db.Query(query, values...)
}

func (m *MySQL) QueryRow(query string, values ...interface{}) *sql.Row {
	return m.db.QueryRow(query, values...)
}

func (m *MySQL) Exec(query string, values ...interface{}) (sql.Result, error) {
	return m.db.Exec(query, values...)
}
