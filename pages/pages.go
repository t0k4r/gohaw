package pages

import (
	"html/template"
	"log"
	"net/http"
)

var Templ *template.Template

func isHx(r *http.Request) bool {
	_, is := r.Header[http.CanonicalHeaderKey("HX-Request")]
	return is
}

func render(w http.ResponseWriter, name string, data any) {
	w.Header().Add("Content-Type", "text/html")
	err := Templ.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fail(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	log.Panic(err)
}
