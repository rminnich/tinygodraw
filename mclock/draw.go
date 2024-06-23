package main

import (
	_ "tinygo.org/x/drivers"
)

func main() {
	display := new()
	redraw(display)
	display.Display()
}
