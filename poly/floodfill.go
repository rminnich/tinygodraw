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
	col.A = 0
	for x := r.Min.X; x < r.Max.X; x++ {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			at := img.AlphaAt(x, y).A
			d("%d,%d:@ %v col %v col.A %v", x, y, at, col, col.A)
			if !painting {
				if at != col.A {
					d("skip %v,%v", x, y)
					continue
				}
				d("start %v,%v", x, y)
				painting = true
			} else {
				if at == col.A {
					d("stop %v,%v", x, y)
					painting = false
					continue
				}
			}
			d("paint %v,%v", x, y)
			img.SetAlpha(x, y, col)
		}
	}

}
