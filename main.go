package main

import (
	"fmt"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func helloSir(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func main() {
	js := mewn.String("./frontend/public/build/bundle.js")
	css := mewn.String("./frontend/public/build/bundle.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "Wails Photo Viewer",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(helloSir)
	app.Run()
}
