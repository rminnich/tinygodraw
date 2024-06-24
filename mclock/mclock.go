/*
  - mclock.c - graphical clock for Plan 9 using draw(2) API
    *
  - Graphical image is based on a clock program for Tektronix vector
  - graphics displays written in PDP-11 BASIC by Dave Robinson at the
  - University of Delaware in the 1970s.
    *
  - 071218 - initial release
  - 071223 - fix window label, fix redraw after hide, add tongue
  - /

#include <u.h>
#include <libc.h>
#include <draw.h>
#include <event.h>

uint16 anghr, angmin, dia, offx, offy;
Image *dots, *back, *blk, *wht, *red, *org, *flesh;
Tm *mtime;
*/
package main

import (
	"image/color"
	"log"
	"math"
	"time"

	"tinygo.org/x/tinydraw"
)

const (
	DBIG        = 600
	XDarkOrange = 0xff8c0000
	Xwheat      = 0xf5deb300
)

var (
	/* hair is head[0..41*2], face is head[27*2..56*2] */
	head = []int16{286, 386, 263, 410, 243, 417, 230, 415, 234, 426, 227, 443, 210, 450, 190, 448,
		172, 435, 168, 418, 175, 400, 190, 398, 201, 400, 188, 390, 180, 375, 178, 363,
		172, 383, 157, 390, 143, 388, 130, 370, 125, 350, 130, 330, 140, 318, 154, 318,
		165, 325, 176, 341, 182, 320, 195, 305, 200, 317, 212, 322, 224, 319, 218, 334,
		217, 350, 221, 370, 232, 382, 250, 389, 264, 387, 271, 380, 275, 372, 276, 381,
		279, 388, 286, 386, 300, 360, 297, 337, 294, 327, 284, 320, 300, 301, 297, 297,
		282, 286, 267, 284, 257, 287, 254, 280, 249, 273, 236, 274, 225, 290, 195, 305}

	mouth  = []int16{235, 305, 233, 297, 235, 285, 243, 280, 250, 282, 252, 288, 248, 290, 235, 305}
	mouth1 = []int16{240, 310, 235, 305, 226, 306}
	mouth2 = []int16{257, 287, 248, 290}
	tongue = []int16{235, 285, 243, 280, 246, 281, 247, 286, 245, 289, 241, 291, 237, 294, 233, 294,
		235, 285}
	tongue1 = []int16{241, 291, 241, 286}
	shirt   = []int16{200, 302, 192, 280, 176, 256, 170, 247, 186, 230, 210, 222, 225, 226, 237, 235,
		222, 291, 200, 302}
	pants = []int16{199, 164, 203, 159, 202, 143, 189, 138, 172, 135, 160, 137, 160, 166, 151, 170,
		145, 180, 142, 200, 156, 230, 170, 247, 186, 230, 210, 222, 225, 226, 237, 235,
		245, 205, 242, 190, 236, 176, 229, 182, 243, 153, 240, 150, 228, 142, 217, 145,
		212, 162, 199, 164}
	eyel   = []int16{294, 327, 296, 335, 293, 345, 285, 345, 280, 337, 281, 325, 284, 320, 294, 327}
	eyer   = []int16{275, 320, 278, 337, 275, 345, 268, 344, 260, 333, 260, 323, 264, 316, 275, 320}
	pupill = []int16{284, 320, 294, 327, 293, 329, 291, 333, 289, 333, 286, 331, 284, 325, 284, 320}
	pupilr = []int16{265, 316, 275, 320, 275, 325, 273, 330, 271, 332, 269, 333, 267, 331, 265, 327,
		265, 316}
	nose = []int16{285, 308, 288, 302, 294, 301, 298, 301, 302, 303, 305, 305, 308, 308, 309, 310,
		310, 312, 310, 316, 308, 320, 305, 323, 302, 324, 297, 324, 294, 322, 288, 317,
		286, 312, 285, 308}
	nose1   = []int16{275, 313, 280, 317, 286, 319}
	buttonl = []int16{201, 210, 194, 208, 190, 196, 191, 187, 199, 188, 208, 200, 201, 210}
	buttonr = []int16{224, 213, 221, 209, 221, 197, 228, 191, 232, 200, 230, 211, 224, 213}
	tail    = []int16{40, 80, 50, 76, 66, 79, 90, 102, 106, 151, 128, 173, 145, 180}
	cuffl   = []int16{202, 143, 197, 148, 188, 150, 160, 137}
	cuffr   = []int16{243, 153, 233, 154, 217, 145}
	legl    = []int16{239, 153, 244, 134, 243, 96, 229, 98, 231, 130, 226, 150, 233, 154, 239, 153}
	legr    = []int16{188, 150, 187, 122, 182, 92, 168, 91, 172, 122, 173, 143, 188, 150}
	shoel   = []int16{230, 109, 223, 107, 223, 98, 228, 90, 231, 76, 252, 70, 278, 73, 288, 82,
		284, 97, 271, 99, 251, 100, 244, 106, 230, 109}
	shoel1 = []int16{223, 98, 229, 98, 243, 96, 251, 100}
	shoel2 = []int16{271, 99, 248, 89}
	shoer  = []int16{170, 102, 160, 100, 160, 92, 163, 85, 157, 82, 160, 73, 178, 66, 215, 63,
		231, 76, 228, 90, 213, 97, 195, 93, 186, 93, 187, 100, 184, 102, 170, 102}
	shoer1 = []int16{160, 92, 168, 91, 182, 92, 186, 93}
	shoer2 = []int16{195, 93, 182, 83}
	tick1  = []int16{302, 432, 310, 446}
	tick2  = []int16{370, 365, 384, 371}
	tick3  = []int16{395, 270, 410, 270}
	tick4  = []int16{370, 180, 384, 173}
	tick5  = []int16{302, 113, 310, 100}
	tick7  = []int16{119, 113, 110, 100}
	tick8  = []int16{40, 173, 52, 180}
	tick9  = []int16{10, 270, 25, 270}
	tick10 = []int16{40, 371, 52, 365}
	tick11 = []int16{110, 446, 119, 432}
	tick12 = []int16{210, 455, 210, 470}
	armh   = []int16{-8, 0, 9, 30, 10, 70, 8, 100, 20, 101, 23, 80, 22, 30, 4, -5}
	armm   = []int16{-8, 0, 10, 80, 8, 130, 22, 134, 25, 80, 4, -5}
	handm  = []int16{8, 140, 5, 129, 8, 130, 22, 134, 30, 137, 27, 143, 33, 163, 30, 168,
		21, 166, 18, 170, 12, 168, 10, 170, 5, 167, 4, 195, -4, 195, -6, 170,
		0, 154, 8, 140}
	handm1 = []int16{0, 154, 5, 167}
	handm2 = []int16{14, 167, 12, 158, 10, 152}
	handm3 = []int16{12, 158, 18, 152, 21, 166}
	handm4 = []int16{20, 156, 29, 151}
	handh  = []int16{20, 130, 15, 135, 6, 129, 4, 155, -4, 155, -6, 127, -8, 121, 4, 108,
		3, 100, 8, 100, 20, 101, 23, 102, 21, 108, 28, 126, 24, 132, 20, 130}
	handh1          = []int16{20, 130, 16, 118}
	offx, offy, dia int16
)

// The API is pairs of points, not Points, so that's a small change.
func xlate(display displayer, in []int16) []int16 {
	out := make([]int16, len(in))
	for i := range out[:len(out)-1] {
		out[i] = int16(offx + (dia*(in[i])+210)/420)
		out[i+1] = int16(offy + (dia*(480-in[i+1])+210)/420)
	}
	return out
}

// let's be a bit gross about this. Just draw the lines for now.
func fillpoly(display displayer, color color.RGBA, points ...int16) {
	// TODO: verify len

	np := len(points)

	if np < 4 {
		log.Printf("fillpoly: only %d points", np)
	}
	for i := 0; i < np-4; i += 4 {
		x0, y0, x1, y1 := points[i], points[i+1], points[i+2], points[i+3]
		log.Printf("line (%d,%d) -> (%d, %d)", x0, y0, x1, y1)
		tinydraw.Line(display, x0, y0, x1, y1, color)

	}
}

func myfill(display displayer, p []int16, color color.RGBA) { // , Image* color)
	out := xlate(display, p)
	log.Printf("xlote %d -> %d", p, out)
	fillpoly(display, color, out...)
}

func mypoly(display displayer, p []int16, color color.RGBA) {
	out := xlate(display, p)
	//	var b int
	//	if dia > DBIG {
	//		b = 1
	//	}
	myfill(display, out, color)
	//fillpoly(display, , out, np, Enddisc, Enddisc, b, color, ZP)
}

func arm(display displayer, p []int16, angle float64) []int16 {

	out := make([]int16, len(p))
	for i := range out[:len(out)-1] {
		cosp := math.Cos(math.Pi * angle / 180.0)
		sinp := math.Sin(math.Pi * angle / 180.0)
		out[i] = int16(float64(p[i])*cosp + float64(p[i+1])*sinp + 210.5)
		out[i] = int16(float64(p[i+1])*cosp - float64(p[i])*sinp + 270.5)
	}
	return out
}

func polyarm(display displayer, p []int16, color color.RGBA, angle float64) {
	tmp := arm(display, p, angle)
	out := xlate(display, tmp)
	mypoly(display, out, color)
	//poly(screen, out, np, Enddisc, Enddisc, dia>DBIG?1:0, color, ZP);
}

func fillarm(display displayer, p []int16, color color.RGBA, angle float64) {

	tmp := arm(display, p, angle)
	out := xlate(display, tmp)
	mypoly(display, out, color)
	//	fillpoly(screen, out, np, ~0, color, ZP);
}

func arms(display displayer, anghr, angmin float64) {
	blk := color.RGBA{A: 255}
	wht := color.RGBA{G: 255, B: 255, R: 255, A: 255}
	/* arms */
	fillarm(display, armh[:8], blk, anghr)
	fillarm(display, armm[:6], blk, angmin)

	/* hour hand */
	fillarm(display, handh[:16], wht, anghr)
	polyarm(display, handh[:16], blk, anghr)
	polyarm(display, handh1[:2], blk, anghr)

	/* minute hand */
	fillarm(display, handm[:18], wht, angmin)
	polyarm(display, handm[:18], blk, angmin)
	polyarm(display, handm1[:2], blk, angmin)
	polyarm(display, handm2[:3], blk, angmin)
	polyarm(display, handm3[:3], blk, angmin)
	polyarm(display, handm4[:2], blk, angmin)
}

func redraw(display displayer) {
	blk := color.RGBA{A: 255}
	red := color.RGBA{R: 255, A: 255}
	wht := color.RGBA{G: 255, B: 255, R: 255, A: 255}
	dots := color.RGBA{G: 64, B: 64, R: 64, A: 255}
	org := color.RGBA{G: 64, B: 0, R: 64, A: 255}
	flesh := color.RGBA{G: 0, B: 64, R: 64, A: 255}
	n := time.Now()
	anghr := float64(n.Hour()*30) + float64(n.Minute()/2)
	angmin := float64(n.Minute() * 6)

	dia = 200
	//dia = Dx(screen->r) < Dy(screen->r) ? Dx(screen->r) : Dy(screen->r);
	//var offx, offy int16
	//offx := screen->r.min.x + (Dx(screen->r) - dia) / 2;
	//offy := screen->r.min.y + (Dy(screen->r) - dia) / 2;

	// draw(screen, screen->r, back, nil, ZP);

	/* first draw the filled areas */
	/* hair is head[0..41*2], face is head[27*2..56*2] */
	myfill(display, head[:41*2], blk) /* hair */
	return
	myfill(display, head[27*2:56*2], flesh) /* face */
	myfill(display, mouth[:8], blk)
	myfill(display, tongue[:9], red)
	myfill(display, shirt[:10], blk)
	myfill(display, pants[:26], red)
	myfill(display, buttonl[:7], wht)
	myfill(display, buttonr[:7], wht)
	myfill(display, eyel[:8], wht)
	myfill(display, eyer[:8], wht)
	myfill(display, pupill[:8], blk)
	myfill(display, pupilr[:9], blk)
	myfill(display, nose[:18], blk)
	myfill(display, shoel[:13], org)
	myfill(display, shoer[:16], org)
	myfill(display, legl[:8], blk)
	myfill(display, legr[:7], blk)

	/* outline the color-filled areas */
	mypoly(display, head[27*2:], blk) /* face */
	mypoly(display, tongue[:9], blk)
	mypoly(display, pants[:26], blk)
	mypoly(display, buttonl[:7], blk)
	mypoly(display, buttonr[:7], blk)
	mypoly(display, eyel[:8], blk)
	mypoly(display, eyer[:8], blk)
	mypoly(display, shoel[:13], blk)
	mypoly(display, shoer[:16], blk)

	/* draw the details */
	mypoly(display, nose1[:3], blk)
	mypoly(display, mouth1[:3], blk)
	mypoly(display, mouth2[:2], blk)
	mypoly(display, tongue1[:2], blk)
	mypoly(display, tail[:7], blk)
	mypoly(display, cuffl[:4], blk)
	mypoly(display, cuffr[:3], blk)
	mypoly(display, shoel1[:4], blk)
	mypoly(display, shoel2[:2], blk)
	mypoly(display, shoer1[:4], blk)
	mypoly(display, shoer2[:2], blk)
	mypoly(display, tick1[:2], dots)
	mypoly(display, tick2[:2], dots)
	mypoly(display, tick3[:2], dots)
	mypoly(display, tick4[:2], dots)
	mypoly(display, tick5[:2], dots)
	mypoly(display, tick7[:2], dots)
	mypoly(display, tick8[:2], dots)
	mypoly(display, tick9[:2], dots)
	mypoly(display, tick10[:2], dots)
	mypoly(display, tick11[:2], dots)
	mypoly(display, tick12[:2], dots)

	arms(display, anghr, angmin)

	display.Display()
	return
}

/*
void
main(void)
{
	Event e;
	Mouse m;
	Menu menu;
	char *mstr[] = {"exit", 0};
	uint16 key, timer, oldmin;

	initdraw(0,0,"mclock");
	back = allocimagemix(display, DPalebluegreen, DWhite);

	dots = allocimage(display, Rect(0,0,1,1), CMAP8, 1, DBlue);
	blk = allocimage(display, Rect(0,0,1,1), CMAP8, 1, DBlack);
	wht = allocimage(display, Rect(0,0,1,1), CMAP8, 1, DWhite);
	red = allocimage(display, Rect(0,0,1,1), CMAP8, 1, DRed);
	org = allocimage(display, Rect(0,0,1,1), CMAP8, 1, XDarkOrange);
	flesh = allocimage(display, Rect(0,0,1,1), CMAP8, 1, Xwheat);

	mtime = localtime(time(0));
	redraw(screen);

	einit(Emouse);
	timer = etimer(0, 30*1000);

	menu.item = mstr;
	menu.lasthit = 0;

	for(;;) {
		key = event(&e);
		if(key == Emouse) {
			m = e.mouse;
			if(m.buttons & 4) {
				if(emenuhit(3, &m, &menu) == 0)
					exits(0);
			}
		} else if(key == timer) {
			oldmin = mtime->min;
			mtime = localtime(time(0));
			if(mtime->min != oldmin) redraw(screen);
		}
	}
}
*/
