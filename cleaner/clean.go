package cleaner

import "github.com/BenLubar/htmlcleaner"

func Clean(content string) string {
	content = htmlcleaner.Preprocess(Config, content)
	return htmlcleaner.Clean(Config, content)
}
