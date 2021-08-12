package main

import (
	_ "embed"
	"log"
	"sktisrfid/pkg/sys"

	"github.com/wailsapp/wails"
)

func basic() string {
	return "Hello World!"
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	card := &sys.Card{}

	RFID, err := NewRFID()
	if err != nil {
		log.Fatal(err)
	}

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 512,
		Title:  "SKTIS RFID",
		JS:     js,
		CSS:    css,
		Colour: "#ffffff",
	})
	app.Bind(RFID)
	app.Bind(basic)
	app.Bind(card)
	app.Run()
}
