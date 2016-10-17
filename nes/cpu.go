package nes

import "fmt"

type Cpu struct {
	PC uint16 //Programing Counter, which instruction to read next
	SP byte   //Stack pointer,

	//Registers
	A byte //Accumlator, copying to and from memory + maths fuctions
	X byte //X Register
	Y byte //Y Register
	S byte //Status, 8 flags
	//7 6 5 4 3 2 1 0
	//Z V   B D I Z C
	//C=Carry, Z=Zero, I=Interupt, D=Decimal,B=Brk/software interupt, V-Overflow,S=Sign, 1=negative

	//64 kb of memory, adressing space of 0x0000 to 0xFFFF
	Memory [0xffff]byte

	RomReader Rom
}

func (self *Cpu) WriteMemory(address uint16, value byte) {
	fmt.Printf("CPU-Writing adress %02x with %d \n", address, value)
	// fmt.Printf(self.Memory)
	self.Memory[address] = value

}

func (self *Cpu) PrintInstruction() {

}

func (self *Cpu) loadRom() {
	// self.RomReader.LoadGame("mario.nes")
	// LoadGame("mario.nes")
}

func (self *Cpu) Init() {
	fmt.Printf("Mode_Absolute %d \n", Mode_Absolute)
	fmt.Printf("Mode_Absolute %% \n", OpTable[0])
	self.PC = 0xFFFC
	self.SP = 0x00
	// self.loadRom()

}
func (self *Cpu) EmulateCycle() {

}
