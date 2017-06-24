package view

import (
	"html/template"
	"strings"
)

// AllHelpers can be passed to view.Funcs to load the view with all helper functions.
var AllHelpers = Funcs{
	"safehtml":    HelperSafeHTML,
	"nl2br":       HelperNL2BR,
	"styles":      HelperStyles,
	"scripts":     HelperScripts,
	"googlefonts": HelperGoogleFontsfunc,
}

func HelperGoogleFontsfunc(fonts ...string) template.HTML {
	return HelperStyles("https://fonts.googleapis.com/css?family=" + strings.Join(fonts, "|"))
}

func HelperNL2BR(s string) string {
	return strings.Replace(s, "\n", "<br>", -1)
}

func HelperSafeHTML(s string) template.HTML {
	return template.HTML(s)
}

func HelperScripts(scripts ...string) (h template.HTML) {
	for _, script := range scripts {
		h += template.HTML(`<script src="` + script + `"></script>`)
	}
	return
}

func HelperStyles(styles ...string) (h template.HTML) {
	for _, style := range styles {
		h += template.HTML(`<link rel="stylesheet" href="` + style + `">`)
	}
	return
}
