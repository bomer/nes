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

}

func (self *Cpu) Init() {
	fmt.Printf("Mode_Absolute %d \n", Mode_Absolute)
	self.PC = 0xFFFC
	self.SP = 0x00

}
