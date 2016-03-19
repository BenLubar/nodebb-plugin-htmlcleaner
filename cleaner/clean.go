package cleaner

import (
	"strings"

	"github.com/BenLubar/htmlcleaner"
)

func Clean(content string) string {
	content = htmlcleaner.Clean(config, content)
	if !strings.HasPrefix(content, "<") {
		content = htmlcleaner.Clean(config, "<p>"+content)
	}
	return content
}
