package pages

import (
	"html/template"
	"log"
	"net/http"
)

func Home() func(http.ResponseWriter, *http.Request) {
	templ, err := template.ParseFiles("templates/Layout.html", "templates/Home.html")
	if err != nil {
		log.Panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html")
		if isHx(r) {
			err := templ.ExecuteTemplate(w, "Main", nil)
			if err != nil {
				log.Panic(err)
			}
		} else {
			err := templ.Execute(w, nil)
			if err != nil {
				log.Panic(err)
			}
		}
	}
}
