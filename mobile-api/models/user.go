package models

import (
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/services"
	"golang.org/x/crypto/bcrypt"
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

func ValidateUser(u *User, ms *services.MySQL) (bool, string) {
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

	if PhoneExists(u.Mobile, ms) == true {
		return false, "Mobile number exists"
	}

	return true, ""
}

func PhoneExists(mobile string, ms *services.MySQL) bool {
	query := "SELECT id FROM users WHERE mobile=?"
	res, _ := ms.Query(query, mobile)

	if res.Next() {
		var userId int
		res.Scan(&userId)
		if userId > 0 {
			return true
		}
	}
	return false
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

func RegisterUser(u *User, ms *services.MySQL) (int64, *app.Msg, error) {

	// const (
	// 	ADMIN_AUTH_SQL = "SELECT id, name, password  FROM users WHERE email=?"
	// )
	// var id, name, bcryptpass string
	status, error := ValidateUser(u, ms)
	if status == false {
		return 0, app.NewErrMsg(error), nil
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	query := "INSERT INTO users(name, mobile, password, blood, sex, location, place_id) VALUES(?,?,?,?,?, POINT(" + u.Lat + ", " + u.Lng + "),?)"
	res, dbError := ms.Exec(query, u.Name, u.Mobile, string(passwordHash[:]), u.Blood, u.Sex, u.PlaceId)
	if dbError != nil {
		return 0, nil, dbError
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return 0, nil, err
		} else {
			return id, nil, nil
		}
	}

	return 0, app.NewErrMsg(error), nil
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
