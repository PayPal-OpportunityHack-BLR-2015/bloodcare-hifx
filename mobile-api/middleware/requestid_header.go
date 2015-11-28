package middleware

import (
	"net/http"

	"github.com/zenazn/goji/web"
	gmiddleware "github.com/zenazn/goji/web/middleware"
)

// RequestIDHeader adds a Request-Id in response header.
func RequestIDHeader(c *web.C, h http.Handler) http.Handler {
	fn := func(resp http.ResponseWriter, req *http.Request) {
		var requestID string
		given := req.Header.Get("X-Request-Id")
		if given != "" {
			requestID = given
		}
		given = req.Header.Get("Request-Id")
		if given != "" {
			requestID = given
		}
		if requestID == "" && c.Env["requestID"] != "" {
			requestID = gmiddleware.GetReqID(*c)
		}
		resp.Header().Set("Request-Id", requestID)
		h.ServeHTTP(resp, req)
	}
	return http.HandlerFunc(fn)
}
