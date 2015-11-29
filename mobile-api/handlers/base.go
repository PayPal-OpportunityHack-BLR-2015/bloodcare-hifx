package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/app"
	"github.com/Sirupsen/logrus"
	"github.com/oxtoacart/bpool"
	"github.com/zenazn/goji/web"
	gmiddleware "github.com/zenazn/goji/web/middleware"
)

//AppH is the signature of our handlers. In addition to goji's handler type
//AppH returns an app.Err
type AppH func(web.C, http.ResponseWriter, *http.Request) *app.Err

// BaseHandler is the base handler. Other handlers embeds this one
type BaseHandler struct {
	Logr    *logrus.Logger
	Config  *app.Config
	bufpool *bpool.BufferPool
}

//NewBaseHandler is the BaseHandler constructor
func NewBaseHandler(l *logrus.Logger, c *app.Config) *BaseHandler {
	return &BaseHandler{Logr: l, Config: c, bufpool: bpool.NewBufferPool(64)}
}

//Route is used because our handlers returns *app.Err
func (b *BaseHandler) Route(h AppH) func(web.C, http.ResponseWriter, *http.Request) {
	fn := func(c web.C, w http.ResponseWriter, r *http.Request) {
		err := h(c, w, r)
		if err != nil {
			reqID := gmiddleware.GetReqID(c)
			b.Logr.WithFields(logrus.Fields{
				"req_id": reqID,
				"err":    err.Error(),
			}).Error("response.err")

			http.Error(w, http.StatusText(err.HTTPStatus), err.HTTPStatus)
		}
	}
	return fn
}

//NotFound
func (b *BaseHandler) NotFound(c web.C, w http.ResponseWriter, r *http.Request) {
	b.Respond(w, 404, []string{})
}

// Respond writes an HTTP response to the given resp,
func (b *BaseHandler) Respond(w http.ResponseWriter, status int, data interface{}) {
	d, err := json.MarshalIndent(data, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(d)
	if nil != err {
		b.Logr.Errorf("web.ioerror %s", err.Error())
	}
	_, err = w.Write([]byte("\n"))
	if nil != err {
		b.Logr.Errorf("web.ioerror %s", err.Error())
	}
}
