package pages

import (
	"database/sql"
	"net/http"
)

var DB *sql.DB

func isHx(r *http.Request) bool {
	_, is := r.Header[http.CanonicalHeaderKey("HX-Request")]
	return is
}
