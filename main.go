//go:generate env GOOS=js GOARCH=wasm go build -o nodebb-plugin-htmlcleaner.wasm

package main

import (
	"fmt"
	"runtime/debug"
	"syscall/js"

	"github.com/BenLubar/nodebb-plugin-htmlcleaner/cleaner"
)

func main() {
	exports := js.Global().Get("module").Get("exports")

	exports.Set("fixPost", cb(async(post(fix))))
	exports.Set("fixSignature", cb(async(signature(fix))))
	exports.Set("fixRaw", cb(async(raw(fix))))

	exports.Set("cleanPost", cb(async(post(clean))))
	exports.Set("cleanSignature", cb(async(signature(clean))))
	exports.Set("cleanRaw", cb(async(raw(clean))))

	exports.Set("renderHelp", cb(renderHelp))

	exports.Set("templatePost", cb(async(post(cleanTemplate))))
	exports.Set("templateSignature", cb(async(signature(cleanTemplate))))
	exports.Set("templateRaw", cb(async(raw(cleanTemplate))))
	exports.Set("templateTopic", cb(templateTopic))
	exports.Set("templateTopics", cb(templateTopics))

	// wait for callbacks
	select {}
}

func cb(fn func(args []js.Value)) js.Callback {
	return js.NewCallback(func(args []js.Value) {
		defer func() {
			if r := recover(); r != nil {
				stack := debug.Stack()
				winston := js.Global().Get("require").Get("main").Call("require", "winston")
				winston.Call("error", fmt.Sprintf("[nodebb-plugin-htmlcleaner] !!PANIC!! %v\n%s", r, stack))

				callback := args[len(args)-1]
				if callback.Type() == js.TypeFunction {
					callback.Invoke(js.Global().Get("Error").New(fmt.Sprintf("HTMLCleaner panicked. Please send a screenshot of this gibberish to @ben_lubar: %v\n%s", r, stack)))
				}
			}
		}()

		fn(args)
	})
}

func async(fn func(args []js.Value)) func(args []js.Value) {
	return func(args []js.Value) {
		go fn(args)
	}
}

func post(fn func(content, uid string) string) func(args []js.Value) {
	return func(args []js.Value) {
		data, callback := args[0], args[1]
		if data.Type() == js.TypeObject && data.Get("postData").Type() == js.TypeObject && data.Get("postData").Get("content").Type() == js.TypeString {
			data.Get("postData").Set("content", fn(data.Get("postData").Get("content").String(), data.Get("postData").Get("uid").String()))
		}
		callback.Invoke(nil, data)
	}
}

func signature(fn func(content, uid string) string) func(args []js.Value) {
	return func(args []js.Value) {
		data, callback := args[0], args[1]
		if data.Type() == js.TypeObject && data.Get("userData").Type() == js.TypeObject && data.Get("userData").Get("signature").Type() == js.TypeString {
			data.Get("userData").Set("signature", fn(data.Get("userData").Get("signature").String(), data.Get("uid").String()))
		}
		callback.Invoke(nil, data)
	}
}

func raw(fn func(content, uid string) string) func(args []js.Value) {
	return func(args []js.Value) {
		raw, callback := args[0], args[1]
		callback.Invoke(nil, fn(raw.String(), ""))
	}
}

func renderHelp(args []js.Value) {
	helpContent, callback := args[0], args[1]
	callback.Invoke(nil, helpContent.String()+cleaner.HelpString)
}

func templateTopic(args []js.Value) {
	data, callback := args[0], args[1]
	cleanTopic(data.Get("topic"))
	callback.Invoke(nil, data)
}

func templateTopics(args []js.Value) {
	data, callback := args[0], args[1]
	topics := data.Get("topics")
	for i := 0; i < topics.Length(); i++ {
		cleanTopic(topics.Index(i))
	}
	callback.Invoke(nil, data)
}

func cleanTopic(topic js.Value) {
	if topic.Get("tags").Type() == js.TypeObject {
		tags := topic.Get("tags")
		for i := 0; i < tags.Length(); i++ {
			tags.Index(i).Set("value", cleanTemplate(tags.Index(i).Get("value").String(), ""))
		}
	}
	topic.Set("title", cleanTemplate(topic.Get("title").String(), ""))
}
