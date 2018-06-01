package view

import (
	"html/template"
	"strings"
)

// allHelpers are used in all new views.
var allHelpers = Funcs{
	"googlefonts": HelperGoogleFonts,
	"nl2br":       HelperNL2BR,
	"safehtml":    HelperSafeHTML,
	"script":      HelperScript,
	"style":       HelperStyle,
}

// HelperGoogleFonts returns the HTML link to Google Fonts's stylesheet of the given font(s).
func HelperGoogleFonts(fonts ...string) template.HTML {
	return HelperStyle("https://fonts.googleapis.com/css?family=" + strings.Join(fonts, "|"))
}

// HelperNL2BR converts "\n" to HTML "<br>".
// Result will not be unescaped, use HelperSafeHTML for that.
func HelperNL2BR(s string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(s), "\n", "<br>", -1))
}

// HelperSafeHTML prevents s to be escaped. Be careful.
func HelperSafeHTML(s string) template.HTML {
	return template.HTML(s)
}

// HelperScript returns HTML script tags for the given script sources.
func HelperScript(src string) template.HTML {
	return template.HTML(`<script src="` + src + `"></script>`)
}

// HelperStyle returns HTML link tags for the given stylesheets.
func HelperStyle(href string) template.HTML {
	return template.HTML(`<link rel="stylesheet" href="` + href + `">`)
}
