//go:build tinygo

package main

import (
	         "tinygo.org/x/drivers"
       "tinygo.org/x/tinydraw/examples/initdisplay"
)

type displayer drivers.Displayer

func new() displayer {
	return initdisplay.InitDisplay()
}
