package models

import (
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/services"
)

type Volunteer struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
type Volunteers []*Volunteer

func (v *Volunteer) String() string {
	return v.Name
}

func Sbd(email, pass string, ms *services.MySQL) (*Volunteer, *app.Msg, error) {
	const (
		ADMIN_AUTH_SQL = "SELECT id, name, password  FROM users WHERE email=?"
	)
	// var id, name, bcryptpass string

	if len(email) == 0 || len(pass) == 0 {
		return nil, app.NewErrMsg("The email or password is empty."), nil
	}
	dbResult, dbError := ms.Query(ADMIN_AUTH_SQL, email)
	if dbError != nil {
		return nil, nil, dbError
	}
	defer dbResult.Close()
	if dbResult.Next() {
		var volunteer *Volunteer
		dbResult.Scan(volunteer.Id, volunteer.Name, volunteer.Password)
		return volunteer, nil, nil
	} else {
		return nil, app.NewErrMsg("The email or password is incorrect."), nil
	}
}
