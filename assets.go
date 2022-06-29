package assets

import "embed"

var (
	//go:embed index.html
	Static embed.FS
)
