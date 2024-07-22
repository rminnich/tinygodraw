//go:build tinygo

package main

import (
	"image"
	"image/color"

	"tinygo.org/x/drivers"
	"tinygo.org/x/tinydraw/examples/initdisplay"
)

type displayer struct {
	d    drivers.Displayer
	x, y int
}

func new() *displayer {
	d := &displayer{d: initdisplay.InitDisplay()}
	x, y := d.d.Size()
	d.x, d.y = int(x), int(y)
	return d
	for x := 0; x < d.x; x++ {
		for y := 0; y < d.y; y++ {
			d.Set(x, y, red)
		}
	}
	return d

}

func (d *displayer) Display() error {
	//d.d.Display()
	return nil
}

func (d *displayer) Size() (x, y int16) {
	return int16(d.x), int16(d.y)
}

func (d *displayer) At(x, y int) color.Color {
	return blk
}

func (d *displayer) SetPixel(x, y int16, c color.RGBA) {
	d.d.SetPixel(x, y, c)
}

func (d *displayer) Set(x, y int, c color.Color) {
	r, g, b, a := c.RGBA()
	d.d.SetPixel(int16(x), int16(y), color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
}

func (d *displayer) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(d.x), int(d.y))
}

func (d *displayer) ColorModel() color.Model {
	return color.RGBAModel
}
