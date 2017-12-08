package template

import (
	"html/template"
	"path"

	"github.com/unrolled/render"
)

var (
	// Render is.
	Render = render.New(render.Options{
		Directory:                 path.Join("templates"),
		Asset:                     nil,
		AssetNames:                nil,
		Layout:                    "layout",
		Extensions:                []string{".html"},
		Funcs:                     []template.FuncMap{},
		Delims:                    render.Delims{Left: "{{", Right: "}}"},
		Charset:                   "UTF-8",
		IndentJSON:                false,
		IndentXML:                 false,
		PrefixJSON:                []byte(""),
		PrefixXML:                 []byte(""),
		HTMLContentType:           "text/html",
		IsDevelopment:             false,
		UnEscapeHTML:              false,
		StreamingJSON:             false,
		RequirePartials:           false,
		DisableHTTPErrorRendering: false,
	})
)
