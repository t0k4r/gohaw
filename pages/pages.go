package pages

import (
	"html/template"
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
