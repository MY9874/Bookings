package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/moyu/bookings/internal/config"
	"github.com/moyu/bookings/internal/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig
var pathToTemplates = "./templates"

// NewRenderer sets the config for thr template package
func NewRenderer(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// tmpl is the name of template we want to render, using html template
func Template(w http.ResponseWriter, r *http.Request, html string, td *models.TemplateData) error {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// get the template cache from the app config

	//get the requested template from cache
	t, ok := tc[html]
	if !ok {
		return errors.New("can't get template from cache")
	}

	// to tell that the error comes from the value that stored in the map
	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)
	err := t.Execute(buf, td)
	if err != nil {
		fmt.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

	return nil
}

// we no longer need to keep track of what html files are inside and add them manually to this function
func CreateTemplateCache() (map[string]*template.Template, error) {
	// myChache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{} //same as last line

	// get all of the files named *.page.html from ./templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplates)) // go to some location and look for the files
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages { // does not care the index
		name := filepath.Base(page) // page is the full path, here base helps get file name
		// ts: template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page) // parse the file 'page' and store it in template 'name'
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.html", pathToTemplates)) // parse .layout files and add to ts
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts // name now should be 'about.page.html' or 'home.page.html'
		// here the html files comes with the layout files

	}

	return myCache, nil

}

// approach1, good but not enough, need to add layout and template path manually
/*
var tc = make(map[string]*template.Template) // template cache, package level variable

func RenderTemplate(w http.ResponseWriter, t string) { //t is the template we want to render
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := tc[t] //pop it to _, and inMap will be True or False depends on if it exist
	if !inMap {
		//need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		//we have the template in the cache
		log.Println("using cached template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}

	//parse the template
	tmpl, err := template.ParseFiles(templates...) //take each entry of the slice and put them in as individual strings
	if err != nil {
		return err
	}

	// add template to cache
	tc[t] = tmpl

	return nil

}


*/
