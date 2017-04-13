package main

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/BenLubar/nodebb-plugin-htmlcleaner/cleaner"
	"github.com/gopherjs/gopherjs/js"
)

// zero width non-breaking space is illegal anyway, so removing it won't cause
// any problems as long as another plugin isn't using the same hack.
var fixer = regexp.MustCompile(`^([>\s]*)<`)

func fix(s, uid string) string {
	return fixer.ReplaceAllString(s, "$1\ufeff<")
}

var nconfURL, _ = url.Parse(js.Module.Get("parent").Call("require", "nconf").Call("get", "url").String())
var user = js.Module.Get("parent").Call("require", "./user.js")

func clean(s, uid string) (content string) {
	defer func() {
		rep := make(chan int, 1)
		if id := js.Global.Call("parseInt", uid, 10).Int(); id > 0 {
			user.Call("getUserField", id, "reputation", func(err *js.Error, reputation int) {
				if err != nil {
					reputation = 0
				}
				rep <- reputation
			})
		} else {
			rep <- 0
		}

		if <-rep < 10 {
			content = cleaner.NoFollow(content, nconfURL)
		}
	}()

	return cleaner.Clean(strings.Replace(s, "\ufeff", "", -1))
}

var templateCleaner = strings.NewReplacer("@", "&#64;", "[", "&#91;", "]", "&#93;")

func cleanTemplate(text, _ string) string {
	return templateCleaner.Replace(text)
}
