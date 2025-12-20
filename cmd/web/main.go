package main

import (
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
}

type appConfig struct {
	useCache         bool
	templateBasePath string
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}

	// get command line arguments (args passed to application)
	flag.BoolVar(&app.config.useCache, "cache", false, "set true for use the cache templates")
	flag.StringVar(&app.config.templateBasePath, "templatePath", "templates", "path of templates file")
	flag.Parse()

	log.Println("html templates are in: ", app.config.templateBasePath)

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
