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
	"googlefonts": HelperGoogleFonts,
}

// HelperGoogleFonts returns an HTML style tags for Google Fonts.
func HelperGoogleFonts(fonts ...string) template.HTML {
	return HelperStyles("https://fonts.googleapis.com/css?family=" + strings.Join(fonts, "|"))
}

// HelperNL2BR converts new line in br HTML tag.
func HelperNL2BR(s string) string {
	return strings.Replace(s, "\n", "<br>", -1)
}

// HelperSafeHTML prevent s to be escaped.
func HelperSafeHTML(s string) template.HTML {
	return template.HTML(s)
}

// HelperScripts returns an HTML script tags for the given src.
func HelperScripts(src ...string) (h template.HTML) {
	for _, script := range src {
		h += template.HTML(`<script src="` + script + `"></script>`)
	}
	return
}

// HelperStyles returns an HTML style tags for the given href.
func HelperStyles(href ...string) (h template.HTML) {
	for _, style := range href {
		h += template.HTML(`<link rel="stylesheet" href="` + style + `">`)
	}
	return
}
