package main

import (
	"sort"
	"strings"

	"github.com/BenLubar/nodebb-plugin-htmlcleaner/cleaner"
	"github.com/gopherjs/gopherjs/js"
)

var helpString = func() string {
	str := "<h2>HTML Cleaner</h2><p>You are allowed to use a subset of HTML.</p>"

	var list []string
	for a, ok := range cleaner.Config.Attr {
		if !ok {
			continue
		}

		list = append(list, a.String())
	}
	if len(list) != 0 {
		sort.Strings(list)
		str += "<p>The following attributes are allowed on all elements: <code>" + strings.Join(list, "</code>, <code>") + "</code></p>"
	}

	list = nil
	for el, attr := range cleaner.Config.Elem {
		buf := []byte("&lt;")
		buf = append(buf, el.String()...)
		for a, ok := range attr {
			if !ok {
				continue
			}

			buf = append(buf, " "...)
			buf = append(buf, a.String()...)
		}
		buf = append(buf, "&gt;"...)
		list = append(list, string(buf))
	}
	if len(list) != 0 {
		sort.Strings(list)
		str += "<p>The following elements are allowed: <code>" + strings.Join(list, "</code>, <code>") + "</code></p>"
	}

	return str
}()

func renderHelp(helpContent string, callback *js.Object) {
	callback.Invoke(nil, helpContent+helpString)
}
