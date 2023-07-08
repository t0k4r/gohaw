package pages

import "net/http"

func Studios() func(http.ResponseWriter, *http.Request) {
	return infoList("studios")
}
