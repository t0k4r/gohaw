package pages

import "net/http"

func Themes() func(http.ResponseWriter, *http.Request) {
	return infoList(getTemplInfo("themes"))
}
