package assets

import "embed"

var (
	//go:embed bootstrap
	Bootstrap embed.FS

	//go:embed templates
	Templates embed.FS
)
