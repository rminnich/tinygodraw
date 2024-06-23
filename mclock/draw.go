package main

import (
	"os"

	_ "tinygo.org/x/drivers"
)

func main() {
	display := new(os.Stdout)
	redraw(display)
	display.Display()
}
