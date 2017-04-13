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

	exports.Set("template", cleanTemplate)
	exports.Set("templatePost", async(post(cleanTemplate)))
	exports.Set("templateSignature", async(signature(cleanTemplate)))
	exports.Set("templateRaw", async(raw(cleanTemplate)))
	exports.Set("templateTopic", templateTopic)
	exports.Set("templateTopics", templateTopics)
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

func templateTopic(data, callback *js.Object) {
	cleanTopic(data.Get("topic"))
	callback.Invoke(nil, data)
}

func templateTopics(data, callback *js.Object) {
	topics := data.Get("topics")
	for i := 0; i < topics.Length(); i++ {
		cleanTopic(topics.Index(i))
	}
	callback.Invoke(nil, data)
}

func cleanTopic(topic *js.Object) {
	if topic.Get("tags").Bool() {
		tags := topic.Get("tags")
		for i := 0; i < tags.Length(); i++ {
			tags.Index(i).Set("value", cleanTemplate(tags.Index(i).Get("value").String(), ""))
		}
	}
	topic.Set("title", cleanTemplate(topic.Get("title").String(), ""))
}
