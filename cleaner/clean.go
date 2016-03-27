package cleaner

import (
	"github.com/BenLubar/htmlcleaner"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func Clean(content string) string {
	nodes := htmlcleaner.Parse(content)

	// convert nodebb-plugin-markdown 5 images to HTML
	dataToSrc(nodes...)

	nodes = htmlcleaner.CleanNodes(Config, nodes)

	// convert it back
	srcToData(nodes...)

	return htmlcleaner.Render(nodes...)
}

const pixel = "data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7"

// https://github.com/julianlam/nodebb-plugin-markdown/commit/3deda2c70bbc9d065e94ec920a4cd79cf29348b9
func dataToSrc(nodes ...*html.Node) {
	for _, n := range nodes {
		if n.Type == html.ElementNode && n.DataAtom == atom.Img {
			attrs := make([]html.Attribute, 0, len(n.Attr))
			hasSrc := false
			hasPixel := false
			for _, a := range n.Attr {
				if a.Namespace == "" && a.Key == "data-src" {
					hasSrc = true
					attrs = append(attrs, html.Attribute{
						Key: "src",
						Val: a.Val,
					})
				} else if a.Namespace == "" && a.Key == "data-state" && a.Val == "unloaded" {
					continue
				} else if a.Namespace == "" && a.Key == "src" && a.Val == pixel {
					hasPixel = true
					continue
				} else {
					attrs = append(attrs, a)
				}
			}
			if hasPixel && !hasSrc {
				attrs = append(attrs, html.Attribute{
					Key: "src",
					Val: pixel,
				})
			}
			n.Attr = attrs
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			dataToSrc(c)
		}
	}
}

func srcToData(nodes ...*html.Node) {
	for _, n := range nodes {
		if n.Type == html.ElementNode && n.DataAtom == atom.Img {
			attrs := make([]html.Attribute, 0, len(n.Attr))
			for _, a := range n.Attr {
				if a.Namespace == "" && a.Key == "src" {
					attrs = append(attrs, html.Attribute{
						Key: "src",
						Val: pixel,
					}, html.Attribute{
						Key: "data-src",
						Val: a.Val,
					}, html.Attribute{
						Key: "data-state",
						Val: "unloaded",
					})
				} else {
					attrs = append(attrs, a)
				}
			}
			n.Attr = attrs
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			srcToData(c)
		}
	}
}
