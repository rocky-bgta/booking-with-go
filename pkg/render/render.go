package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/rocky_bgta/booking-with-go/models"
	"github.com/rocky_bgta/booking-with-go/pkg/config"
)

var app *config.AppConfig

// NewTemplate initializes the app config for template rendering
func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateDate) *models.TemplateDate {
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateDate) {

	var tc map[string]*template.Template

	if app.UseCache {
		// get template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all the files named *page.gohtml from ./templates
	pages, err := filepath.Glob("./templates/*.gohtml")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.gohtml
	for _, page := range pages {
		name := filepath.Base(page)
		actualPage, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		match, err := filepath.Glob("./templates/*layout.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(match) > 0 {
			// Where They Get Linked
			actualPage, err = actualPage.ParseGlob("./templates/*layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = actualPage
	}

	return myCache, nil
}
