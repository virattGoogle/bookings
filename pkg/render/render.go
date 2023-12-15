package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/virattGoogle/bookings/pkg/config"
	"github.com/virattGoogle/bookings/pkg/models"
)
var app *config.Appconfig
func Newtemplate(a *config.Appconfig){
	app =a
}

func RenderTemplate(w http.ResponseWriter,r *http.Request, tmpl string,td *models.Templatedata) {

	tc:= app.TemplateCache

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not get template")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultdata(td,r)

	err := t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func AddDefaultdata(td *models.Templatedata, r *http.Request) *models.Templatedata {
	td.CSRFToken =nosurf.Token(r)
	
	return td
}

func Createttemplatecache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./Templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./Templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./Templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil

}
