package models

import (
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/services"
)

type User struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Mobile   string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Blood    string `json:"blood,omitempty"`
	Sex      string `json:"sex,omitempty"`
	Lat      string `json:"lat,omitempty"`
	Lng      string `json:"lng,omitempty"`
	PlaceId  string `json:"placeId,omitempty"`
}
type Users []*User

func (u *User) String() string {
	return u.Name
}

func (u *User) ValidateUser() (bool, string) {
	if len(u.Name) == 0 {
		return false, "Name is empty"
	}

	if len(u.Mobile) == 0 {
		return false, "Mobile is empty"
	}

	if len(u.Password) == 0 {
		return false, "Password is empty"
	}

	if IsValidBloodType(u.Blood) == false {
		return false, "Invalid blood type"
	}

	return true, ""
}

func IsValidBloodType(bloodType string) bool {
	switch bloodType {
	case
		"A+",
		"A−",
		"B+",
		"B−",
		"A1B+",
		"AB+",
		"AB−",
		"O+",
		"O−":
		return true
	}
	return false
}

func RegisterUser(u *User, ms *services.MySQL) (bool, *app.Msg, error) {

	// const (
	// 	ADMIN_AUTH_SQL = "SELECT id, name, password  FROM users WHERE email=?"
	// )
	// var id, name, bcryptpass string
	status, error := u.ValidateUser()
	if status == false {
		return false, app.NewErrMsg(error), nil
	}
	query := "INSERT INTO users(name, mobile, password, blood, sex, lat, lng, place_id) VALUES(?,?,?,?,?,?,?,?)"
	_, dbError := ms.Exec(query, u.Name, u.Mobile, u.Password, u.Blood, u.Sex, u.Lat, u.Lng, u.PlaceId)
	return false, app.NewErrMsg(error), nil
	/*
		dbResult, dbError := ms.Query(ADMIN_AUTH_SQL, email)
		if dbError != nil {
			return nil, nil, dbError
		}
		defer dbResult.Close()
		if dbResult.Next() {
			var user *User
			dbResult.Scan(user.Id, user.Name, user.Password)
			return user, nil, nil
		} else {
			return nil, app.NewErrMsg("The email or password is incorrect."), nil
		}*/
}
