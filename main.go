//go:generate gopherjs build

package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	exports := js.Module.Get("exports")

	exports.Set("fix", fix)
	exports.Set("fixPost", post(fix))
	exports.Set("fixSignature", signature(fix))
	exports.Set("fixRaw", raw(fix))

	exports.Set("clean", clean)
	exports.Set("cleanPost", post(clean))
	exports.Set("cleanSignature", signature(clean))
	exports.Set("cleanRaw", raw(clean))

	exports.Set("renderHelp", renderHelp)
}

func post(fn func(string) string) func(data, callback *js.Object) {
	return func(data, callback *js.Object) {
		if data != nil && data.Get("postData") != nil && data.Get("postData").Get("content") != nil {
			data.Get("postData").Set("content", fn(data.Get("postData").Get("content").String()))
		}
		callback.Invoke(nil, data)
	}
}

func signature(fn func(string) string) func(data, callback *js.Object) {
	return func(data, callback *js.Object) {
		if data != nil && data.Get("userData") != nil && data.Get("userData").Get("signature") != nil {
			data.Get("userData").Set("signature", fn(data.Get("userData").Get("signature").String()))
		}
		callback.Invoke(nil, data)
	}
}

func raw(fn func(string) string) func(raw string, callback *js.Object) {
	return func(raw string, callback *js.Object) {
		callback.Invoke(nil, fn(raw))
	}
}
