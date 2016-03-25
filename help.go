package main

import (
	"github.com/BenLubar/nodebb-plugin-htmlcleaner/cleaner"
	"github.com/gopherjs/gopherjs/js"
)

var helpString = func() string {
	buf := []byte("<h2>HTML Cleaner</h2><p>You are allowed to use a subset of HTML.</p>")

	first := true
	for a, ok := range cleaner.Config.Attr {
		if !ok {
			continue
		}

		if first {
			buf = append(buf, "<p>The following attributes are allowed on all elements: "...)
			first = false
		} else {
			buf = append(buf, ", "...)
		}

		buf = append(buf, "<code>"...)
		buf = append(buf, a.String()...)
		buf = append(buf, "</code>"...)
	}
	if !first {
		buf = append(buf, "</p>"...)
	}

	first = true
	for el, attr := range cleaner.Config.Elem {
		if first {
			buf = append(buf, "<p>The following elements are allowed: "...)
			first = false
		} else {
			buf = append(buf, ", "...)
		}

		buf = append(buf, "<code>&lt;"...)
		buf = append(buf, el.String()...)
		for a, ok := range attr {
			if !ok {
				continue
			}

			buf = append(buf, " "...)
			buf = append(buf, a.String()...)
		}
		buf = append(buf, "&gt;</code>"...)
	}
	if !first {
		buf = append(buf, "</p>"...)
	}

	return string(buf)
}()

func renderHelp(helpContent string, callback *js.Object) {
	callback.Invoke(nil, helpContent+helpString)
}
