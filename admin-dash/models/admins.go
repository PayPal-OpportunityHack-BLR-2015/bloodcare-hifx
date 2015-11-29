package models

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/services"
)

//Admin models an admin user
type Admin struct {
	ID       string
	Email    string
	Name     string
	Password string
	Status   string
}

type Admins []*Admin

func (a *Admin) String() string {
	return a.Name
}

func AuthAdmin(email, pass string, db *services.MySQL) (*Admin, *app.Msg, error) {
	const (
		ADMIN_AUTH_SQL = "SELECT id, name, password, status  FROM admin_users WHERE email=?"
	)
	var id, name, bcryptpass, status string

	if len(email) == 0 || len(pass) == 0 {
		return nil, app.NewErrMsg("The email or password is empty."), nil
	}

	rows, err := db.Query(ADMIN_AUTH_SQL, email)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, app.NewErrMsg("The email or password is incorrect."), nil
	}

	rows.Scan(&id, &name, &bcryptpass, &status)
	perr := bcrypt.CompareHashAndPassword([]byte(bcryptpass), []byte(pass))
	if perr != nil {
		return nil, app.NewErrMsg("The email or password is incorrect."), nil
	}
	if status == "inactive" {
		return nil, app.NewErrMsg("Please contact sysadmin"), nil
	}
	return &Admin{ID: id, Name: name, Email: email}, nil, nil
}

func ListAdmins(db *services.MySQL) (*Admins, error) {
	const (
		ADMINS_LIST_SQL = "SELECT id, name, email, password, status FROM admin_users"
	)
	var (
		results  Admins
		id       string
		name     string
		email    string
		password string
		status   string
	)
	rows, err := db.Query(ADMINS_LIST_SQL)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&id, &name, &email, &password, &status)
		results = append(results,
			&Admin{
				ID:       id,
				Name:     name,
				Email:    email,
				Password: password,
				Status:   status,
			})
	}
	return &results, nil
}
