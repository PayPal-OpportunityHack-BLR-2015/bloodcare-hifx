package handlers

import (
	"net/http"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/services"
	"github.com/zenazn/goji/web"
)

// VolunteerHandler hold the services used for login & auth
type VolunteerHandler struct {
	*BaseHandler
	RS *services.Redis
	MS *services.MySQL
}

func NewVolunteerHandler(b *BaseHandler, rs *services.Redis, ms *services.MySQL) *VolunteerHandler {
	return &VolunteerHandler{BaseHandler: b, RS: rs, MS: ms}
}

func (v *VolunteerHandler) DoLogin(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	appErr := app.NewErr()
	email := r.FormValue("email")

	if len(email) == 0 {
		appErr.MissingParametersErrors("email")
		v.Error(&c, w, *appErr)
	}

	volunteer, volunteerExist, err := v.MySql.GetVolunteer(email)
	if err != nil {
		v.Error(&c, w, app.InternalServerError.SetErr(err.Error()))
		return
	}

	if !volunteerExist {
		v.Error(&c, w, *appErr.NotFoundErrors("volunteer"))
		return
	}
	v.Respond(w, 200, volunteer)
	return nil
}
