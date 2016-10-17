package nes

import (
	"fmt"
)

type Nes struct {
	Cpu Cpu
	Rom Rom
}

func (nes *Nes) WriteMemory(address uint16, value byte) {
	// Cpu.Memory[address] = value
	fmt.Printf("NES : Writing adress %02x with %02x, old =  \n", address, value) //Cpu.Memory[address]

}

func (nes *Nes) Init() {
	nes.Cpu.Init()
	nes.Rom.LoadGame("mario.nes", nes.Cpu)

}
