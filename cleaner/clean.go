package cleaner

import "github.com/BenLubar/htmlcleaner"

func Clean(content string) string {
	return htmlcleaner.Clean(Config, content)
}
