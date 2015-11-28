package handlers

import (
	"net/http"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/services"
	"github.com/zenazn/goji/web"
)

//DashboardHandler holds the services used for showing the dashboard
type DashboardHandler struct {
	*BaseHandler
	DB *services.MySQL
}

func NewDashboardHandler(bh *BaseHandler, db *services.MySQL) *DashboardHandler {
	return &DashboardHandler{BaseHandler: bh, DB: db}
}

//ShowDashboard shows the dashboard
func (d *DashboardHandler) ShowDashboard(c web.C, w http.ResponseWriter, r *http.Request) *app.Err {
	data := d.getTplVars(c)
	data["selDashboard"] = "active"
	err := d.RenderTpl(w, "dashboard.html", &data)
	if err != nil {
		return app.InternalServerError.SetErr(err.Error())
	}
	return nil
}
