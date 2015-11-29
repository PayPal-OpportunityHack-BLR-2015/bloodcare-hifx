package handlers

import (
	"net/http"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/models"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/services"
	"github.com/zenazn/goji/web"
)

type RequestHandler struct {
	*BaseHandler
	RS *services.Redis
	MS *services.MySQL
}

func NewRequestHandler(b *BaseHandler, rs *services.Redis, ms *services.MySQL) *RequestHandler {
	return &RequestHandler{BaseHandler: b, RS: rs, MS: ms}
}

func (u *RequestHandler) MakeBloodRequest(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
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
func (u *RequestHandler) RemoveBloodRequest(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	bloodReq := models.BloodRequest{}
	bloodReq.ReqId = r.FormValue("req_id")
	bloodReq.UserId = r.FormValue("user_id")
	_, _, err := models.DeleteBloodRequest(&bloodReq, u.MS)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	return nil //TODO  success
}

func (u *RequestHandler) GetRequestDetails(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	bloodReq := models.BloodRequest{}
	bloodReq.ReqId = r.FormValue("req_id")
	status, bloodRes, err := models.GetBloodRequest(&bloodReq, u.MS)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	if status != false {
		u.Respond(w, 200, bloodRes)
	}
	u.NotFound(c, w, r)
	return nil
}
