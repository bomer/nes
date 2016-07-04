package nes

import (
	"fmt"
)

type Nes struct {
	Cpu Cpu
}

func WriteMemory(address uint16, value byte) {
	fmt.Printf("Writing adress %02x with %d, old =  ", address, value) //Cpu.Memory[address]

}

func (nes *Nes) Init() {
	nes.Cpu.Init()
	LoadGame("mario.nes")

}
