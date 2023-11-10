package pages

import (
	"context"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func isHx(r *http.Request) bool {
	_, is := r.Header[http.CanonicalHeaderKey("HX-Request")]
	return is
}

func render(w http.ResponseWriter, comp templ.Component) {
	w.Header().Add("Content-Type", "text/html")
	err := comp.Render(context.Background(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fail(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	log.Panic(err)
}
