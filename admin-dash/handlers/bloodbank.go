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
	DB *services.MySQL
}

func NewBloodBankHandler(bh *BaseHandler, db *services.MySQL) *BloodBankHandler {
	return &BloodBankHandler{BaseHandler: bh, DB: db}
}

//ShowBloodBankForm shows the  input form
func (b *BloodBankHandler) ShowBloodBankForm(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	data := b.getTplVars(c)
	data["selBloodBank"] = "active"

	err := b.RenderTpl(w, "bloodbank.html", &data)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	return nil
}

//TODO:unit tests
func (b *BloodBankHandler) FetchBloodBanks(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	data := b.getTplVars(c)
	data["selBloodBank"] = "active"

	bloodBanks, err := models.ListBloodBanks(b.DB)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	data["bloodBanks"] = bloodBanks

	err = b.RenderTpl(w, "bloodbanks.html", &data)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
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
