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

//This is really a helper function that lets me manually step through op code execution one at a time
func Pause() {
	var i int
	a, _ := fmt.Scanf("Paused.. enter to continue%d", &i)
	fmt.Printf("%d", a)
}

//Starts NES system. This controls the main loop and emulation of CPU Cycles
func (nes *Nes) Init() {

	nes.Cpu.Init()
	// fmt.Printf("%%", nes.Cpu.Memory) // Check to see if ROM loaded in CPU RAM
	fmt.Println("Mario Loaded")
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
