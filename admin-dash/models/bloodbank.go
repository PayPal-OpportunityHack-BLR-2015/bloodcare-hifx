package models

import (
	"fmt"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/services"
)

//BloodBank model
type BloodBank struct {
	Id   int
	Name string
	// Add Fields here
}

func (a *BloodBank) String() string {
	return fmt.Sprintf("id:%s", a.Id)
}

func BloodBanks(page int, db *services.MySQL) (*BloodBank, *app.Msg, error) {

	return &BloodBank{}, nil, nil
}

func InsertBloodBank(name string, db *services.MySQL) (string, *app.Msg) {
	const (
		BBANK_INSERT_SQL = "INSERT INTO bloodbanks () VALUES (?, ?, ?, ?)"
	)
	return "", nil
}

func UpdateBloodBank(id int, name string, db *services.MySQL) (string, *app.Msg) {
	const (
		BBANK_UPDATE_SQL = ""
	)
	return "", nil
}

func FetchBloodBankDetails(id int, db *services.MySQL) (string, *app.Msg) {
	const (
		JOB_FETCH_SQL = "SELECT (stage) FROM bloodcare_jobs WHERE id=? AND userId=?"
	)
	return "", nil
}
