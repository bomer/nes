// Package nes contains everything that lives within the NES system. main entry point above it should be portable.
package nes

import (
	"fmt"
	"time"
)

//Nes is the top level, representing the physical system, it contains all NES Compontents, trying to keep as close to the original system at possible
type Nes struct {
	Cpu Cpu
	Rom Rom
	Ppu Ppu
}

//Pause is a helper function that lets me manually step through op code execution one at a time
func Pause() {
	var i int
	a, _ := fmt.Scanf("Paused.. enter to continue%d", &i)
	fmt.Printf("%d", a)
}

// CyclesPerSecond = 1.79mhz
const CyclesPerSecond = 1789773
const CyclesPerFrame = 29780.5

//Init Starts NES system. This controls the main loop and emulation of CPU Cycles
func (nes *Nes) Init() {

	nes.Cpu.Init()
	// fmt.Printf("%%", nes.Cpu.Memory) // Check to see if ROM loaded in CPU RAM
	fmt.Println("Mario Loaded")
	//Run emulator on another go-routine
	//Else emulator runs to slow on main thread.
	// go func() {

	emuticker := time.NewTicker(time.Second / CyclesPerSecond) //TODO - Replace with nes CPU FREQ
	for {
		nes.Cpu.EmulateCycle()
		nes.Ppu.EmulateCycle()
		nes.Ppu.EmulateCycle()
		nes.Ppu.EmulateCycle()
		<-emuticker.C
	}
	// }()

}
