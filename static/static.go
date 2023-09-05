package static

import "embed"

//go:embed *.css
//go:embed *.html
//go:embed *.ico
var Files embed.FS
