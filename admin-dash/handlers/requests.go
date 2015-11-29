package handlers

import (
	"net/http"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/models"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/services"
	"github.com/zenazn/goji/web"
)

//RequestHandler holds the services used for showing the dashboard
type RequestHandler struct {
	*BaseHandler
	DB *services.MySQL
}

func NewRequestHandler(bh *BaseHandler, db *services.MySQL) *RequestHandler {
	return &RequestHandler{BaseHandler: bh, DB: db}
}


//TODO:unit tests
func (b *RequestHandler) FetchRequests(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	data := b.getTplVars(c)
	data["selRequest"] = "active"

	requests, err := models.ListRequests(b.DB)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	data["requests"] = requests

	err = b.RenderTpl(w, "requests.html", &data)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	return nil
}
