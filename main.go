//go:generate gopherjs build

package main

import (
	"strings"

	"github.com/BenLubar/htmlcleaner"
	"github.com/gopherjs/gopherjs/js"
	"golang.org/x/net/html/atom"
)

func main() {
	exports := js.Module.Get("exports")
	exports.Set("clean", clean)
	exports.Set("cleanPost", cleanPost)
	exports.Set("cleanSignature", cleanSignature)
	exports.Set("cleanRaw", cleanRaw)
}

var config = &htmlcleaner.Config{
	Elem: map[atom.Atom]map[atom.Atom]bool{
		atom.A: {
			atom.Href: true,
		},
		atom.Img: {
			atom.Src: true,
			atom.Alt: true,
		},
		atom.Video: {
			atom.Src:      true,
			atom.Poster:   true,
			atom.Controls: true,
		},
		atom.Audio: {
			atom.Src:      true,
			atom.Controls: true,
		},

		atom.B: nil,
		atom.I: nil,
		atom.U: nil,
		atom.S: nil,

		atom.Em:     nil,
		atom.Strong: nil,
		atom.Strike: nil,

		atom.Big:   nil,
		atom.Small: nil,
		atom.Sup:   nil,
		atom.Sub:   nil,

		atom.Ins: nil,
		atom.Del: nil,

		atom.Abbr:    nil,
		atom.Address: nil,
		atom.Cite:    nil,
		atom.Q:       nil,

		atom.P:          nil,
		atom.Blockquote: nil,

		atom.Pre:  nil,
		atom.Code: nil,
		atom.Kbd:  nil,
		atom.Tt:   nil,

		atom.Details: nil,
		atom.Summary: nil,

		atom.H1: nil,
		atom.H2: nil,
		atom.H3: nil,
		atom.H4: nil,
		atom.H5: nil,
		atom.H6: nil,

		atom.Ul: {
			atom.Start: true,
		},
		atom.Ol: {
			atom.Start: true,
		},
		atom.Li: {
			atom.Value: true,
		},

		atom.Hr: nil,
		atom.Br: nil,

		atom.Div:   nil,
		atom.Table: nil,

		atom.Thead: nil,
		atom.Tbody: nil,
		atom.Tfoot: nil,

		atom.Tr: nil,
		atom.Th: nil,
		atom.Td: nil,

		atom.Caption: nil,
	},

	Attr: map[atom.Atom]bool{
		atom.Title: true,
	},

	AllowJavascriptURL: false,

	EscapeComments: true, // work around https://github.com/psychobunny/templates.js/issues/54
}

func clean(content string) string {
	content = htmlcleaner.Clean(config, content)
	if !strings.HasPrefix(content, "<") {
		content = htmlcleaner.Clean(config, "<p>"+content)
	}
	return content
}

func cleanPost(data, callback *js.Object) {
	if data != nil && data.Get("postData") != nil && data.Get("postData").Get("content") != nil {
		data.Get("postData").Set("content", clean(data.Get("postData").Get("content").String()))
	}
	callback.Invoke(nil, data)
}

func cleanSignature(data, callback *js.Object) {
	if data != nil && data.Get("userData") != nil && data.Get("userData").Get("signature") != nil {
		data.Get("userData").Set("signature", clean(data.Get("userData").Get("signature").String()))
	}
	callback.Invoke(nil, data)
}

func cleanRaw(raw string, callback *js.Object) {
	callback.Invoke(nil, clean(raw))
}
