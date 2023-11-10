package static

import (
	"embed"
	"net/http"
)

//go:embed *.css
var files embed.FS

func Handler() http.Handler {
	return http.FileServer(http.FS(files))
}
