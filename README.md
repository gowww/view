# [![gowww](https://avatars.githubusercontent.com/u/18078923?s=20)](https://github.com/gowww) view [![GoDoc](https://godoc.org/github.com/gowww/view?status.svg)](https://godoc.org/github.com/gowww/view) [![Build](https://travis-ci.org/gowww/view.svg?branch=master)](https://travis-ci.org/gowww/view) [![Coverage](https://coveralls.io/repos/github/gowww/view/badge.svg?branch=master)](https://coveralls.io/github/gowww/view?branch=master) [![Go Report](https://goreportcard.com/badge/github.com/gowww/view)](https://goreportcard.com/report/github.com/gowww/view)

Package [view](https://godoc.org/github.com/gowww/view) provides a superset of the standard html/template package.  
It allows to keep global data for templates and parse strings, files and directories (also recursively).

## Installing

1. Get package:

	```Shell
	go get -u github.com/gowww/view
	````

2. Import it in your code:

	```Go
	import "github.com/gowww/view"
	```

## Usage

### Parsing

Use [New](https://godoc.org/github.com/gowww/view#New) to make a new view unit:

```Go
v := view.New()
```

Use [Data](https://godoc.org/github.com/gowww/view#Data) to add global data for view templates.
It can be called multiple times:

```Go
v.Data(Data{"foo": "Foo"})
v.Data(Data{"bar": "Bar"})
```

Use [Funcs](https://godoc.org/github.com/gowww/view#Funcs) to add functions for view templates.
It can be called multiple times:

```Go
v.Funcs(view.AllHelpers)
v.Funcs(Funcs{"pathescape": url.PathEscape,})
```

Use [Parse](https://godoc.org/github.com/gowww/view#Parse) to parse text directly:

```Go
v.Parse("<h1>{{.app}} — {{.title}}</h1><p>{{.foobar}}</p>")
```

Use [ParseFiles](https://godoc.org/github.com/gowww/view#ParseFiles) to parse names files:
```Go
v.ParseFiles("main.gohtml", "admin.gohtml")
```

Use [ParseGlob](https://godoc.org/github.com/gowww/view#ParseGlob) to parse files matching a path pattern:
```Go
v.ParseGlob("views/*.gohtml")
```

Use [ParseDir](https://godoc.org/github.com/gowww/view#ParseDir) to recursively parse all files from a directory:
```Go
v.ParseDir("views")
```

All this can be chained and called multiple times:

```Go
tmpl := "<h1>{{.app}} — {{.title}}</h1><p>{{.foobar}}</p>"
data1 := Data{"app": "App"}
data2 := Data{"foobar": "Foobar"}
v := New().Data(data1).Data(data2).Funcs(view.AllHelpers).Parse(tmpl)
```

### Execution

Use [Execute](https://godoc.org/github.com/gowww/view#Execute) to execute main template:

```Go
w := new(bytes.Buffer)
v.Execute(w, Data{"title": "Example"})
```

Use [ExecuteTemplate](https://godoc.org/github.com/gowww/view#ExecuteTemplate) to execute a named template:

```Go
w := new(bytes.Buffer)
v.ExecuteTemplate(w, "home", Data{"title": "Example"})
```