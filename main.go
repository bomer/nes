package main

import (
	"fmt"
	"github.com/bomer/nes/nes"
)

var myNes nes.Nes

func main() {

	fmt.Printf("Initing...")
	myNes.Cpu.Quiet = true
	myNes.Rom.LoadGame("mario.nes", &myNes.Cpu)
	myNes.Cpu.Quiet = false
	myNes.Init()
	// myNes.Cpu

	// for {
	// 	myNes.Cpu.EmulateCycle()
	// }

}
