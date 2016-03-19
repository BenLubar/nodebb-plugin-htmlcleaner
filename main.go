//go:generate gopherjs build

package main

import (
	"github.com/BenLubar/nodebb-plugin-htmlcleaner/cleaner"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	exports := js.Module.Get("exports")
	exports.Set("clean", cleaner.Clean)
	exports.Set("cleanPost", cleanPost)
	exports.Set("cleanSignature", cleanSignature)
	exports.Set("cleanRaw", cleanRaw)
}

func cleanPost(data, callback *js.Object) {
	if data != nil && data.Get("postData") != nil && data.Get("postData").Get("content") != nil {
		data.Get("postData").Set("content", cleaner.Clean(data.Get("postData").Get("content").String()))
	}
	callback.Invoke(nil, data)
}

func cleanSignature(data, callback *js.Object) {
	if data != nil && data.Get("userData") != nil && data.Get("userData").Get("signature") != nil {
		data.Get("userData").Set("signature", cleaner.Clean(data.Get("userData").Get("signature").String()))
	}
	callback.Invoke(nil, data)
}

func cleanRaw(raw string, callback *js.Object) {
	callback.Invoke(nil, cleaner.Clean(raw))
}
