package nes

import (
	"fmt"
)

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
	Memory [0xFFFF + 1]byte

	RomReader Rom

	//non Global items, used to track instruction info
	instruction byte       //Current value stored at memory[pc]
	info        OpCodeInfo //Stores information about how to read the full op code
}

func (self *Cpu) WriteMemory(address uint16, value byte) {
	// fmt.Printf("CPU-Writing adress %02x with %d \n", address, value)
	//TODO. Extra mapping, mirrors, etc.
	self.Memory[address] = value
}

func (self *Cpu) DecodeInstruction() {
	fmt.Printf("About to run instruction at %d", self.PC)
	self.instruction = self.Memory[self.PC]
	fmt.Printf("memory val = %02x \n", self.instruction)
	info := OpTable[int(self.instruction)]
	fmt.Printf("Instruction info %+v \n", info)
	fmt.Printf("Mode - %s, Operation - %s \n", info.ModeString(), info.OperationString())

	fmt.Printf("\n")

	// var address uint16 //Address of what we're going to read based on the MODE
	switch info.Mode {
	case Mode_Absolute:
		// address=self.Memory
		b1 := byte(self.Memory[self.PC+1])
		b2 := byte(self.Memory[self.PC+2])

		fmt.Printf("Op Code %02x , B1=%02x b2=%02x", self.instruction, b1, b2)
	case Mode_Implied:

	}

}

func (self *Cpu) loadRom() {
	// self.RomReader.LoadGame("mario.nes")
	// LoadGame("mario.nes")
}

func (self *Cpu) Init() {
	fmt.Printf("Mode_Absolute %d \n", Mode_Absolute)
	fmt.Printf("Mode_Absolute %+v \n", OpTable[0x00])
	var i int
	a, _ := fmt.Scanf("%d", &i)
	fmt.Printf("%d", a)

	self.PC = 0xFFFC
	self.SP = 0x00
	// self.loadRom()

}
func (self *Cpu) EmulateCycle() {
	self.DecodeInstruction()

}
