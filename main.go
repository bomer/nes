package main

import (
	"fmt"
	"github.com/bomer/nes/nes"
)

var myNes nes.Nes

func main() {

	fmt.Printf("Initing...")
	myNes.Cpu.Init()

	for {
		myNes.Cpu.EmulateCycle()
	}

}
