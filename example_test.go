package view_test

import (
	"bytes"
	"fmt"

	"github.com/gowww/view"
)

func Example() {
	tmpl := "<h1>{{.app}} — {{.title}}</h1>"
	v := view.New().Data(view.Data{"app": "App"}).Funcs(view.AllHelpers).Parse(tmpl)
	w := new(bytes.Buffer)
	v.Execute(w, view.Data{"title": "Example"})
	fmt.Println(w)
}
