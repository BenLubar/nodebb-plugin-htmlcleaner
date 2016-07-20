package cleaner

import (
	"net/url"
	"regexp"

	"github.com/BenLubar/htmlcleaner"
	"github.com/gopherjs/gopherjs/js"
	"golang.org/x/net/html/atom"
)

var Config = &htmlcleaner.Config{
	Elem: map[atom.Atom]map[atom.Atom]bool{
		atom.A: {
			atom.Href: true,
			atom.Rel:  true,
		},
		atom.Img: {
			atom.Src:    true,
			atom.Alt:    true,
			atom.Class:  true,
			atom.Width:  true,
			atom.Height: true,
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

		atom.Pre: {
			atom.Class: true,
		},
		atom.Code: {
			atom.Class: true,
		},
		atom.Kbd: nil,
		atom.Tt:  nil,

		atom.Details: {
			atom.Open: true,
		},
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

	ValidateURL: func(u *url.URL) (ok bool) {
		defer func() {
			if recover() != nil {
				ok = false
			}
		}()
		js.Global.Call("require", "url").Call("parse", u.String())
		return true
	},

	EscapeComments: true, // work around https://github.com/psychobunny/templates.js/issues/54

	WrapText: true,

	AttrMatch: map[atom.Atom]map[atom.Atom]*regexp.Regexp{
		atom.A: {
			atom.Rel: regexp.MustCompile(`\Anofollow\z`),
		},
		atom.Img: {
			atom.Class: regexp.MustCompile(`\A((emoji|img-markdown|img-responsive)(\s+|\s*\z))*\z`),
		},
		atom.Pre: {
			atom.Class: regexp.MustCompile(`\A((markdown-highlight)(\s+|\s*\z))*\z`),
		},
		atom.Code: {
			atom.Class: regexp.MustCompile(`\A((hljs|language-[a-z0-9]+)(\s+|\s*\z))*\z`),
		},
		atom.Table: {
			atom.Class: regexp.MustCompile(`\A((table|table-bordered|table-striped)(\s+|\s*\z))*\z`),
		},
	},
}
