package render

import (
	"bytes"
	"html/template"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/rohitCodeRed/go_crypto/config"
	"github.com/rohitCodeRed/go_crypto/model"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *model.CoinData) *model.CoinData {
	return td
}

// RenderTemplate renders a template
func RenderTemplate(c *gin.Context, tmpl string, td *model.CoinData) {

	var tc map[string]*template.Template

	//tc, _ = CreateTemplate()
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplate()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template..")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)

	_ = t.Execute(buf, td)
	// if err != nil {
	// 	log.Println(err)
	// }

	// render the template
	_, err := buf.WriteTo(c.Writer)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplate() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
