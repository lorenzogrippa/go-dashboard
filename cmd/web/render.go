package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type templateData struct {
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, tname string, t *templateData) {

	var tmpl *template.Template

	if app.config.useCache {
		if templateFromMap, ok := app.templateMap[tname]; ok {
			tmpl = templateFromMap
		}
	}

	if tmpl == nil {
		newTemplateFromMap, err := app.buildTemplateFromDisk(tname)

		if err != nil {
			log.Println("Error: template ", tname, " -> ", err.Error())

			return
		}

		log.Println("template build from disk")

		tmpl = newTemplateFromMap
	}

	if t == nil {
		t = &templateData{}
	}

	if err := tmpl.ExecuteTemplate(w, tname, t); err != nil {
		log.Println("Error executing template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) buildTemplateFromDisk(t string) (*template.Template, error) {
	templateSlice := []string{
		fmt.Sprintf("%s/base.layout.gohtml", app.config.templateBasePath),
		fmt.Sprintf("%s/partials/head.partial.gohtml", app.config.templateBasePath),
		fmt.Sprintf("%s/partials/navigation.partial.gohtml", app.config.templateBasePath),
		fmt.Sprintf("%s/partials/footer.partial.gohtml", app.config.templateBasePath),

		fmt.Sprintf("%s/%s", app.config.templateBasePath, t),
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}

	app.templateMap[t] = tmpl

	return tmpl, nil
}
