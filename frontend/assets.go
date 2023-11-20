package frontend

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var assets embed.FS

func Assets() fs.FS {
	assets, err := fs.Sub(assets, "dist")
	if err != nil {
		panic(err)
	}
	return assets
}
