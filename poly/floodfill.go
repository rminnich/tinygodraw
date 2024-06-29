package main

import (
	"image"
	"image/color"
	"log"
)

var d = log.Printf

func floodFill(img *image.Alpha, start image.Point, col color.Alpha) {

	var painting bool
	r := img.Bounds()
	for y := r.Min.Y; y < r.Max.Y; y++ {
		painting = false
		for x := r.Min.X; x < r.Max.X; x++ {
			at := img.AlphaAt(x, y).A
			d("%d,%d=%v col %v col.A %v", y, x, at, col, col.A)
			if !painting {
				if at != col.A {
					d("skip %v,%v", y, x)
					continue
				}
				d("start %v,%v", y, x)
				painting = true
				continue
			} else {
				if at == col.A {
					d("stop %v,%v", y, x)
					painting = false
					continue
				}
			}
			d("paint %v,%v", y, x)
			img.SetAlpha(x, y, col)
		}
	}

}
