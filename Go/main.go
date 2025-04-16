package main

import (
	"fmt"

	"honnef.co/go/js/dom/v2"
)

func main() {
	fmt.Println("hello world")
	document := dom.GetWindow().Document()
	body := document.QuerySelector("body")
	h1 := document.CreateElement("h1")
	h1.SetInnerHTML("hello world")
	body.AppendChild(h1)
	body.SetAttribute("style", "background-color: red;")
}
