package handlers

import (
	"net/http"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/models"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/services"
	"github.com/zenazn/goji/web"
)

// UserHandler hold the services used for login & auth
type UserHandler struct {
	*BaseHandler
	RS *services.Redis
	MS *services.MySQL
}

func NewUserHandler(b *BaseHandler, rs *services.Redis, ms *services.MySQL) *UserHandler {
	return &UserHandler{BaseHandler: b, RS: rs, MS: ms}
}

func (u *UserHandler) DoRegistration(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	//appErr := app.NewErr()

	// if len(email) == 0 {
	// 	appErr.MissingParametersErrors("email")
	// 	v.Error(&c, w, *appErr)
	// }

	user := models.User{}
	user.Name = r.FormValue("name")
	user.Mobile = r.FormValue("mobile")
	user.Password = r.FormValue("password")
	user.Blood = r.FormValue("blood")
	user.Sex = r.FormValue("sex")
	user.Lat = r.FormValue("lat")
	user.Lng = r.FormValue("lng")
	_, _, err := models.RegisterUser(&user, u.MS)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	/*user, userExist, err := v.MySql.Getuser(email)


	if !userExist {
		v.Error(&c, w, *appErr.NotFoundErrors("user"))
		return
	}
	v.Respond(w, 200, user)
	*/
	return nil
}
