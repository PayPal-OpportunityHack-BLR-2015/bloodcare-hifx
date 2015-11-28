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

/*func (m *MySQL) GetVolunteer(email string) (models.Volunteer, bool, error) {

	var volunteer models.Volunteer
	dbResult, dbError := m.db.Query("SELECT id, name, email, password "+
		"FROM users WHERE email = ?", email)

	if dbError != nil {
		return volunteer, false, dbError
	}
	defer dbResult.Close()

	if dbResult.Next() {
		dbResult.Scan(&volunteer.Id, &volunteer.Name, &volunteer.Password)
		return volunteer, true, nil
	} else {
		return volunteer, false, nil
	}
}
*/
