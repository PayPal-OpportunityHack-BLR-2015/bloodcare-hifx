package models

import (
	"fmt"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/services"
	"github.com/paulmach/go.geo"
)

//Donor model
type Donor struct {
	ID       string
	Name     string
	Mobile   string
	Blood    string
	Gender   string
	Location *geo.Point
	Place    string
	Created  string
}
type Donors []*Donor

func (a *Donor) String() string {
	//TODO: Fix :(
	return fmt.Sprintf("id:%s", a.ID)
}

func ListDonors(db *services.MySQL) (*Donors, error) {
	const (
		DONORS_LIST_SQL = "SELECT id, name, mobile, blood, sex, location, place_id, created FROM users"
	)
	var (
		results  Donors
		id       string
		name     string
		mobile   string
		blood    string
		gender   string
		location *geo.Point
		place    string
		created  string
	)
	rows, err := db.Query(DONORS_LIST_SQL)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&id, &name, &mobile, &blood, &gender, &location, &place, &created)
		results = append(results,
			&Donor{
				ID:       id,
				Name:     name,
				Mobile:   mobile,
				Blood:    blood,
				Gender:   gender,
				Location: location,
				Place:    place,
				Created:  created,
			})
	}
	return &results, nil

}

func InsertDonor(name string, db *services.MySQL) (string, *app.Msg) {
	const (
		BBANK_INSERT_SQL = "INSERT INTO Donors () VALUES (?, ?, ?, ?)"
	)
	return "", nil
}

func UpdateDonor(id int, name string, db *services.MySQL) (string, *app.Msg) {
	const (
		BBANK_UPDATE_SQL = ""
	)
	return "", nil
}

func FetchDonorDetails(id int, db *services.MySQL) (string, *app.Msg) {
	const (
		JOB_FETCH_SQL = "SELECT (stage) FROM bloodcare_jobs WHERE id=? AND userId=?"
	)
	return "", nil
}
