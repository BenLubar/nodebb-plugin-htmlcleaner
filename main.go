//go:generate gopherjs build

package main

import (
	"github.com/BenLubar/nodebb-plugin-htmlcleaner/cleaner"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	exports := js.Module.Get("exports")

	exports.Set("fix", fix)
	exports.Set("fixPost", async(post(fix)))
	exports.Set("fixSignature", async(signature(fix)))
	exports.Set("fixRaw", async(raw(fix)))

	exports.Set("clean", clean)
	exports.Set("cleanPost", async(post(clean)))
	exports.Set("cleanSignature", async(signature(clean)))
	exports.Set("cleanRaw", async(raw(clean)))

	exports.Set("renderHelp", renderHelp)
}

func async(fn func(data, callback *js.Object)) func(data, callback *js.Object) {
	return func(data, callback *js.Object) {
		go fn(data, callback)
	}
}

func post(fn func(content, uid string) string) func(data, callback *js.Object) {
	return func(data, callback *js.Object) {
		if data != nil && data.Get("postData") != nil && data.Get("postData").Get("content") != nil {
			data.Get("postData").Set("content", fn(data.Get("postData").Get("content").String(), data.Get("postData").Get("uid").String()))
		}
		callback.Invoke(nil, data)
	}
}

func signature(fn func(content, uid string) string) func(data, callback *js.Object) {
	return func(data, callback *js.Object) {
		if data != nil && data.Get("userData") != nil && data.Get("userData").Get("signature") != nil {
			data.Get("userData").Set("signature", fn(data.Get("userData").Get("signature").String(), data.Get("uid").String()))
		}
		callback.Invoke(nil, data)
	}
}

func raw(fn func(content, uid string) string) func(raw, callback *js.Object) {
	return func(raw, callback *js.Object) {
		callback.Invoke(nil, fn(raw.String(), ""))
	}
}

func renderHelp(helpContent string, callback *js.Object) {
	callback.Invoke(nil, helpContent+cleaner.HelpString)
}
