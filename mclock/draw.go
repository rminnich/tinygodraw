package main

import (
	"log"
	"time"

	_ "tinygo.org/x/drivers"
)

func main() {
	display := new()
	for {
		redraw(display)
		time.Sleep(time.Second)
	}
	log.Printf("all done")
}
