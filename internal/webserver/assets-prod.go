package webserver

import (
	"io/fs"

	"github.com/hulb/shiori/internal"
)

var assets fs.FS

func init() {
	var err error
	assets, err = fs.Sub(internal.Assets, "view")
	if err != nil {
		panic(err)
	}
}
