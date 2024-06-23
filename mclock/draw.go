package main

import (
	_ "tinygo.org/x/drivers"
	"os"
)

func main() {
	display := new(os.Stdout)
	redraw(display)
	display.Display()
}
