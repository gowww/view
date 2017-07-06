package view

import (
	"html/template"
	"strings"
)

// AllHelpers can be passed to view.Funcs to load the view with all helper functions.
var AllHelpers = Funcs{
	"googlefonts": HelperGoogleFonts,
	"nl2br":       HelperNL2BR,
	"safehtml":    HelperSafeHTML,
	"scripts":     HelperScripts,
	"styles":      HelperStyles,
}

// HelperGoogleFonts returns the HTML link to Google Fonts's stylesheet of the given font(s).
func HelperGoogleFonts(fonts ...string) template.HTML {
	return HelperStyles("https://fonts.googleapis.com/css?family=" + strings.Join(fonts, "|"))
}

// HelperNL2BR converts "\n" to HTML "<br>".
// Result will not be unescaped, use HelperSafeHTML for that.
func HelperNL2BR(s string) string {
	return strings.Replace(s, "\n", "<br>", -1)
}

// HelperSafeHTML prevents s to be escaped. Be careful.
func HelperSafeHTML(s string) template.HTML {
	return template.HTML(s)
}

// HelperScripts returns HTML script tags for the given script sources.
func HelperScripts(src ...string) (h template.HTML) {
	for _, script := range src {
		h += template.HTML(`<script src="` + script + `"></script>`)
	}
	return
}

// HelperStyles returns HTML link tags for the given stylesheets.
func HelperStyles(href ...string) (h template.HTML) {
	for _, style := range href {
		h += template.HTML(`<link rel="stylesheet" href="` + style + `">`)
	}
	return
}
