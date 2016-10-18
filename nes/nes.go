package nes

import (
	"fmt"
	"time"
)

type Nes struct {
	Cpu Cpu
	Rom Rom
}

func (nes *Nes) OLDWriteMemory(address uint16, value byte) {
	// Cpu.Memory[address] = value
	fmt.Printf("NES : Writing adress %02x with %02x, old =  \n", address, value) //Cpu.Memory[address]

}

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
