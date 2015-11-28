package handlers

import (
	"net/http"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/models"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/services"
	"github.com/zenazn/goji/web"
)

// AdminHandler hold the services used for login & auth
type AdminHandler struct {
	*BaseHandler
	CS *services.Cassandra
}

//NewAdminHandler is the AdminHandler constructor
func NewAdminHandler(b *BaseHandler, cs *services.Cassandra) *AdminHandler {
	return &AdminHandler{BaseHandler: b, CS: cs}
}

//ShowLogin displays the login screen
func (a *AdminHandler) ShowLogin(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	data := map[string]string{"assetsurl": a.Config.AssetsUrl}
	err := a.RenderTpl(w, "login.html", &data)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	return nil
}

//DoLogin performs the login
func (a *AdminHandler) DoLogin(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	email := r.FormValue("email")
	pass := r.FormValue("password")
	admin, msg, err := models.AuthAdmin(email, pass, a.CS)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	if msg != nil {
		data := map[string]string{"assetsurl": a.Config.AssetsUrl, "Msg": msg.Message}
		err := a.RenderTpl(w, "login.html", &data)
		if err != nil {
			return app.InternalServerError.SetErr(err.Error())
		}
		return nil
	}

	session := map[string]string{"userid": admin.ID, "username": admin.Name}
	a.setSession(session, w)

	http.Redirect(w, r, "/", http.StatusFound)
	return nil
}

//DoLogout logs the user out
func (a *AdminHandler) DoLogout(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	cookie := &http.Cookie{
		Name:   app.COOKIE_NAME,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", http.StatusFound)
	return nil
}

//ShowAdmins shows the administrators list
func (a *AdminHandler) ShowAdmins(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	data := a.getTplVars(c)
	data["selAdministrators"] = "active"

	admins, err := models.ListAdmins(a.CS)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	data["adminsList"] = admins

	err = a.RenderTpl(w, "admins.html", &data)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	return nil
}
