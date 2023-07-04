package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Doaa-Ismail/go_course/booking/pkg/config"
	"github.com/Doaa-Ismail/go_course/booking/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplates(w http.ResponseWriter, html string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateRenderTemplatesTest()
		// if err != nil {
		// 	fmt.Println("Error Getting Cache Page : ", err)
		// }
	}
	t, ok := tc[html]
	if !ok {
		// fmt.Println("Error ok!! ,", err)
		log.Fatal("Couldnt get template from template cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)

	//parseTemplates, _ := template.ParseFiles("./Templates/" + html)
	//err = parseTemplates.Execute(w, nil)
	if err != nil {
		fmt.Println("Error Parsing templates, ", err)
		return
	}
}

func CreateRenderTemplatesTest() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./Templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		//fmt.Println("Page is currently ", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./Templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./Templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
