package main

import (
	"fmt"
	"os"

	"github.com/HiFX/env_parser"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/handlers"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/middleware"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/services"
	"github.com/Sirupsen/logrus"
	"github.com/goji/glogrus"
	"github.com/zenazn/goji"
	gmiddleware "github.com/zenazn/goji/web/middleware"
)

var (
	logr    *logrus.Logger
	appName string
)

func init() {
	logr = logrus.New()
	//	logr.Formatter = new(logrus.JSONFormatter)

	appName = "OHACK_API"

	goji.Abandon(gmiddleware.Logger)             //Remove default logger
	goji.Abandon(gmiddleware.Recoverer)          //Remove default Recoverer
	goji.Use(middleware.RequestIDHeader)         //Add RequestIDHeader Middleware
	glogrus := glogrus.NewGlogrus(logr, appName) //Add custom logger Middleware
	goji.Use(glogrus)
	goji.Use(middleware.NewRecoverer(logr)) //Add custom recoverer

}

func main() {
	//Profiling
	// go func() {
	// 	log.Println(http.ListenAndServe(":6060", nil))
	// }()

	var (
		config *app.Config
	)

	envParser := env_parser.NewEnvParser()
	envParser.Name(appName)
	envParser.Separator("_")
	envSrc := app.Envs{}
	envParseError := envParser.Map(&envSrc)
	app.Chk(envParseError)

	app.PrintWelcome()

	switch envSrc.Mode {
	case app.MODE_DEV:
		logr.Level = logrus.InfoLevel
	case app.MODE_PROD:
		logr.Level = logrus.WarnLevel
	case app.MODE_DEBUG:
		logr.Level = logrus.DebugLevel
	}

	config = app.NewConfig(envSrc.AssetsUrl, envSrc.UploadPath)

	logFile, fileError := os.OpenFile(envSrc.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	defer logFile.Close()
	if fileError == nil {
		logr.Out = logFile
	} else {
		fmt.Println("invalid log file; \n, Error : ", fileError, "\nopting standard output..")
	}
	redisService, _ := services.NewRedis(envSrc.RedisUrl)
	// app.Chk(reErr)

	sqlConnectionStringFormat := "%s:%s@tcp(%s:%s)/%s"
	sqlConnectionString := fmt.Sprintf(sqlConnectionStringFormat, envSrc.MysqlUser, envSrc.MysqlPassword,
		envSrc.MysqlHost, envSrc.MysqlPort, envSrc.MysqlDbName)
	mySqlService := services.NewMySQL(sqlConnectionString, 10)

	//TODO check
	baseHandler := handlers.NewBaseHandler(logr, config)
	userHandler := handlers.NewUserHandler(baseHandler, redisService, mySqlService)

	goji.Post("/register", baseHandler.Route(userHandler.DoRegistration))
	goji.NotFound(baseHandler.NotFound)

	goji.Serve()
}
