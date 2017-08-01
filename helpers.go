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
func HelperScripts(srcs ...string) (h template.HTML) {
	for _, src := range srcs {
		h += template.HTML(`<script src="` + src + `"></script>`)
	}
	return
}

// HelperStyles returns HTML link tags for the given stylesheets.
func HelperStyles(hrefs ...string) (h template.HTML) {
	for _, href := range hrefs {
		h += template.HTML(`<link rel="stylesheet" href="` + href + `">`)
	}
	return
}
