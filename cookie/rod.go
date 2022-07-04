package main

import "github.com/go-rod/rod"

func maind() {
	browser := rod.New().MustConnect()

	// Even you forget to close, rod will close it after main process ends.
	defer browser.MustClose()

	// Create a new page
	page := browser.MustPage("https://www.kuaishou.com")
	c, _ := page.Cookies(nil)
	for _, v := range c {
		println(v.Name, v.Value)
	}

	cs := browser.MustGetCookies()
	for _, v := range cs {
		println(v.Name, v.Value)
	}
}
