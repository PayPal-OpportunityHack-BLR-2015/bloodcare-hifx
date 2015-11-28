package handlers

import (
	"net/http"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/services"
	"github.com/zenazn/goji/web"
)

//BloodBankHandler holds the services used for showing the dashboard
type BloodBankHandler struct {
	*BaseHandler
	DB *services.MySQL
}

func NewBloodBankHandler(bh *BaseHandler, db *services.MySQL) *BloodBankHandler {
	return &BloodBankHandler{BaseHandler: bh, DB: db}
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

//TODO:unit tests
func (f *BloodBankHandler) FetchBloodBanks(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	return nil
}

//TODO:unit tests
func (f *BloodBankHandler) InsertBloodBank(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	return nil
}

//TODO:unit tests
func (f *BloodBankHandler) ShowOneBloodBank(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	return nil
}
