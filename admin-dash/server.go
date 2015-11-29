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
		mode      string
		mysqlHost string
		redisHost string
		assetsUrl string
		config    *app.Config
		errors    []string
	)
	flag.StringVar(&mode, "mode", "dev", "dev|debug|prod")
	flag.StringVar(&assetsUrl, "assets-url", "", "http://static.assets.url")
	flag.StringVar(&mysqlHost, "mysql-host", "", "mysql.host")
	flag.StringVar(&redisHost, "redis-host", "", "redis.host")
	app.EnvParse(appName)

	if assetsUrl == "" {
		errors = append(errors, "assets-url")
	}
	if mysqlHost == "" {
		errors = append(errors, "mysql-host")
	}
	if redisHost == "" {
		errors = append(errors, "redis-host")
	}

	app.PrintWelcome()
	app.ParseErrors(errors)

	logr.Level = app.GetLogrMode(mode)
	config = app.NewConfig(assetsUrl)
	mysql := services.NewMySQL(mysqlHost, 100)


	bH := handlers.NewBaseHandler(logr, config, templates, cookieHandler)

	adminH := handlers.NewAdminHandler(bH, mysql)
	dashH := handlers.NewDashboardHandler(bH, mysql)
	bloodbankH := handlers.NewBloodBankHandler(bH, mysql)

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
