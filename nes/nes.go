// Package nes contains everything that lives within the NES system. main entry point above it should be portable.
package nes

import (
	"fmt"
	"time"
)

// Nes is the top level, representing the physical system, it contains all NES Compontents, trying to keep as close to the original system at possible
type Nes struct {
	Cpu Cpu
	Rom Rom
	Ppu Ppu
}

// Pause is a helper function that lets me manually step through op code execution one at a time
func Pause() {
	var i int
	a, _ := fmt.Scanf("Paused.. enter to continue%d", &i)
	fmt.Printf("%d", a)
}

// CyclesPerSecond = 1.79mhz
const CyclesPerSecond = 1789773
const CyclesPerFrame = 29780

// Init Starts NES system. This controls the main loop and emulation of CPU Cycles
func (nes *Nes) Init() {

	nes.Cpu.Init()
	// fmt.Printf("%%", nes.Cpu.Memory) // Check to see if ROM loaded in CPU RAM
	fmt.Println("Mario Loaded")

	//Run emulator on another go-routine
	//Else emulator runs to slow on main thread.
	// go func() {
	// Target 60 Frames Per Second (approx 16.67ms per frame)
	frameTicker := time.NewTicker(time.Second / 60)
	defer frameTicker.Stop()

	for {
		// Run a whole frame's worth of cycles as fast as native Go can loop
		cyclesThisFrame := 0
		for cyclesThisFrame < CyclesPerFrame {

			// Execute 1 CPU cycle
			nes.Cpu.EmulateCycle()

			// Execute 3 PPU cycles for every 1 CPU cycle
			nes.Ppu.EmulateCycle()
			nes.Ppu.EmulateCycle()
			nes.Ppu.EmulateCycle()

			cyclesThisFrame++
		}

		// Now that a full frame is drawn, wait here to throttle it to 60 FPS
		<-frameTicker.C

		// (Optional) This is where you would blit your FullBuffer to a GUI
		// or clear it for the next frame!
	}

}
