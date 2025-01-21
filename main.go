package main

import (
	"fmt"

	"github.com/bomer/nes/nes"
)

var myNes nes.Nes

func main() {

	fmt.Printf("Initing...")
	myNes.Cpu.Quiet = true
	myNes.Rom.LoadGame("mario.nes", &myNes)
	myNes.Cpu.Quiet = false
	myNes.Cpu.Debug = true
	myNes.Cpu.DebugLines = 128
	myNes.Cpu.System = &myNes
	myNes.Init()
	// myNes.Cpu

	// for {
	// 	myNes.Cpu.EmulateCycle()
	// }

}
