package main

import (
	"strings"

	"github.com/BenLubar/nodebb-plugin-htmlcleaner/cleaner"
)

// zero width non-breaking space is illegal anyway, so removing it won't cause
// any problems as long as another plugin isn't using the same hack.
var fixer = strings.NewReplacer("\n<", "\n\ufeff<")
var remover = strings.NewReplacer("\ufeff", "")

func fix(s string) string {
	return fixer.Replace(s)
}

func clean(s string) string {
	return cleaner.Clean(remover.Replace(s))
}
