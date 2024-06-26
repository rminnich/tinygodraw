// draw pjw on a display. For more info: search pjw images bell labs
package main

//go:generate go run pjw.go

import (
	"image/color"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinydraw/examples/initdisplay"
)

func main() {
	display := initdisplay.InitDisplay()

	tinydraw.FilledRectangle(display, 0, 0, 256, 256, color.RGBA{R: 255, G: 0, B: 128, A: 255})
	for _, p := range pixels {
		tinydraw.FilledRectangle(display, p.x, p.y, 1, 1, color.RGBA{R: p.val, G: p.val, B: p.val, A: 255})
	}
	display.Display()

}
