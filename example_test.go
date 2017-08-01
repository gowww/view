package view_test

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"

	"github.com/gowww/view"
)

func Example() {
	// Make the views with global data and functions.
	tmpl := "<h1>{{.app}} — {{.title}}</h1>"
	v := view.New().Data(view.Data{
		"app": "App",
	}).Funcs(view.Funcs{
		"pathescape": url.PathEscape,
		"trim":       strings.Trim,
	}).Parse(tmpl)

	// Get a view with local data.
	w := new(bytes.Buffer)
	v.Execute(w, view.Data{
		"title": "Example",
	})
	fmt.Println(w)
}
