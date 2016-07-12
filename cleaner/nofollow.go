package cleaner

import (
	"net/url"

	"github.com/BenLubar/htmlcleaner"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func NoFollow(content string, base *url.URL) string {
	nodes := htmlcleaner.ParseDepth(content, 0)
	for _, n := range nodes {
		noFollow(n, base)
	}
	return htmlcleaner.Render(nodes...)
}

func noFollow(n *html.Node, base *url.URL) {
	if n.Type != html.ElementNode {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		noFollow(c, base)
	}

	if n.DataAtom == atom.A {
		for _, a := range n.Attr {
			if a.Key == "href" {
				u, err := base.Parse(a.Val)
				if err != nil || u.Host != base.Host || u.Scheme != base.Scheme {
					break
				}
				return
			}
		}

		for i, a := range n.Attr {
			if a.Key == "rel" {
				n.Attr[i].Val = "nofollow"
				return
			}
		}

		n.Attr = append(n.Attr, html.Attribute{
			Key: "rel",
			Val: "nofollow",
		})
	}
}
