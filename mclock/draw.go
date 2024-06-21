// draw pjw on a display. For more info: search pjw images bell labs
package main

//go:generate go run pjw.go

import (
	_ "tinygo.org/x/drivers"
	"tinygo.org/x/tinydraw/examples/initdisplay"
)

func main() {
	display := initdisplay.InitDisplay()
	redraw(display)
	display.Display()
}
