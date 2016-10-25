package nes

import (
	"fmt"
	"time"
)

//Contains all NES Compontents, trying to keep as close to the original system at possible
type Nes struct {
	Cpu Cpu
	Rom Rom
}

//Starts NES system. This controlls the main loop and emulation of CPU Cycles
func (nes *Nes) Init() {
	nes.Cpu.Init()
	nes.Rom.LoadGame("mario.nes", &nes.Cpu)
	// fmt.Printf("%%", nes.Cpu.Memory) // Check to see if ROM loaded in CPU RAM

	//Run emulator on another go-routine
	//Else emulator runs to slow on main thread.
	// go func() {
	emuticker := time.NewTicker(time.Second / 30) //TODO - Replace with nes CPU FREQ, 360=faster,
	for {
		nes.Cpu.EmulateCycle()
		<-emuticker.C
	}
	// }()

}
