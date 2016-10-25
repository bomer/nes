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

	var address uint16 //Address of what we're going to read based on the MODE
	switch info.Mode {
	case Mode_Absolute:
		// address=self.Memory
		// abcd stored in x=34 x+1=12
		b1 := byte(self.Memory[self.PC+1])
		b2 := byte(self.Memory[self.PC+2])

		fmt.Printf("Op Code %02x , B1=%02x b2=%02x", self.instruction, b1, b2)
		address = uint16(b2>>8 | b1)
		fmt.Printf("%02x", address)
		break
	case Mode_AbsoluteX:
		b1 := byte(self.Memory[self.PC+1])
		b2 := byte(self.Memory[self.PC+2])
		address = uint16(b2>>8|b1) + uint16(self.X)
		break

	case Mode_AbsoluteY:
		b1 := byte(self.Memory[self.PC+1])
		b2 := byte(self.Memory[self.PC+2])
		address = uint16(b2>>8|b1) + uint16(self.Y)
		break

	case Mode_Indirect: // TODO, need to do indirect_X and Y. Contains bug
		break

	case Mode_Immediate:
		address = self.PC + 1
		break

	case Mode_Accumulator:
		address = 0
		break

	case Mode_Implied:
		address = 0
		break

	case Mode_Relative: //Crazy one
		bb := uint16(self.Memory[self.PC+1])
		//if number >128, then its negative, mimicing signed byte. Minus 128 in this case
		if bb < 128 {
			address = self.PC + 2 + bb
		} else {
			address = self.PC + 2 + bb - 128
		}
		break
	case Mode_ZeroPage: //Read only one one byte refference as 16 bit
		address = uint16(self.Memory[self.PC+1])
		break

	case Mode_ZeroPageX:
		address = uint16(uint16(self.Memory[self.PC+1]) + uint16(self.X))
		break
	case Mode_ZeroPageY:
		address = uint16(uint16(self.Memory[self.PC+1]) + uint16(self.Y))
		break
	}
	fmt.Printf("Got Address %02x", address)

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
