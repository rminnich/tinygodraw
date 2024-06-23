//go:build !tinygo

package main

import (
	"image/color"
	"log"
	"os"
)

type displayer struct {
	f *os.File
}

func new() displayer {
	return displayer{f: os.Stdout}
}

func (d displayer) Display() error {
	log.Printf("Display")
	return nil
}

func (d displayer) Size() (x, y int16) {
	return 1024, 1024
}
func (d displayer) SetPixel(x, y int16, c color.RGBA) {
	log.Printf("Set (%d, %d) to %v", x, y, c)
}
