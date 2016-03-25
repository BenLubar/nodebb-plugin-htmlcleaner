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
		var attrs []string
		for a, ok := range attr {
			if !ok {
				continue
			}

			attrs = append(attrs, " "+a.String())
		}
		sort.Strings(attrs)
		list = append(list, el.String()+strings.Join(attrs, ""))
	}
	if len(list) != 0 {
		sort.Strings(list)
		str += "<p>The following elements are allowed: <code>&lt;" + strings.Join(list, "&gt;</code>, <code>&lt;") + "&gt;</code></p>"
	}

	return str
}()

func renderHelp(helpContent string, callback *js.Object) {
	callback.Invoke(nil, helpContent+helpString)
}
