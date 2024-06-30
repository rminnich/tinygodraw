package main

import (
	"image"
	"image/color"
	"log"
)

var d = log.Printf
// floodFill fills the area connected to the start point with the specified color
func floodFill(img *image.Alpha, start image.Point, col color.Alpha) {
	q := []image.Point{start}
	initialColor := img.AlphaAt(start.X, start.Y)

	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		if p.X < 0 || p.Y < 0 || p.X >= img.Bounds().Dx() || p.Y >= img.Bounds().Dy() {
			continue
		}

		if img.AlphaAt(p.X, p.Y) == initialColor {
			img.SetAlpha(p.X, p.Y, col)

			q = append(q, image.Point{p.X + 1, p.Y})
			q = append(q, image.Point{p.X - 1, p.Y})
			q = append(q, image.Point{p.X, p.Y + 1})
			q = append(q, image.Point{p.X, p.Y - 1})
		}
	}
}

