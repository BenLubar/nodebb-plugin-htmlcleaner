package cleaner

import (
	"net/url"
	"regexp"
	"syscall/js"

	"github.com/BenLubar/htmlcleaner"
	"golang.org/x/net/html/atom"
)

var Config = (&htmlcleaner.Config{
	ValidateURL: func(u *url.URL) (ok bool) {
		if !htmlcleaner.SafeURLScheme(u) {
			return false
		}

		defer func() {
			if recover() != nil {
				ok = false
			}
		}()
		js.Global().Call("require", "url").Call("parse", u.String())
		return true
	},

	WrapText: false, // https://what.thedailywtf.com/post/1049805
}).
	GlobalAttrAtom(atom.Title).
	ElemAttrAtom(atom.A, atom.Href).
	ElemAttrAtomMatch(atom.A, atom.Rel, regexp.MustCompile(`\Anofollow\z`)).
	ElemAttrAtom(atom.Img, atom.Src, atom.Alt, atom.Width, atom.Height).
	ElemAttrAtomMatch(atom.Img, atom.Class, regexp.MustCompile(`\A((emoji|img-markdown|img-responsive|2ac3a3adc5551647133578b5d9c21ea4)(\s+|\s*\z))*\z`)).
	ElemAttrAtom(atom.Video, atom.Src, atom.Poster, atom.Controls).
	ElemAttrAtom(atom.Audio, atom.Src, atom.Controls).
	ElemAttrAtom(atom.Source, atom.Src, atom.Type).
	ElemAtom(atom.B, atom.I, atom.U, atom.S).
	ElemAtom(atom.Em, atom.Strong, atom.Strike).
	ElemAtom(atom.Big, atom.Small, atom.Sup, atom.Sub).
	ElemAtom(atom.Ins, atom.Del).
	ElemAtom(atom.Abbr, atom.Address, atom.Cite, atom.Q).
	ElemAtom(atom.P, atom.Blockquote).
	ElemAttrAtomMatch(atom.Pre, atom.Class, regexp.MustCompile(`\A((markdown-highlight)(\s+|\s*\z))*\z`)).
	ElemAttrAtomMatch(atom.Code, atom.Class, regexp.MustCompile(`\A((hljs|language-[a-z0-9]+)(\s+|\s*\z))*\z`)).
	ElemAtom(atom.Kbd, atom.Tt).
	ElemAttrAtom(atom.Details, atom.Open).
	ElemAtom(atom.Summary).
	ElemAtom(atom.H1, atom.H2, atom.H3, atom.H4, atom.H5, atom.H6).
	ElemAttrAtom(atom.Ul, atom.Start).
	ElemAttrAtom(atom.Ol, atom.Start).
	ElemAttrAtom(atom.Li, atom.Value).
	ElemAtom(atom.Hr, atom.Br).
	ElemAtom(atom.Div, atom.Span).
	ElemAttrAtomMatch(atom.Table, atom.Class, regexp.MustCompile(`\A((table|table-bordered|table-striped)(\s+|\s*\z))*\z`)).
	ElemAtom(atom.Thead, atom.Tbody, atom.Tfoot).
	ElemAtom(atom.Tr, atom.Th, atom.Td).
	ElemAtom(atom.Caption).
	ElemAtom(atom.Dl, atom.Dt, atom.Dd)
