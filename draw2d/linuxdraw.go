//go:build !tinygo

package main

import (
	"image"
	"image/color"
	"log"
	"os"
)

type displayer struct {
	f *os.File
	bs [][]color.Color
	x, y int
}

var _ image.Image = &displayer{}

func new() displayer {
	d := displayer{}
	d.bs = make([][]color.Color, 320)
	for x := 0; x < 320; x++ {
		d.bs[x] = make([]color.Color, 240)
		for y := 0; y < 240; y++ {
			d.bs[x][y] = color.RGBA{}
		}
	}
	log.Printf("d.bs %v", d.bs)
	return d
}

func (d displayer) Display() error {
	log.Printf("Display")
	return nil
}


func (d displayer) Size() (x, y int16) {
	return int16(d.x), int16(d.y)
}

func (d displayer) At(x, y int) color.Color {
	return d.bs[x][y]
}

func (d displayer) SetPixel(x, y int16, c color.RGBA) {
	d.bs[x][y] = color.Color(c)
	//	log.Printf("Set (%d, %d) to %v", x, y, c)
}

func (d displayer) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(d.x), int(d.y))
}

func (d displayer) ColorModel() color.Model {
	return color.RGBAModel
}

