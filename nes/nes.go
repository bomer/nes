package nes

import (
	"fmt"
)

type Nes struct {
	Cpu Cpu
}

func WriteMemory(address uint16, value byte) {
	nes.Cpu.Memory [address]=value
	fmt.Printf("Writing adress %02x with %d, old =  \n", address, value) //Cpu.Memory[address]

}

func (nes *Nes) Init() {
	nes.Cpu.Init()
	LoadGame("mario.nes")

}
