package templates

import "embed"

//go:embed *.go.html
var Files embed.FS
