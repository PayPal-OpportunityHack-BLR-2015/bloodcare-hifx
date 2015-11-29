package models

import (
	"fmt"

	"github.com/paulmach/go.geo"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/services"
)

//BloodBank model
type BloodBank struct {
	ID       string
	Name     string
	Type     string
	Location *geo.Point
	Created  string
}
type BloodBanks []*BloodBank

func (a *BloodBank) String() string {
	return fmt.Sprintf("id:%s", a.ID)
}

func ListBloodBanks(db *services.MySQL) (*BloodBanks, error) {
	const (
		BBANK_LIST_SQL = "SELECT id, name, type, location, created FROM organisations"
	)
	var (
		results  BloodBanks
		id       string
		name     string
		typ      string
		location *geo.Point
		created  string
	)
	rows, err := db.Query(BBANK_LIST_SQL)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&id, &name, &typ, &location, &created)
		results = append(results,
			&BloodBank{
				ID:       id,
				Name:     name,
				Type:     typ,
				Location: location,
				Created:  created,
			})
	}
	return &results, nil

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
