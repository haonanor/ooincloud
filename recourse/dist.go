package recourse

import (
	"embed"
	_ "embed"
)

//go:embed dist/assets
var StaticPath embed.FS

//go:embed dist/index.html
var MainIndexPath string
