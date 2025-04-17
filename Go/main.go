package main

import (
	"bytes"
	"embed"
	"encoding/base64"
	"fmt"

	"honnef.co/go/js/dom/v2"
)

//go:embed Windows_Logo_1995.svg
var f embed.FS

const heightFactor = .86

func main() {
	fmt.Println("hello world")
	document, _, ctx := setup()
	spawn := createLogoHandler(document, ctx)
	spawn(0, 0, 100)
	spawn(100, 0, 100)
	spawn(200, 0, 100)
	spawn(0, 100, 100)
}

func setup() (dom.Document, *dom.HTMLBodyElement, *dom.CanvasRenderingContext2D) {
	window := dom.GetWindow()
	document := window.Document()
	body := document.QuerySelector("body").(*dom.HTMLBodyElement)
	canvas := document.CreateElement("canvas").(*dom.HTMLCanvasElement)
	// canvas.SetAttribute("style", "width: 100%; height: 100vh;")
	canvas.SetWidth(window.InnerWidth())
	canvas.SetHeight(window.InnerHeight())
	body.AppendChild(canvas)
	body.SetAttribute("style", "background-color: white; margin: 0; overflow: hidden;")
	return document, body, canvas.GetContext2d()
}

func logo() string {
	var buffer bytes.Buffer
	img, _ := f.ReadFile("Windows_Logo_1995.svg")
	buffer.Write(img)
	svg := buffer.String()
	fmt.Println("Before encoding:")
	fmt.Println(svg)
	fmt.Println("Encoded:")
	b64 := base64.StdEncoding.EncodeToString(img)
	// escaped := strings.ReplaceAll(url.QueryEscape(strings.ReplaceAll(svg, "\n", "")), "+", "%20")
	encoded := fmt.Sprintf("data:image/svg+xml;base64,%v", b64)
	fmt.Println(encoded)
	return encoded
}

func createLogoHandler(document dom.Document, ctx *dom.CanvasRenderingContext2D) func(float64, float64, float64) {
	svg := logo()
	img := document.CreateElement("img").(*dom.HTMLImageElement)
	img.SetAttribute("src", svg)
	return func(dx float64, dy float64, dWidth float64) {
		ctx.DrawImageWithDst(img, dx, dy, dWidth, dWidth*heightFactor)
	}
}
