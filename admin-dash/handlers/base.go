package handlers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/securecookie"
	"github.com/oxtoacart/bpool"
	"github.com/zenazn/goji/web"
	gmiddleware "github.com/zenazn/goji/web/middleware"
)

//AppH is the signature of our handlers. In addition to goji's handler type
//AppH returns an app.Err
type AppH func(web.C, http.ResponseWriter, *http.Request) *app.Err

// BaseHandler is the base handler. Other handlers embeds this one
type BaseHandler struct {
	Logr      *logrus.Logger `inject:"private"`
	Templates map[string]*template.Template
	Config    *app.Config
	bufpool   *bpool.BufferPool
	cookieH   *securecookie.SecureCookie
}

//NewBaseHandler is the BaseHandler constructor
func NewBaseHandler(l *logrus.Logger, c *app.Config, tpls map[string]*template.Template, cookieHandler *securecookie.SecureCookie) *BaseHandler {
	return &BaseHandler{Logr: l, Config: c, Templates: tpls, bufpool: bpool.NewBufferPool(64), cookieH: cookieHandler}
}

// Respond writes an HTTP response to the given resp,
func (b *BaseHandler) Respond(w http.ResponseWriter, status int, data string) {

	w.WriteHeader(status)
	_, err := w.Write([]byte(data))
	if nil != err {
		b.Logr.Errorf("web.ioerror %s", err.Error())
	}
}

//Render renders the template
func (b *BaseHandler) Render(w http.ResponseWriter, name, template string, data interface{}) error {
	tpl, ok := b.Templates[name]
	if !ok {
		return app.InternalServerError.SetErr("Tpl does not exist:" + name)
	}

	// Create a buffer to write to and check if any errors were encountered.
	buf := b.bufpool.Get()
	err := tpl.ExecuteTemplate(buf, template, data)
	if err != nil {
		b.bufpool.Put(buf)
		return app.InternalServerError.SetErr(err.Error())
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
	b.bufpool.Put(buf)
	return nil
}

//RenderTpl renders the base template
func (b *BaseHandler) RenderTpl(w http.ResponseWriter, name string, data interface{}) error {
	return b.Render(w, name, "base", data)
}

//RenderModal renders the modal template
func (b *BaseHandler) RenderModal(w http.ResponseWriter, name string, data interface{}) error {
	return b.Render(w, name, "modal", data)
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
	w.WriteHeader(http.StatusNotFound)
	data := map[string]string{"assetsurl": b.Config.AssetsUrl}
	b.RenderTpl(w, "404.html", data)
}

func (b *BaseHandler) setSession(session map[string]string, w http.ResponseWriter) {
	if encoded, err := b.cookieH.Encode(app.COOKIE_NAME, session); err == nil {
		cookie := &http.Cookie{
			Name:  app.COOKIE_NAME,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func (b *BaseHandler) getAllSession(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie(app.COOKIE_NAME)
	if nil != err {
		return nil, err
	}
	session := make(map[string]string)
	err = b.cookieH.Decode(app.COOKIE_NAME, cookie.Value, &session)
	if nil != err {
		return nil, err
	}
	return session, nil
}

func (b *BaseHandler) getSession(key string, r *http.Request) (string, error) {
	var (
		val string
		ok  bool
	)
	session, err := b.getAllSession(r)
	if err != nil {
		return "", err
	}
	if val, ok = session[key]; !ok {
		return "", errors.New("Key not found")
	}
	return val, nil
}

func (b *BaseHandler) getTplVars(c web.C) map[string]interface{} {
	return map[string]interface{}{"assetsurl": b.Config.AssetsUrl, "username": c.Env["username"].(string), "userid": c.Env["userid"].(string)}
}
