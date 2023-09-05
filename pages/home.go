package pages

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	if isHx(r) {
		render(w, "Nav", nil)
	} else {
		render(w, "pageHome.go.html", nil)
	}
}
