package main

import (
	"image"
	"image/color"
	"log"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/samples/gopher2"
)

type mouse struct {
	col    color.RGBA
	points []float64
}

var (
	blk   = color.RGBA{A: 255}
	blu   = color.RGBA{B: 255, A: 255}
	red   = color.RGBA{R: 255, A: 255}
	wht   = color.RGBA{G: 255, B: 255, R: 255, A: 255}
	dots  = color.RGBA{G: 64, B: 64, R: 64, A: 255}
	org   = color.RGBA{G: 64, B: 0, R: 64, A: 255}
	flesh = color.RGBA{G: 0, B: 64, R: 64, A: 255}
	/* hair is head[0..41*2], face is head[27*2..56*2] */
	head = mouse{
		col: wht,
		points: []float64{286, 386, 263, 410, 243, 417, 230, 415, 234, 426, 227, 443, 210, 450, 190, 448,
			172, 435, 168, 418, 175, 400, 190, 398, 201, 400, 188, 390, 180, 375, 178, 363,
			172, 383, 157, 390, 143, 388, 130, 370, 125, 350, 130, 330, 140, 318, 154, 318,
			165, 325, 176, 341, 182, 320, 195, 305, 200, 317, 212, 322, 224, 319, 218, 334,
			217, 350, 221, 370, 232, 382, 250, 389, 264, 387, 271, 380, 275, 372, 276, 381,
			279, 388, 286, 386, 300, 360, 297, 337, 294, 327, 284, 320, 300, 301, 297, 297,
			282, 286, 267, 284, 257, 287, 254, 280, 249, 273, 236, 274, 225, 290, 195, 305},
	}

	mouth = mouse{
		col:    red,
		points: []float64{235, 305, 233, 297, 235, 285, 243, 280, 250, 282, 252, 288, 248, 290, 235, 305}}
	mouth1 = mouse{
		col:    red,
		points: []float64{240, 310, 235, 305, 226, 306}}
	mouth2 = mouse{
		col:    red,
		points: []float64{257, 287, 248, 290}}
	tongue = mouse{
		col: flesh,
		points: []float64{235, 285, 243, 280, 246, 281, 247, 286, 245, 289, 241, 291, 237, 294, 233, 294,
			235, 285}}
	tongue1 = mouse{
		col:    wht,
		points: []float64{241, 291, 241, 286}}
	shirt = mouse{
		col: wht,
		points: []float64{200, 302, 192, 280, 176, 256, 170, 247, 186, 230, 210, 222, 225, 226, 237, 235,
			222, 291, 200, 302}}
	pants = mouse{
		col: wht,
		points: []float64{199, 164, 203, 159, 202, 143, 189, 138, 172, 135, 160, 137, 160, 166, 151, 170,
			145, 180, 142, 200, 156, 230, 170, 247, 186, 230, 210, 222, 225, 226, 237, 235,
			245, 205, 242, 190, 236, 176, 229, 182, 243, 153, 240, 150, 228, 142, 217, 145,
			212, 162, 199, 164}}
	eyel = mouse{
		col:    wht,
		points: []float64{294, 327, 296, 335, 293, 345, 285, 345, 280, 337, 281, 325, 284, 320, 294, 327}}
	eyer = mouse{
		col:    wht,
		points: []float64{275, 320, 278, 337, 275, 345, 268, 344, 260, 333, 260, 323, 264, 316, 275, 320}}
	pupill = mouse{
		col:    wht,
		points: []float64{284, 320, 294, 327, 293, 329, 291, 333, 289, 333, 286, 331, 284, 325, 284, 320}}
	pupilr = mouse{
		col: wht,
		points: []float64{265, 316, 275, 320, 275, 325, 273, 330, 271, 332, 269, 333, 267, 331, 265, 327,
			265, 316}}
	nose = mouse{
		col: wht,
		points: []float64{285, 308, 288, 302, 294, 301, 298, 301, 302, 303, 305, 305, 308, 308, 309, 310,
			310, 312, 310, 316, 308, 320, 305, 323, 302, 324, 297, 324, 294, 322, 288, 317,
			286, 312, 285, 308}}
	nose1 = mouse{
		col:    wht,
		points: []float64{275, 313, 280, 317, 286, 319}}
	buttonl = mouse{
		col:    red,
		points: []float64{201, 210, 194, 208, 190, 196, 191, 187, 199, 188, 208, 200, 201, 210}}
	buttonr = mouse{
		col:    red,
		points: []float64{224, 213, 221, 209, 221, 197, 228, 191, 232, 200, 230, 211, 224, 213}}
	tail = mouse{
		col:    wht,
		points: []float64{40, 80, 50, 76, 66, 79, 90, 102, 106, 151, 128, 173, 145, 180}}
	cuffl = mouse{
		col:    wht,
		points: []float64{202, 143, 197, 148, 188, 150, 160, 137}}
	cuffr = mouse{
		col:    wht,
		points: []float64{243, 153, 233, 154, 217, 145}}
	legl = mouse{
		col:    wht,
		points: []float64{239, 153, 244, 134, 243, 96, 229, 98, 231, 130, 226, 150, 233, 154, 239, 153}}
	legr = mouse{
		col:    wht,
		points: []float64{188, 150, 187, 122, 182, 92, 168, 91, 172, 122, 173, 143, 188, 150}}
	shoel = mouse{
		col: wht,
		points: []float64{230, 109, 223, 107, 223, 98, 228, 90, 231, 76, 252, 70, 278, 73, 288, 82,
			284, 97, 271, 99, 251, 100, 244, 106, 230, 109}}
	shoel1 = mouse{
		col:    wht,
		points: []float64{223, 98, 229, 98, 243, 96, 251, 100}}
	shoel2 = mouse{
		col:    wht,
		points: []float64{271, 99, 248, 89}}
	shoer = mouse{
		col: wht,
		points: []float64{170, 102, 160, 100, 160, 92, 163, 85, 157, 82, 160, 73, 178, 66, 215, 63,
			231, 76, 228, 90, 213, 97, 195, 93, 186, 93, 187, 100, 184, 102, 170, 102}}
	shoer1 = mouse{
		col:    wht,
		points: []float64{160, 92, 168, 91, 182, 92, 186, 93}}
	shoer2 = mouse{
		col:    wht,
		points: []float64{195, 93, 182, 83}}
	tick1 = mouse{
		col:    blu,
		points: []float64{302, 432, 310, 446}}
	tick2 = mouse{
		col:    blu,
		points: []float64{370, 365, 384, 371}}
	tick3 = mouse{
		col:    blu,
		points: []float64{395, 270, 410, 270}}
	tick4 = mouse{
		col:    blu,
		points: []float64{370, 180, 384, 173}}
	tick5 = mouse{
		col:    blu,
		points: []float64{302, 113, 310, 100}}
	tick7 = mouse{
		col:    blu,
		points: []float64{119, 113, 110, 100}}
	tick8 = mouse{
		col:    blu,
		points: []float64{40, 173, 52, 180}}
	tick9 = mouse{
		col:    blu,
		points: []float64{10, 270, 25, 270}}
	tick10 = mouse{
		col:    blu,
		points: []float64{40, 371, 52, 365}}
	tick11 = mouse{
		col:    blu,
		points: []float64{110, 446, 119, 432}}
	tick12 = mouse{
		col:    blu,
		points: []float64{210, 455, 210, 470}}
	armh = mouse{
		col:    wht,
		points: []float64{-8, 0, 9, 30, 10, 70, 8, 100, 20, 101, 23, 80, 22, 30, 4, -5}}
	armm = mouse{
		col:    wht,
		points: []float64{-8, 0, 10, 80, 8, 130, 22, 134, 25, 80, 4, -5}}
	handm = mouse{
		col: wht,
		points: []float64{8, 140, 5, 129, 8, 130, 22, 134, 30, 137, 27, 143, 33, 163, 30, 168,
			21, 166, 18, 170, 12, 168, 10, 170, 5, 167, 4, 195, -4, 195, -6, 170,
			0, 154, 8, 140}}
	handm1 = mouse{
		col:    wht,
		points: []float64{0, 154, 5, 167}}
	handm2 = mouse{
		col:    wht,
		points: []float64{14, 167, 12, 158, 10, 152}}
	handm3 = mouse{
		col:    wht,
		points: []float64{12, 158, 18, 152, 21, 166}}
	handm4 = mouse{
		col:    wht,
		points: []float64{20, 156, 29, 151}}
	handh = mouse{
		col: wht,
		points: []float64{20, 130, 15, 135, 6, 129, 4, 155, -4, 155, -6, 127, -8, 121, 4, 108,
			3, 100, 8, 100, 20, 101, 23, 102, 21, 108, 28, 126, 24, 132, 20, 130}}
	handh1 = mouse{
		col:    wht,
		points: []float64{20, 130, 16, 118}}

	all = []mouse{
		handh1,
		handh,
		handm4,
		handm3,
		handm2,
		handm1,
		handm,
		armm,
		armh,
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
		legr,
		legl,
		cuffr,
		cuffl,
		tail,
		buttonr,
		buttonl,
		nose1,
		nose,
		pupilr,
		pupill,
		eyer,
		eyel,
		pants,
		shirt,
		tongue1,
		tongue,
		mouth2,
		mouth1,
		mouth,
		head,
	}
)

func main() {

	if false {
		dest := image.NewRGBA(image.Rect(0, 0, 480, 640.0))
		gc := draw2dimg.NewGraphicContext(dest)
		s, err := gopher2.Main(gc, "png")
		log.Printf("%q %v", s, err)
	}
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 480, 640.0))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)

	// Draw a closed shape
	gc.MoveTo(10, 10) // should always be called first for a new path
	gc.LineTo(100, 50)
	gc.QuadCurveTo(100, 10, 10, 10)
	gc.Close()
	gc.FillStroke()

	gc.MoveTo(110, 110) // should always be called first for a new path
	gc.LineTo(200, 250)
	gc.QuadCurveTo(120, 10, 10, 10)
	gc.Close()
	gc.SetFillColor(red)
	gc.SetStrokeColor(blu)
	gc.SetLineWidth(8)
	gc.FillStroke()

	for i, m := range all {
		m := m
		s := m.points
		gc := draw2dimg.NewGraphicContext(dest)
		log.Printf("%d: col %v", i, m.col)
		gc.MoveTo(480-float64(s[0]), 640-float64(s[1]))
		for i := 2; i < len(s); i += 2 {
			gc.LineTo(480-float64(s[i]), 640-float64(s[i+1]))
		}
		gc.Close()
		gc.SetLineWidth(1)
		gc.SetFillColor(m.col)
		gc.SetStrokeColor(m.col)
		gc.FillStroke()
	}
	// Set the font luximbi.ttf
	gc.SetFontData(draw2d.FontData{Name: "luxi", Family: draw2d.FontFamilyMono, Style: draw2d.FontStyleBold | draw2d.FontStyleItalic})
	// Set the fill text color to black
	gc.SetFillColor(image.White)
	gc.SetFontSize(14)
	// Display Hello World
	gc.FillStringAt("Hello World", 8, 52)
	gc.FillStroke()
	gc.FillStringAt("0,0", 0, 0)
	gc.FillStroke()
	gc.FillStringAt("10,10", 10, 10)
	gc.FillStroke()
	gc.FillStringAt("100,100", 100, 100)
	gc.FillStroke()
	gc.FillStringAt("200,200", 200, 200)
	gc.FillStroke()
	gc.FillStringAt("300,300", 300, 300)
	gc.FillStroke()

	// Save to file
	draw2dimg.SaveToPngFile("hello.png", dest)
}
