//go:build ignore
// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

//go:generate go run assets-generator.go

func main() {
	err := vfsgen.Generate(http.Dir("internal/view"), vfsgen.Options{
		Filename:     "internal/webserver/assets-prod.go",
		PackageName:  "webserver",
		BuildTags:    "!dev",
		VariableName: "assets",
	})

	if err != nil {
		log.Fatalln(err)
	}
}
