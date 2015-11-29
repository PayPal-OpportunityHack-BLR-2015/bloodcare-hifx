package handlers

import (
	"net/http"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/models"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/services"
	"github.com/zenazn/goji/web"
)

//DonorHandler holds the services used for showing the dashboard
type DonorHandler struct {
	*BaseHandler
	DB *services.MySQL
}

func NewDonorHandler(bh *BaseHandler, db *services.MySQL) *DonorHandler {
	return &DonorHandler{BaseHandler: bh, DB: db}
}


//TODO:unit tests
func (b *DonorHandler) FetchDonors(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	data := b.getTplVars(c)
	data["selDonor"] = "active"

	donors, err := models.ListDonors(b.DB)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	data["donors"] = donors

	err = b.RenderTpl(w, "donors.html", &data)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	return nil
}
