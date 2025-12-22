package main

import (
	"dashboard/configuration"
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	App         *configuration.Application
}

type appConfig struct {
	useCache         bool
	templateBasePath string
	staticResources  string
	dns              string // database connection string
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}

	// get command line arguments (args passed to application)
	initializeConfig(&app)

	// get db
	db, err := initDB(app.config.dns)

	if err != nil {
		log.Panic(err)
	}

	app.App = configuration.New(db)

	if app.config.useCache {
		log.Println("Ok use cache for the templates")
	}

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func initializeConfig(app *application) {
	flag.BoolVar(&app.config.useCache, "cache", false, "set true for use the cache templates")
	flag.StringVar(&app.config.templateBasePath, "templatePath", "./templates", "path of templates file")
	flag.StringVar(&app.config.staticResources, "staticResourcesPath", "./static/", "path of static resources")
	flag.StringVar(&app.config.dns, "dns", "mariadb:password@rcp(localhost:3306)/breeders?parseTime=true", "dns for database connection string")
	flag.Parse()

	log.Println("html templates are in:\t\t", app.config.templateBasePath)
	log.Println("static resources path are in:\t", app.config.staticResources)
	log.Println("dns db string:\t\t\t", app.config.dns)
}
