package cleaner

import (
	"github.com/BenLubar/htmlcleaner"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func Clean(content string) string {
	content = htmlcleaner.Preprocess(Config, content)
	content = htmlcleaner.Clean(Config, content)
	nodes := htmlcleaner.ParseDepth(content, 0)
	for _, n := range nodes {
		ensureControls(n)
	}
	return htmlcleaner.Render(nodes...)
}

func ensureControls(n *html.Node) {
	if n.Type != html.ElementNode {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ensureControls(c)
	}

	if n.DataAtom == atom.Video || n.DataAtom == atom.Audio {
		for _, a := range n.Attr {
			if a.Key == "controls" {
				return
			}
		}

		n.Attr = append(n.Attr, html.Attribute{
			Key: "controls",
		})
	}
}
