package main

import (
	"flag"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/handlers"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/middleware"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/services"
	"github.com/Sirupsen/logrus"
	"github.com/goji/glogrus"
	"github.com/gorilla/securecookie"
	"github.com/saj1th/envtoflag"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	gmiddleware "github.com/zenazn/goji/web/middleware"
)

var (
	logr          *logrus.Logger
	appName       string
	templates     map[string]*template.Template
	cookieHandler *securecookie.SecureCookie
)

func init() {
	logr = logrus.New()
	//	logr.Formatter = new(logrus.JSONFormatter)

	appName = "bloodcare.dash"

	cookieHandler = securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32))

	goji.Abandon(gmiddleware.Logger)             //Remove default logger
	goji.Abandon(gmiddleware.Recoverer)          //Remove default Recoverer
	goji.Use(middleware.RequestIDHeader)         //Add RequestIDHeader Middleware
	glogrus := glogrus.NewGlogrus(logr, appName) //Add custom logger Middleware
	goji.Use(glogrus)
	goji.Use(middleware.NewRecoverer(logr)) //Add custom recoverer

	initTpl()
}

func initTpl() {
	templates = make(map[string]*template.Template)
	layouts, err := filepath.Glob("views/*.html") //Get the list of template files
	app.Chk(err)
	includes, err := filepath.Glob("views/includes/*.html") //Get the list of include files
	app.Chk(err)
	for _, layout := range layouts {
		files := append(includes, layout) // Parse the layout tpl file and all the includes
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles(files...))
	}
}

func main() {
	//Profiling
	// go func() {
	// 	log.Println(http.ListenAndServe(":6060", nil))
	// }()

	var (
		mode           string
		cassandra_host string
		redis_host     string
		upload_path    string
		assets_url     string
		config         *app.Config
		errors         []string
	)
	flag.StringVar(&mode, "mode", "dev", "dev|debug|prod")
	flag.StringVar(&assets_url, "assets-url", "", "http://static.assets.url")
	flag.StringVar(&cassandra_host, "cassandra-host", "", "cassandra.host")
	flag.StringVar(&redis_host, "redis-host", "", "redis.host")
	flag.StringVar(&upload_path, "upload-path", "", "upload.path")
	envtoflag.Parse(appName)

	if assets_url == "" {
		errors = append(errors, "assets-url")
	}
	if cassandra_host == "" {
		errors = append(errors, "cassandra-host")
	}
	if redis_host == "" {
		errors = append(errors, "redis-host")
	}
	if upload_path == "" {
		errors = append(errors, "upload-path")
	}
	app.PrintWelcome()
	app.ParseErrors(errors)

	logr.Level = app.GetLogrMode(mode)
	config = app.NewConfig(assets_url, upload_path)

	cassandra := services.NewCassandra(cassandra_host, "bloodcare")
	minions := services.NewMinion(redis_host)
	bH := handlers.NewBaseHandler(logr, config, templates, cookieHandler)

	adminH := handlers.NewAdminHandler(bH, cassandra)
	dashH := handlers.NewDashboardHandler(bH, cassandra)
	bloodbankH := handlers.NewBloodBankHandler(bH, cassandra, minions)

	goji.Get("/login/", http.RedirectHandler("/login", 301))
	goji.Get("/login", bH.Route(adminH.ShowLogin))
	goji.Post("/login", bH.Route(adminH.DoLogin))
	goji.NotFound(bH.NotFound)

	admin := web.New()
	goji.Handle("/*", admin)
	admin.NotFound(bH.NotFound)
	admin.Use(middleware.NewAuth(cookieHandler))

	admin.Get("/", bH.Route(dashH.ShowDashboard))
	admin.Get("/bloodbanks", bH.Route(bloodbankH.FetchBloodBanks))
	admin.Get("/bloodbank", bH.Route(bloodbankH.ShowBloodBankForm))
	admin.Post("/bloodbank", bH.Route(bloodbankH.InsertBloodBank))
	admin.Get("/bloodbank/:id", bH.Route(bloodbankH.ShowOneBloodBank))

	goji.Serve()
}
