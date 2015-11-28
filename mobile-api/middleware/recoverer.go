package middleware

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/facebookgo/stack"
	"github.com/zenazn/goji/web"
	gmiddleware "github.com/zenazn/goji/web/middleware"
)

// Logr hold the logrus var
type recoverer struct {
	h http.Handler
	c *web.C
	l *logrus.Logger
}

func (rec recoverer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	reqID := gmiddleware.GetReqID(*rec.c)

	defer func() {
		if r := recover(); r != nil {
			var err error
			switch x := r.(type) {
			case error:
				err = x
			default:
				err = fmt.Errorf("%v", x)
			}
			rec.l.WithFields(logrus.Fields{
				"req_id": reqID,
				"method": req.Method,
				"uri":    req.RequestURI,
				"remote": req.RemoteAddr,
				"err":    err.Error(),
				"stack":  stack.Callers(0),
			}).Warn("service.err")
			http.Error(resp, http.StatusText(500), 500)
		}
	}()

	rec.h.ServeHTTP(resp, req)

}

// NewRecoverer returns recoverer
func NewRecoverer(l *logrus.Logger) func(*web.C, http.Handler) http.Handler {
	fn := func(c *web.C, h http.Handler) http.Handler {
		return recoverer{h, c, l}
	}
	return fn
}
