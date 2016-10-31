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

func (self *Cpu) ReadAddress() uint16 {
	b1 := uint16(self.Memory[self.PC+1])
	b2 := uint16(self.Memory[self.PC+2])
	fmt.Printf("Op Code %02x , B1=%02x b2=%02x", self.instruction, b1, b2)
	address := uint16(b2)<<8 | b1
	return address
}
func (self *Cpu) ReadWrappedAddress(a uint16) uint16 {
	//a:= passed
	b := (a & 0xFF00) | uint16(byte(a)+1)
	b1 := uint16(self.Memory[a])
	b2 := uint16(self.Memory[b])
	address := uint16(b2)<<8 | b1
	return address
}

func (self *Cpu) DecodeInstruction() {
	fmt.Printf("About to run instruction at %d", self.PC)
	self.instruction = self.Memory[self.PC]
	fmt.Printf("memory val = %02x \n", self.instruction)
	self.info = OpTable[int(self.instruction)]
	fmt.Printf("Instruction self.info %+v \n", self.info)
	fmt.Printf("Mode - %s, Operation - %s \n", self.info.ModeString(), self.info.OperationString())

	fmt.Printf("\n")

	var address uint16 //Address of what we're going to read based on the MODE
	switch self.info.Mode {
	case Mode_Absolute:
		// address=self.Memory
		// abcd stored in x=34 x+1=12
		address = self.ReadAddress()
		break
	case Mode_AbsoluteX:
		address = self.ReadAddress() + uint16(self.X)
		break

	case Mode_AbsoluteY:
		address = self.ReadAddress() + uint16(self.Y)
		break

	case Mode_Indirect: // TODO, need to do indirect_X and Y. Contains bug
		address = self.ReadWrappedAddress(self.ReadAddress())
		break

	// case Mode_IndirectX:
	// 	address = self.ReadWrappedAddress(self.ReadAddress() + self.X)

	// case Mode_IndirectY:
	// 	address = self.ReadWrappedAddress(self.ReadAddress() + self.Y)

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
	//Run Operation
	self.info.RunOperation(self)
	fmt.Println("Op Executed \n")
	// self.PC += self.info.No_Bytes
	fmt.Println("Done with this op.... \n\n")

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

// func Brk(self *Cpu) {
// 	fmt.Println("BRK OP")
// 	fmt.Printf("%v", self)
// }

func Adc(self *Cpu) {
	fmt.Println("Running Op Adc")
}
func And(self *Cpu) {
	fmt.Println("Running Op And")
}
func Asl(self *Cpu) {
	fmt.Println("Running Op Asl")
}
func Bcc(self *Cpu) {
	fmt.Println("Running Op Bcc")
}
func Bcs(self *Cpu) {
	fmt.Println("Running Op Bcs")
}
func Beq(self *Cpu) {
	fmt.Println("Running Op Beq")
}
func Bit(self *Cpu) {
	fmt.Println("Running Op Bit")
}
func Bmi(self *Cpu) {
	fmt.Println("Running Op Bmi")
}
func Bne(self *Cpu) {
	fmt.Println("Running Op Bne")
}
func Bpl(self *Cpu) {
	fmt.Println("Running Op Bpl")
}
func Brk(self *Cpu) {
	fmt.Println("Running Op Brk, cpu info  - ")
	fmt.Printf("%+v", self.info)
}
func Bvc(self *Cpu) {
	fmt.Println("Running Op Bvc")
}
func Bvs(self *Cpu) {
	fmt.Println("Running Op Bvs")
}
func Clc(self *Cpu) {
	fmt.Println("Running Op Clc")
}
func Cld(self *Cpu) {
	fmt.Println("Running Op Cld")
}
func Cli(self *Cpu) {
	fmt.Println("Running Op Cli")
}
func Clv(self *Cpu) {
	fmt.Println("Running Op Clv")
}
func Cmp(self *Cpu) {
	fmt.Println("Running Op Cmp")
}
func Cpx(self *Cpu) {
	fmt.Println("Running Op Cpx")
}
func Cpy(self *Cpu) {
	fmt.Println("Running Op Cpy")
}
func Dec(self *Cpu) {
	fmt.Println("Running Op Dec")
}
func Dex(self *Cpu) {
	fmt.Println("Running Op Dex")
}
func Dey(self *Cpu) {
	fmt.Println("Running Op Dey")
}
func Eor(self *Cpu) {
	fmt.Println("Running Op Eor")
}
func Inc(self *Cpu) {
	fmt.Println("Running Op Inc")
}
func Inx(self *Cpu) {
	fmt.Println("Running Op Inx")
}
func Iny(self *Cpu) {
	fmt.Println("Running Op Iny")
}
func Jmp(self *Cpu) {
	fmt.Println("Running Op Jmp")
}
func Jsr(self *Cpu) {
	fmt.Println("Running Op Jsr")
}
func Lda(self *Cpu) {
	fmt.Println("Running Op Lda")
}
func Ldx(self *Cpu) {
	fmt.Println("Running Op Ldx")
}
func Ldy(self *Cpu) {
	fmt.Println("Running Op Ldy")
}
func Lsr(self *Cpu) {
	fmt.Println("Running Op Lsr")
}
func Nop(self *Cpu) {
	fmt.Println("Running Op Nop")
}
func Ora(self *Cpu) {
	fmt.Println("Running Op Ora")
}
func Pha(self *Cpu) {
	fmt.Println("Running Op Pha")
}
func Php(self *Cpu) {
	fmt.Println("Running Op Php")
}
func Pla(self *Cpu) {
	fmt.Println("Running Op Pla")
}
func Plp(self *Cpu) {
	fmt.Println("Running Op Plp")
}
func Rol(self *Cpu) {
	fmt.Println("Running Op Rol")
}
func Ror(self *Cpu) {
	fmt.Println("Running Op Ror")
}
func Rti(self *Cpu) {
	fmt.Println("Running Op Rti")
}
func Rts(self *Cpu) {
	fmt.Println("Running Op Rts")
}
func Sbc(self *Cpu) {
	fmt.Println("Running Op Sbc")
}
func Sec(self *Cpu) {
	fmt.Println("Running Op Sec")
}
func Sed(self *Cpu) {
	fmt.Println("Running Op Sed")
}
func Sei(self *Cpu) {
	fmt.Println("Running Op Sei")
}
func Sta(self *Cpu) {
	fmt.Println("Running Op Sta")
}
func Stx(self *Cpu) {
	fmt.Println("Running Op Stx")
}
func Sty(self *Cpu) {
	fmt.Println("Running Op Sty")
}
func Tax(self *Cpu) {
	fmt.Println("Running Op Tax")
}
func Tay(self *Cpu) {
	fmt.Println("Running Op Tay")
}
func Tsx(self *Cpu) {
	fmt.Println("Running Op Tsx")
}
func Txa(self *Cpu) {
	fmt.Println("Running Op Txa")
}
func Txs(self *Cpu) {
	fmt.Println("Running Op Txs")
}
func Tya(self *Cpu) {
	fmt.Println("Running Op Tya")
}
