package handlers

import (
	"net/http"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/models"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/services"
	"github.com/zenazn/goji/web"
)

// UserHandler hold the services used for login & auth
type Request struct {
	*BaseHandler
	RS *services.Redis
	MS *services.MySQL
}

func NewRequestHandler(b *BaseHandler, rs *services.Redis, ms *services.MySQL) *UserHandler {
	return &UserHandler{BaseHandler: b, RS: rs, MS: ms}
}

func (u *UserHandler) MakeBloodRequest(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	bloodReq := models.BloodRequest{}
	bloodReq.UserId = r.FormValue("user_id")
	bloodReq.Date = r.FormValue("date")
	bloodReq.Blood = r.FormValue("blood")
	bloodReq.Phone = r.FormValue("phone")
	bloodReq.Description = r.FormValue("description")
	bloodReq.Lat = r.FormValue("lat")
	bloodReq.Lng = r.FormValue("lng")
	bloodReq.PlaceId = r.FormValue("place_id")
	_, _, err := models.CreateBloodRequest(&bloodReq, u.MS)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	return nil //TODO  success
}
