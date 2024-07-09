package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"runtime"

	"github.com/llgcode/draw2d/draw2dimg"
)

type screen struct {
	x, y float64
	*draw2dimg.GraphicContext
}

type mouse struct {
	fill    color.RGBA
	pen     color.RGBA
	points  []float64
	noclose bool
}

var (
	blk   = color.RGBA{A: 255}
	blu   = color.RGBA{B: 255, A: 255}
	grn   = color.RGBA{G: 255, A: 255}
	yel   = color.RGBA{G: 255, R: 255, A: 255}
	red   = color.RGBA{R: 255, A: 255}
	pnk   = color.RGBA{R: 255, G: 128, B: 128, A: 255}
	face  = color.RGBA{R: 128, G: 0, B: 128, A: 255}
	wht   = color.RGBA{G: 255, B: 255, R: 255, A: 255}
	dots  = color.RGBA{G: 64, B: 64, R: 64, A: 255}
	org   = color.RGBA{G: 64, B: 0, R: 64, A: 255}
	clr   = color.RGBA{G: 0, B: 0, R: 0, A: 0}
	flesh = color.RGBA{G: 0, B: 64, R: 64, A: 255}
	/* hair is head[0..41*2], face is head[27*2..56*2] */
	head = mouse{
		fill: wht, pen: blk,
		points: []float64{286, 386, 263, 410, 243, 417, 230, 415, 234, 426, 227, 443, 210, 450, 190, 448,
			172, 435, 168, 418, 175, 400, 190, 398, 201, 400, 188, 390, 180, 375, 178, 363,
			172, 383, 157, 390, 143, 388, 130, 370, 125, 350, 130, 330, 140, 318, 154, 318,
			165, 325, 176, 341, 182, 320, 195, 305, 200, 317, 212, 322, 224, 319, 218, 334,
			217, 350, 221, 370, 232, 382, 250, 389, 264, 387, 271, 380, 275, 372, 276, 381,
			279, 388, 286, 386, 300, 360, 297, 337, 294, 327, 284, 320, 300, 301, 297, 297,
			282, 286, 267, 284, 257, 287, 254, 280, 249, 273, 236, 274, 225, 290, 195, 305},
	}

	mouth = mouse{
		fill: red, pen: blk,
		points: []float64{235, 305, 233, 297, 235, 285, 243, 280, 250, 282, 252, 288, 248, 290, 235, 305}}
	mouth1 = mouse{
		fill: red, pen: blk,
		points: []float64{240, 310, 235, 305, 226, 306}}
	mouth2 = mouse{
		fill: red, pen: blk,
		points: []float64{257, 287, 248, 290}}
	tongue = mouse{
		fill: red, pen: blk,
		points: []float64{235, 285, 243, 280, 246, 281, 247, 286, 245, 289, 241, 291, 237, 294, 233, 294,
			235, 285}}
	tongue1 = mouse{
		fill: red, pen: blk,
		points: []float64{241, 291, 241, 286}}
	shirt = mouse{
		fill: pnk, pen: blk,
		points: []float64{200, 302, 192, 280, 176, 256, 170, 247, 186, 230, 210, 222, 225, 226, 237, 235,
			222, 291, 200, 302}}
	pants = mouse{
		fill: yel, pen: blk,
		points: []float64{199, 164, 203, 159, 202, 143, 189, 138, 172, 135, 160, 137, 160, 166, 151, 170,
			145, 180, 142, 200, 156, 230, 170, 247, 186, 230, 210, 222, 225, 226, 237, 235,
			245, 205, 242, 190, 236, 176, 229, 182, 243, 153, 240, 150, 228, 142, 217, 145,
			212, 162, 199, 164}}
	eyel = mouse{
		fill: wht, pen: blk,
		points: []float64{294, 327, 296, 335, 293, 345, 285, 345, 280, 337, 281, 325, 284, 320, 294, 327}}
	eyer = mouse{
		fill: wht, pen: blk,
		points: []float64{275, 320, 278, 337, 275, 345, 268, 344, 260, 333, 260, 323, 264, 316, 275, 320}}
	pupill = mouse{
		fill: blk, pen: blk,
		points: []float64{284, 320, 294, 327, 293, 329, 291, 333, 289, 333, 286, 331, 284, 325, 284, 320}}
	pupilr = mouse{
		fill: blk, pen: blk,
		points: []float64{265, 316, 275, 320, 275, 325, 273, 330, 271, 332, 269, 333, 267, 331, 265, 327,
			265, 316}}
	nose = mouse{
		fill: blk, pen: blk,
		points: []float64{285, 308, 288, 302, 294, 301, 298, 301, 302, 303, 305, 305, 308, 308, 309, 310,
			310, 312, 310, 316, 308, 320, 305, 323, 302, 324, 297, 324, 294, 322, 288, 317,
			286, 312, 285, 308}}
	nose1 = mouse{
		fill: blk, pen: blk,
		points: []float64{275, 313, 280, 317, 286, 319}}
	buttonl = mouse{
		fill: red, pen: blk,
		points: []float64{201, 210, 194, 208, 190, 196, 191, 187, 199, 188, 208, 200, 201, 210}}
	buttonr = mouse{
		fill: red, pen: blk,
		points: []float64{224, 213, 221, 209, 221, 197, 228, 191, 232, 200, 230, 211, 224, 213}}
	tail = mouse{
		fill: clr, pen: blk,
		noclose: true,
		points:  []float64{40, 80, 50, 76, 66, 79, 90, 102, 106, 151, 128, 173, 145, 180}}
	cuffl = mouse{
		fill: wht, pen: blk,
		noclose: true,
		points:  []float64{202, 143, 197, 148, 188, 150, 160, 137}}
	cuffr = mouse{
		fill: wht, pen: blk,
		noclose: true,
		points:  []float64{243, 153, 233, 154, 217, 145}}
	legl = mouse{
		fill: wht, pen: blk,
		points: []float64{239, 153, 244, 134, 243, 96, 229, 98, 231, 130, 226, 150, 233, 154, 239, 153}}
	legr = mouse{
		fill: wht, pen: blk,
		points: []float64{188, 150, 187, 122, 182, 92, 168, 91, 172, 122, 173, 143, 188, 150}}
	shoel = mouse{
		fill: grn, pen: blk,
		points: []float64{230, 109, 223, 107, 223, 98, 228, 90, 231, 76, 252, 70, 278, 73, 288, 82,
			284, 97, 271, 99, 251, 100, 244, 106, 230, 109}}
	shoel1 = mouse{
		fill: grn, pen: blk,
		points: []float64{223, 98, 229, 98, 243, 96, 251, 100}}
	shoel2 = mouse{
		fill: grn, pen: blk,
		points: []float64{271, 99, 248, 89}}
	shoer = mouse{
		fill: red, pen: blk,
		points: []float64{170, 102, 160, 100, 160, 92, 163, 85, 157, 82, 160, 73, 178, 66, 215, 63,
			231, 76, 228, 90, 213, 97, 195, 93, 186, 93, 187, 100, 184, 102, 170, 102}}
	shoer1 = mouse{
		fill: red, pen: blk,
		points: []float64{160, 92, 168, 91, 182, 92, 186, 93}}
	shoer2 = mouse{
		fill: red, pen: blk,
		points: []float64{195, 93, 182, 83}}
	tick1 = mouse{
		fill: blu, pen: blu,
		points: []float64{302, 432, 310, 446}}
	tick2 = mouse{
		fill: blu, pen: blu,
		points: []float64{370, 365, 384, 371}}
	tick3 = mouse{
		fill: blu, pen: blu,
		points: []float64{395, 270, 410, 270}}
	tick4 = mouse{
		fill: blu, pen: blu,
		points: []float64{370, 180, 384, 173}}
	tick5 = mouse{
		fill: blu, pen: blu,
		points: []float64{302, 113, 310, 100}}
	tick7 = mouse{
		fill: blu, pen: blu,
		points: []float64{119, 113, 110, 100}}
	tick8 = mouse{
		fill: blu, pen: blu,
		points: []float64{40, 173, 52, 180}}
	tick9 = mouse{
		fill: blu, pen: blu,
		points: []float64{10, 270, 25, 270}}
	tick10 = mouse{
		fill: blu, pen: blu,
		points: []float64{40, 371, 52, 365}}
	tick11 = mouse{
		fill: blu, pen: blu,
		points: []float64{110, 446, 119, 432}}
	tick12 = mouse{
		fill: blu, pen: blu,
		points: []float64{210, 455, 210, 470}}
	armh = mouse{
		fill: wht, pen: blk,
		points: []float64{-8, 0, 9, 30, 10, 70, 8, 100, 20, 101, 23, 80, 22, 30, 4, -5}}
	armm = mouse{
		fill: wht, pen: blk,
		points: []float64{-8, 0, 10, 80, 8, 130, 22, 134, 25, 80, 4, -5}}
	handm = mouse{
		fill: blu, pen: blk,
		points: []float64{8, 140, 5, 129, 8, 130, 22, 134, 30, 137, 27, 143, 33, 163, 30, 168,
			21, 166, 18, 170, 12, 168, 10, 170, 5, 167, 4, 195, -4, 195, -6, 170,
			0, 154, 8, 140}}
	handm1 = mouse{
		fill: blu, pen: blk,
		points: []float64{0, 154, 5, 167}}
	handm2 = mouse{
		fill: blu, pen: blk,
		points: []float64{14, 167, 12, 158, 10, 152}}
	handm3 = mouse{
		fill: blu, pen: blk,
		points: []float64{12, 158, 18, 152, 21, 166}}
	handm4 = mouse{
		fill: blu, pen: blk,
		points: []float64{20, 156, 29, 151}}
	handh = mouse{
		fill: red, pen: blk,
		points: []float64{20, 130, 15, 135, 6, 129, 4, 155, -4, 155, -6, 127, -8, 121, 4, 108,
			3, 100, 8, 100, 20, 101, 23, 102, 21, 108, 28, 126, 24, 132, 20, 130}}
	handh1 = mouse{
		fill: red, pen: blk,
		points: []float64{20, 130, 16, 118}}

	showtime = []mouse{
		handh1,
		handh,
		handm4,
		handm3,
		handm2,
		handm1,
		handm,
		armm,
		armh,
	}
	all = []mouse{
		tick12,
		tick11,
		tick10,
		tick9,
		tick8,
		tick7,
		tick5,
		tick4,
		tick3,
		tick2,
		tick1,
		shoer2,
		shoer1,
		shoer,
		shoel2,
		shoel1,
		shoel,
		tongue,
		tongue1,
		pants,
		shirt,
		eyer,
		eyel,
		pupilr,
		pupill,
		nose1,
		nose,
		cuffr,
		cuffl,
		tail,
		mouth2,
		mouth1,
		mouth,
		buttonr,
		buttonl,
		legr,
		legl,
	}
)

func armpoints(gc screen, p []float64, angle float64) []float64 {

	out := make([]float64, len(p))
	for i := range out[:len(out)-1] {
		cosp := math.Cos(math.Pi * angle / 180.0)
		sinp := math.Sin(math.Pi * angle / 180.0)
		out[i] = (float64(p[i])*cosp + float64(p[i+1])*sinp + gc.x/4)
		out[i+1] = (float64(p[i+1])*cosp - float64(p[i])*sinp + gc.y/4)
	}
	return out
}

func poly(gc screen, m mouse) {
	s := m.points
	log.Printf("poly %d points", len(s))
	gc.MoveTo(float64(s[0]), float64(s[1]))
	for i := 2; i < len(s); i += 2 {
		gc.LineTo(float64(s[i]), float64(s[i+1]))
	}
	if !m.noclose {
		gc.Close()
	}
	gc.SetLineWidth(1)
	gc.SetFillColor(m.fill)
	gc.SetStrokeColor(m.pen)
	gc.FillStroke()
}

func arm(s screen, m mouse, slice int, color color.RGBA, angle float64) {
	log.Printf("arm: input: %v", m.points[:slice])
	m.points = armpoints(s, m.points[:slice], angle)
	log.Printf("arm: output: %v", m.points[:slice])

	poly(s, m)
}

func arms(display screen, anghr, angmin float64) {
	/* arms */
	arm(display, armh, 8, blk, anghr)

	/* hour hand */
	arm(display, handh, 16, wht, anghr)
	arm(display, handh, 16, blk, anghr)
	arm(display, handh1, 2, blk, anghr)

	/* minute hand */
	arm(display, armm, 6, blk, angmin)
	arm(display, handm, 2*18, wht, angmin)
	arm(display, handm, 2*18, blk, angmin)
	return
	arm(display, handm1, 2*2, blk, angmin)
	arm(display, handm2, 2*3, blk, angmin)
	arm(display, handm3, 2*3, blk, angmin)
	arm(display, handm4, 2*2, blk, angmin)
}

func main() {
	display := new()

	ix, iy := display.Size()
	flash := func(c color.RGBA) {
		for x := int16(0); x < ix; x++ {
			for y := int16(0); y < iy; y++ {
				display.SetPixel(x, y, c)
			}
		}
	}
	flash(red)
	display.Display()
	x, y := float64(ix), float64(iy)
	flash(blu)
	// Initialize the graphic context on an RGBA image
	r := image.Rect(0, 0, int(ix), int(iy))
	log.Printf("-------------> %d", r)
	return
	flash(yel)
	dest := image.NewAlpha(r)
	flash(red)
	gc := screen{GraphicContext: draw2dimg.NewGraphicContext(dest), x: x, y: y}

	flash(grn)
	canvas := mouse{fill: wht, pen: wht, points: []float64{0, 0, y, 0, y, x, 0, x}}

	/* hair is head[0..41*2], face is head[27*2..56*2] */
	hair := mouse{fill: blk, pen: blk, points: head.points[:41*2]}
	face := mouse{fill: face, pen: blk, points: head.points[27*2:]}
	for i, m := range append([]mouse{canvas, hair, face}, all...) {
		m := m
		s := m.points
		if false {
			log.Printf("%d: pen %v, fill %v", i, m.pen, m.fill)
		}
		gc.MoveTo(float64(s[0]), float64(s[1]))
		for i := 2; i < len(s); i += 2 {
			gc.LineTo(float64(s[i]), float64(s[i+1]))
		}
		if !m.noclose {
			gc.Close()
		}
		gc.SetLineWidth(1)
		gc.SetFillColor(m.fill)
		gc.SetStrokeColor(m.pen)
		gc.FillStroke()
	}
	//	arms(gc, 30, 60)

	for x := int16(0); x < ix; x++ {
		for y := int16(0); y < iy; y++ {
			r, g, b, a := dest.RGBA64At(int(x), int(y)).RGBA()
			c := color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
			display.SetPixel(x, y, c)
		}
	}
	display.Display()
	// Save to file
	if runtime.GOOS == "linux" {
		draw2dimg.SaveToPngFile("hello.png", dest)
	}
}
