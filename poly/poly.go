package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	// Define the set of (x, y) pairs for the polygon vertices
	vertices := []image.Point{
		{50, 50},
		{150, 50},
		{150, 150},
		{50, 150},
	}

	// Create a new RGBA image
	img := image.NewRGBA(image.Rect(0, 0, 200, 200))

	// Fill the entire image with white color
	d("white %v", color.White)
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Fill the polygon with a specific color
	fillPolygon(img, vertices, color.Black)

	// Save the image to a file
	file, err := os.Create("polygon.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		panic(err)
	}
}

// fillPolygon fills the given polygon on the image with the specified color
func fillPolygon(img *image.RGBA, vertices []image.Point, col color.Color) {
	// Create a mask image to draw the polygon on
	mask := image.NewAlpha(img.Bounds())

	// Draw the polygon on the mask
	for i := 0; i < len(vertices); i++ {
		j := (i + 1) % len(vertices)
		drawLine(mask, vertices[i], vertices[j], color.Alpha(col))
	}

	// Fill the polygon using flood fill
	floodFill(mask, vertices[0], color.Alpha{255})

	// Apply the mask to the original image
	draw.DrawMask(img, img.Bounds(), &image.Uniform{col}, image.Point{}, mask, image.Point{}, draw.Over)
}

// drawLine draws a line on the image from p1 to p2 with the specified color
func drawLine(img *image.Alpha, p1, p2 image.Point, col color.Alpha) {
	dx := abs(p2.X - p1.X)
	dy := abs(p2.Y - p1.Y)
	sx := -1
	if p1.X < p2.X {
		sx = 1
	}
	sy := -1
	if p1.Y < p2.Y {
		sy = 1
	}
	err := dx - dy

	for {
		img.SetAlpha(p1.X, p1.Y, col)
		if p1 == p2 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			p1.X += sx
		}
		if e2 < dx {
			err += dx
			p1.Y += sy
		}
	}
}

func contains(img image.Image, p image.Point) bool {
	if p.X < 0 || p.Y < 0 || p.X >= img.Bounds().Dx() || p.Y >= img.Bounds().Dy() {
		return false
	}
	return true
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
