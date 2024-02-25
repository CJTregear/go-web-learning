package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RnderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}
	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	//parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl", "./templates/base.layout.tmpl")
	//err = parsedTemplate.Execute(w, nil)
	//if err != nil {
	//	fmt.Println("error parsing template", err)
	//	return
	//}
}

func createTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	//get all of the files named *.page.tmpl from the ./templates folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.tmpl
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

//var tc = make(map[string]*template.Template)
//
//func RenderTemplate(w http.ResponseWriter, t string) {
//	var tmpl *template.Template
//	var err error
//
//	//check to see if we already have the template in our cache
//	_, inMap := tc[t]
//	if !inMap {
//		err = createTemplateCache(t)
//		log.Println("Creating template and adding to cache")
//		if err != nil {
//			log.Println(err)
//		}
//	} else {
//		log.Println("Using cached template")
//	}
//	tmpl = tc[t]
//
//	err = tmpl.Execute(w, nil)
//	if err != nil {
//		log.Println(err)
//	}

//func createTemplateCache(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
//	}
//
//	//parse the template
//	tmpl, err := template.ParseFiles(templates...)
//	if err != nil {
//		return err
//	}
//	//add template to cache
//	tc[t] = tmpl
//	return nil
//}
