package view

import (
	"bytes"
	"fmt"
	"testing"
)

func TestView(t *testing.T) {
	tmpl := "<h1>{{.app}} — {{.title}}</h1><p>{{.foobar}}</p>"
	v := New().Data(Data{"app": "App"}).Data(Data{"foobar": "Foobar"}).Funcs(AllHelpers).Parse(tmpl)
	w := new(bytes.Buffer)
	v.Execute(w, Data{"title": "Example"})
	fmt.Println(w)
}
