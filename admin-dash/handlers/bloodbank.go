package handlers

import (
	"net/http"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/models"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/services"
	"github.com/zenazn/goji/web"
)

//BloodBankHandler holds the services used for showing the dashboard
type BloodBankHandler struct {
	*BaseHandler
	CS *services.Cassandra
	MN *services.Minion
}

func NewBloodBankHandler(bh *BaseHandler, cs *services.Cassandra, mn *services.Minion) *BloodBankHandler {
	return &BloodBankHandler{BaseHandler: bh, CS: cs, MN: mn}
}

//ShowBloodBankForm shows the  input form
func (f *BloodBankHandler) ShowBloodBankForm(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	data := f.getTplVars(c)
	data["selBloodBank"] = "active"

	err := f.RenderTpl(w, "bloodbank.html", &data)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	return nil
}

//FetchBloodBanks shows the bloodbanks
func (f *BloodBankHandler) FetchBloodBanks(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	data := f.getTplVars(c)

	page := r.FormValue("page")
	bloodbank, _, _ := models.FetchBloodBankDetails(id, date, f.CS)

	err := f.RenderTpl(w, "bloodbank.html", &data)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	return nil
}
