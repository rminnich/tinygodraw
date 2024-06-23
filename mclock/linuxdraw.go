//go:build !tinygo

package main

import (
	"image/color"
	"os"
)

type displayer struct {
	f *os.File
}

func new(f *os.File) displayer {
	return displayer{f: f}
}

func (d displayer) Display() error {
	panic("here")
}

func (d displayer) Size() (x, y int16) {
	return 1024, 1024
}
func (d displayer) SetPixel(x, y int16, c color.RGBA) {
	panic("setpixel")
}
