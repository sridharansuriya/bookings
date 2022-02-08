package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/sridharansuriya/bookings/pkg/config"
	"github.com/sridharansuriya/bookings/pkg/models"
)

var (
	functions = template.FuncMap{}
	app       *config.AppConfig
)

func SetAppConfig(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders a template on to the web page
func RenderTemplate(rw http.ResponseWriter, templateName string, td *models.TemplateData) {
	// getting this from app config is the ideal way
	var tc map[string]*template.Template
	var err error
	if !app.UseCache {
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		tc = app.TemplateCache
	}
	buf := new(bytes.Buffer)
	t, ok := tc[templateName]
	if !ok {
		log.Fatal("Unable to find template in cache")
	}
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	_, err = buf.WriteTo(rw)

	if err != nil {
		log.Println(err)
	}
}

// Creates a template cache that can be used to render templates
func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := make(map[string]*template.Template)

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return templateCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return templateCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[name] = ts
	}
	return templateCache, nil
}
