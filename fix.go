package main

import (
	"net/url"
	"regexp"
	"strings"
	"syscall/js"

	"github.com/BenLubar/nodebb-plugin-htmlcleaner/cleaner"
)

// zero width non-breaking space is illegal anyway, so removing it won't cause
// any problems as long as another plugin isn't using the same hack.
var fixer1 = regexp.MustCompile(`^([>\s]*)<`)
var fixer2 = regexp.MustCompile(`\[([ Xx])\]`)
var fixer3 = regexp.MustCompile(`<!--`)

func fix(s, uid string) string {
	s = fixer1.ReplaceAllString(s, "$1\ufeff<")
	s = fixer2.ReplaceAllString(s, "[\ufeff$1]")
	s = fixer3.ReplaceAllLiteralString(s, "<!--\u200b") // prevent templating system from barfing
	return s
}

var nconfURL, _ = url.Parse(js.Global().Get("require").Get("main").Call("require", "nconf").Call("get", "url").String())
var user = js.Global().Get("require").Get("main").Call("require", "./src/user")

func clean(s, uid string) (content string) {
	defer func() {
		rep := make(chan int, 1)
		if id := js.Global().Call("parseInt", uid, 10).Int(); id > 0 {
			var cb js.Callback
			cb = js.NewCallback(func(args []js.Value) {
				var reputation int
				if len(args) > 1 && (args[0].Type() == js.TypeUndefined || args[0].Type() == js.TypeNull) {
					reputation = args[1].Int()
				}
				rep <- reputation

				cb.Release()
			})

			user.Call("getUserField", id, "reputation", cb)
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
