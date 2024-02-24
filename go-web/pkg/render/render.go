package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RnderTemplate renders template using html/template
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl", "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check to see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		err = createTemplateCache(t)
		log.Println("Creating template and adding to cache")
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Using cached template")
	}
	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
	}

	//parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	//add template to cache
	tc[t] = tmpl
	return nil
}
