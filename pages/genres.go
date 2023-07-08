package pages

import "net/http"

func Genres() func(http.ResponseWriter, *http.Request) {
	return infoList("genres")
}
