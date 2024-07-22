//go:build !tinygo

package main

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
)

type displayer struct {
	f    *os.File
	x, y int
}

var _ draw.Image = &displayer{}

func new() *displayer {
	d := displayer{}
	x, y := 320, 240
	d.x, d.y = int(x), int(y)
	return &d
}

func (d *displayer) Display() error {
	log.Printf("Display")
	return nil
}

func (d *displayer) Size() (x, y int16) {
	return int16(d.x), int16(d.y)
}

func (d *displayer) At(x, y int) color.Color {
	return blk
}

func (d *displayer) SetPixel(x, y int16, c color.RGBA) {
	log.Printf("Set (%d, %d) to %v", x, y, c)
}

func (d *displayer) Set(x, y int, c color.Color) {
	log.Printf("Set (%d, %d) to %v", x, y, c)
}

func (d *displayer) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(d.x), int(d.y))
}

func (d *displayer) ColorModel() color.Model {
	return color.RGBAModel
}
