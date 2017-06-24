/*
Package view provides a superset of the standard html/template package.

It allows to keep global data for templates and parse strings, files and directories (also recursively).
*/
package view

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
)

// View contains templates, global data and functions.
type View struct {
	tmpl *template.Template
	data Data
}

// Data is a map of data passed to a view template.
type Data map[string]interface{}

// Funcs is a map of functions usable inside view templates.
type Funcs map[string]interface{}

// New init a view unit.
func New() *View {
	return &View{tmpl: template.New("main")}
}

// Data adds global data for view templates.
func (v *View) Data(data Data) *View {
	if v.data == nil {
		v.data = data
	} else {
		for n, f := range data {
			v.data[n] = f
		}
	}
	return v
}

// Funcs adds functions for view templates.
func (v *View) Funcs(funcs Funcs) *View {
	v.tmpl.Funcs(template.FuncMap(funcs))
	return v
}

func templateMust(t *template.Template, err error) *template.Template {
	if err != nil {
		panic("view: " + err.Error())
	}
	return t
}

// Parse parses text and associates the resulting templates with view.
func (v *View) Parse(text string) *View {
	v.tmpl = templateMust(v.tmpl.Parse(text))
	return v
}

// ParseFiles parses the named files and associates the resulting templates with view.
func (v *View) ParseFiles(filenames ...string) *View {
	v.tmpl = templateMust(v.tmpl.ParseFiles(filenames...))
	return v
}

// ParseGlob parses the files identified by the pattern and associates the resulting templates with view.
func (v *View) ParseGlob(pattern string) *View {
	v.tmpl = templateMust(v.tmpl.ParseGlob(pattern))
	return v
}

// ParseDir parses recursively all files in directory and associates the resulting templates with view.
func (v *View) ParseDir(dirpath string) *View {
	if err := filepath.Walk(dirpath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		_, err = v.tmpl.ParseFiles(path)
		return err
	}); err != nil {
		panic("view: " + err.Error())
	}
	return v
}

// setExecuteData populates referenced data mapped with the view global data.
func (v *View) setExecuteData(data *Data) {
	if v.data != nil {
		if data == nil {
			*data = v.data
		} else {
			for n, f := range v.data {
				(*data)[n] = f
			}
		}
	}
}

// Execute writes w with the main template.
func (v *View) Execute(w io.Writer, data Data) error {
	v.setExecuteData(&data)
	return v.tmpl.Execute(w, data)
}

// ExecuteTemplate writes w with the named template.
func (v *View) ExecuteTemplate(w io.Writer, name string, data Data) error {
	v.setExecuteData(&data)
	return v.tmpl.ExecuteTemplate(w, name, data)
}
