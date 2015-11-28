package middleware

import (
	"net/http"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/gorilla/securecookie"
	"github.com/zenazn/goji/web"
)

type auth struct {
	h       http.Handler
	c       *web.C
	cookieH *securecookie.SecureCookie
}

func (a auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ok bool
	cookie, err := r.Cookie(app.COOKIE_NAME)
	if nil != err {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	session := make(map[string]string)
	err = a.cookieH.Decode(app.COOKIE_NAME, cookie.Value, &session)
	if nil != err {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if a.c.Env["userid"], ok = session["userid"]; !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if a.c.Env["username"], ok = session["username"]; !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	a.h.ServeHTTP(w, r)
}

// NewAuth returns the Auth middleware
func NewAuth(cookieHandler *securecookie.SecureCookie) func(*web.C, http.Handler) http.Handler {
	fn := func(c *web.C, h http.Handler) http.Handler {
		return auth{h: h, c: c, cookieH: cookieHandler}
	}
	return fn
}
