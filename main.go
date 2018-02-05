package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
)

func main() {

	m := martini.Classic()

	m.Use(gzip.All(gzip.Options{
		CompressionLevel: gzip.BestCompression,
	}))

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Layout:     "layouts/layout",           // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		// Funcs:      []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
		Delims:     render.Delims{Left: "[[", Right: "]]"}, // Sets delimiters to the specified strings.
		Charset:    "UTF-8",                                // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                                   // Output human readable JSON
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(http.StatusOK, "index", nil)
	})
	m.Run()
}
