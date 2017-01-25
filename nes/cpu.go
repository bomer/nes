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
	S byte //Status, 8 flags, for more info see statusflags.go file

	//64 kb of memory, adressing space of 0x0000 to 0xFFFF
	Memory [0xFFFF + 1]byte

	//Address to be used for read ops
	address uint16

	RomReader Rom

	//non Global items, used to track instruction info
	instruction byte       //Current value stored at memory[pc]
	info        OpCodeInfo //Stores information about how to read the full op code

	Quiet bool //
}

//Made possible by http://stackoverflow.com/questions/47981/how-do-you-set-clear-and-toggle-a-single-bit-in-c-c
func (self *Cpu) SetFlag(flag int, tovalue bool) {
	fmt.Printf("Setting flag at positon %d to %s", flag, tovalue)
	fmt.Printf("Before - %b", self.S)
	//check
	// n |= (1 << self.S)
	if tovalue {
		self.S |= 1 << uint8(flag)
	} else {
		self.S &= ^(1 << uint8(flag))
	}
	fmt.Printf("After - %b \n", (self.S))

}

//Made possible by http://stackoverflow.com/questions/47981/how-do-you-set-clear-and-toggle-a-single-bit-in-c-c
func (self *Cpu) GetFlag(flag int) bool {
	fmt.Printf("Getting flag at positon %d ", flag)
	//check byte, will store the value of the pos like 128 or 0
	var n byte

	//Checking with logic and
	n = self.S & (1 << uint8(flag))
	if n == 0 {
		return false
	}
	return true

}

func (self *Cpu) WriteMemory(address uint16, value byte) {
	if !self.Quiet {
		fmt.Printf("CPU-Writing adress %02x with %d \n", address, value)
	}
	//TODO. Extra mapping, mirrors, etc.
	self.Memory[address] = value
}

func (self *Cpu) ReadAddressByte(start uint16) uint8 {
	return uint8(self.Memory[start])
}

func (self *Cpu) ReadAddress(start uint16) uint16 {
	b1 := uint16(self.Memory[start])
	b2 := uint16(self.Memory[start+1])
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

func (self *Cpu) Push(value byte) {
	pos := uint16(0x100) + uint16(self.SP)
	self.WriteMemory(pos, value)
	self.SP--
}

func (self *Cpu) Pull() byte {
	self.SP++
	pos := uint16(0x100) + uint16(self.SP)
	ret := self.ReadAddressByte(pos)
	return ret
}

//Check for negative & zero, common on sets +  calculations
func (self *Cpu) CheckNZ(value byte) {
	if value == 0 {
		self.SetFlag(Status_Z, true)
	} else {
		self.SetFlag(Status_Z, false)
	}
	fmt.Printf("checking 7th bit of %b", value)
	checkbit := value >> 7 & 1
	// checkbit := value & 128 // if 7th bit == 1
	if checkbit == 1 {
		self.SetFlag(Status_N, true)
	} else {
		self.SetFlag(Status_N, false)
	}
}

func (self *Cpu) DecodeInstruction() {
	fmt.Printf("About to run instruction at %d\n", self.PC)
	self.instruction = self.Memory[self.PC]
	fmt.Printf("Instruction %02x \n", self.instruction)
	self.info = OpTable[int(self.instruction)]
	// fmt.Printf("Instruction self.info %+v \n", self.info)
	fmt.Printf("Mode - %s, Operation - %s \n", self.info.ModeString(), self.info.OperationString())

	fmt.Printf("\n")

	var address uint16 //Address of what we're going to read based on the MODE
	switch self.info.Mode {
	case Mode_Absolute:
		// address=self.Memory
		// abcd stored in x=34 x+1=12
		address = self.ReadAddress(self.PC + 1)
		break
	case Mode_AbsoluteX:
		address = self.ReadAddress(self.PC+1) + uint16(self.X)
		break

	case Mode_AbsoluteY:
		address = self.ReadAddress(self.PC+1) + uint16(self.Y)
		break

	case Mode_Indirect: // TODO, need to do indirect_X and Y. Contains bug
		address = self.ReadWrappedAddress(self.ReadAddress(self.PC + 1))
		break

	case Mode_IndirectX:
		address = self.ReadWrappedAddress(self.ReadAddress(self.PC+1) + uint16(self.X))
		break

	case Mode_IndirectY:
		address = self.ReadWrappedAddress(self.ReadAddress(self.PC+1) + uint16(self.Y))
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
		fmt.Printf("bb = %v", bb)
		//if number >128, then its negative, mimicing signed byte. Minus 128 in this case
		//Removed a +2 on both of these lines because we always add 2 at the end.
		if bb < 128 {
			address = self.PC + bb
		} else {
			address = self.PC - (bb - 128)
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
	self.address = address
	//Run Operation
	self.info.RunOperation(self)
	fmt.Println("Op Executed \n")
	self.PC += uint16(self.info.No_Bytes)
	fmt.Println("Done with this op.... \n\n")
	Pause()

}

func (self *Cpu) loadRom() {
	// self.RomReader.LoadGame("mario.nes")
	// LoadGame("mario.nes")
}

func (self *Cpu) Init() {
	// Test mode lookup table
	// fmt.Printf("Mode_Absolute %d \n", Mode_Absolute)
	// fmt.Printf("Mode_Absolute %+v \n", OpTable[0x00])
	Pause()

	self.PC = 0xFFFC //Loads back a step then reads ahead like a normal op code
	self.PC = self.ReadAddress(self.PC)
	self.SP = 0xff
	self.X = 0
	self.Y = 0

}
func (self *Cpu) EmulateCycle() {
	self.DecodeInstruction()

}

//ADC  Add Memory to Accumulator with Carry
//Must set Carry and Overflow Flag
func Adc(self *Cpu) {
	fmt.Println("Running Op Adc")
	// A + M + C -> A, C
	a := self.A
	m := self.ReadAddressByte(self.address)
	fmt.Printf("Adding a (%b) and mem(%b)", a, m)
	var c byte = 0
	if self.GetFlag(Status_C) {
		c = 1
	}

	self.A = a + m + c
	self.CheckNZ(self.A)

	if int(a)+int(m)+int(c) > 0xFF { // if overflow
		fmt.Println("CARRY Flag enabled cause > 0xff!")
		Sec(self) //set carry flag
	} else {
		Clc(self) //clear carry flag
	}
	//if overflow, that is negative flag on when shouldnt be
	//if only 1 of a or m had flag, but after a or the combination didnt, overflow!!
	if (a^m)&0x80 == 0 && (a^self.A)&0x80 != 0 {
		fmt.Println("OVERFLOW HIT!")
		self.SetFlag(Status_V, true)
	} else {
		self.SetFlag(Status_V, false)

	}

}

//AND, AND Memory with Accumulator
// A AND M -> A
func And(self *Cpu) {
	fmt.Println("Running Op And")
	m := self.ReadAddressByte(self.address)
	self.A &= m
	self.CheckNZ(self.A)

}

//ASL  Shift Left One Bit (Memory or Accumulator)
// C <- [76543210] <- 0
func Asl(self *Cpu) {
	fmt.Println("Running Op Asl")
	if self.info.Mode == Mode_Accumulator { // Interact with self.A
		carry := self.A >> 7
		if carry > 0 {
			self.SetFlag(Status_C, true)
		} else {
			self.SetFlag(Status_C, false)
		}
		self.A = self.A << 1
		self.CheckNZ(self.A)
	} else { // Interact with memory read byte, read first, then modify, and write back
		m := self.ReadAddressByte(self.address)
		carry := m >> 7
		if carry > 0 {
			self.SetFlag(Status_C, true)
		} else {
			self.SetFlag(Status_C, false)
		}
		m = m << 1
		self.WriteMemory(self.address, m)
		self.CheckNZ(m)

	}
}
func Bcc(self *Cpu) {
	fmt.Println("Running Op Bcc")
	if self.GetFlag(Status_C) == false {
		self.PC = self.address
	}
}
func Bcs(self *Cpu) {
	fmt.Println("Running Op Bcs")
	if self.GetFlag(Status_C) == true {
		self.PC = self.address
	}
}

//BEQ Branch on Result Zero
func Beq(self *Cpu) {
	fmt.Println("Running Op Beq")
	if self.GetFlag(Status_Z) == true {
		self.PC = self.address
	}
}

//TODO OH GOD WHAT IS THIS!
func Bit(self *Cpu) {
	fmt.Println("Running Op Bit")
}

//BMI  Branch on Result Minus
func Bmi(self *Cpu) {
	fmt.Println("Running Op Bmi")
	if self.GetFlag(Status_N) == true {
		self.PC = self.address
	}
}

//BNE  Branch on Result not Zero
func Bne(self *Cpu) {
	fmt.Println("Running Op Bne")
	if self.GetFlag(Status_Z) == false {
		self.PC = self.address
	}
}

// BPL  Branch on Result Plus
func Bpl(self *Cpu) {
	fmt.Println("Running Op Bpl")
	if self.GetFlag(Status_N) == false {
		self.PC = self.address
	}
}
func Brk(self *Cpu) {
	fmt.Println("Running Op Brk, cpu info  - ")
	fmt.Printf("%+v", self.info)
}

// BVC  Branch on Overflow Clear
func Bvc(self *Cpu) {
	fmt.Println("Running Op Bvc")
	if self.GetFlag(Status_V) == false {
		self.PC = self.address
	}
}

// BVS  Branch on Overflow Set
func Bvs(self *Cpu) {
	fmt.Println("Running Op Bvs")
	if self.GetFlag(Status_V) == true {
		self.PC = self.address
	}
}
func Clc(self *Cpu) {
	fmt.Println("Running Op Clc")
	self.SetFlag(Status_C, false)
}

//Clear Decimal Flag
func Cld(self *Cpu) {
	fmt.Println("Running Op Cld")
	self.SetFlag(Status_D, false)
}
func Cli(self *Cpu) {
	fmt.Println("Running Op Cli")
	self.SetFlag(Status_I, false)
}
func Clv(self *Cpu) {
	fmt.Println("Running Op Clv")
	self.SetFlag(Status_V, false)
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

//DEC  Decrement Memory by One
func Dec(self *Cpu) {
	fmt.Println("Running Op Dec")
	m := self.ReadAddressByte(self.address)
	m--
	self.WriteMemory(self.address, m)
	self.CheckNZ(m)
}

//DEX  Decrement Index X by One
func Dex(self *Cpu) {
	fmt.Println("Running Op Dex")
	self.X--
	self.CheckNZ(self.X)
}

//DEY  Decrement Index Y by One
func Dey(self *Cpu) {
	fmt.Println("Running Op Dey")
	self.Y--
	self.CheckNZ(self.Y)
}

//EOR , Excluse OR
// A EOR M -> A
func Eor(self *Cpu) {
	fmt.Println("Running Op Eor")
	m := self.ReadAddressByte(self.address)
	self.A ^= m
	self.CheckNZ(self.A)
}

//INC  Increment Memory by One
func Inc(self *Cpu) {
	fmt.Println("Running Op Inc")
	m := self.ReadAddressByte(self.address)
	m++
	self.WriteMemory(self.address, m)
	self.CheckNZ(m)
}

// INX  Increment X Reg by One
func Inx(self *Cpu) {
	fmt.Println("Running Op Inx")
	self.X++
	self.CheckNZ(self.X)
}

//INY  Increment Y Reg by One
func Iny(self *Cpu) {
	fmt.Println("Running Op Iny")
	self.Y++
	self.CheckNZ(self.Y)
}
func Jmp(self *Cpu) {
	fmt.Println("Running Op Jmp")
}
func Jsr(self *Cpu) {
	fmt.Println("Running Op Jsr")
}

//Load memory (M) from Address (self.address) into Accumulator
func Lda(self *Cpu) {
	fmt.Println("Running Op Lda")
	self.A = self.ReadAddressByte(self.address)
	fmt.Printf("Setting Accumulator to.. %d\n (binary of %b)", self.A, self.A)
	self.CheckNZ(self.A)
}
func Ldx(self *Cpu) {
	fmt.Println("Running Op Ldx")
	self.X = self.ReadAddressByte(self.address)
	fmt.Printf("Setting X to.. %d\n", self.X)
	self.CheckNZ(self.X)
}
func Ldy(self *Cpu) {
	fmt.Println("Running Op Ldy")
	self.Y = self.ReadAddressByte(self.address)
	fmt.Printf("Setting Y to.. %d\n", self.Y)
	self.CheckNZ(self.Y)
}
func Lsr(self *Cpu) {
	fmt.Println("Running Op Lsr")
}
func Nop(self *Cpu) {
	fmt.Println("Running Op Nop")
}

// ORA, Or memory with accumulator
// A OR M -> A
func Ora(self *Cpu) {
	fmt.Println("Running Op Ora")
	m := self.ReadAddressByte(self.address)
	self.A |= m
	self.CheckNZ(self.A)
}
func Pha(self *Cpu) {
	fmt.Println("Running Op Pha")
	self.Push(self.A)
}
func Php(self *Cpu) {
	fmt.Println("Running Op Php")
}
func Pla(self *Cpu) {
	fmt.Println("Running Op Pla")
	self.A = self.Pull()
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
	// A - M - C -> A
	a := self.A
	m := self.ReadAddressByte(self.address)
	fmt.Printf("subtracting a (%b) and mem(%b)", a, m)
	var c byte = 0
	if self.GetFlag(Status_C) {
		c = 1
	}
	fmt.Printf("Before....: %d", self.A)
	self.A = a - m - (1 - c)
	self.CheckNZ(self.A)
	fmt.Printf("After...: %d", self.A)
	if int(a)-int(m)-int(1-c) >= 0 { // if overflow
		fmt.Println("CARRY Flag enabled cause >= 0")
		Sec(self) //set carry flag
	} else {
		Clc(self) //clear carry flag
	}
	//if overflow, that is negative flag on when shouldnt be
	//if only 1 of a or m had flag, but after a or the combination didnt, overflow!!
	if (a^m)&0x80 != 0 && (a^self.A)&0x80 != 0 {
		fmt.Println("OVERFLOW HIT!")
		self.SetFlag(Status_V, true)
	} else {
		self.SetFlag(Status_V, false)
	}
}

//Set Status Flag of C - Carry Flag to on (00 10 00 00)
func Sec(self *Cpu) { //Set Carry Flag
	fmt.Println("Running Op Sec")
	self.SetFlag(Status_C, true)
}

//Set Status Flag of D - Decimal Flag to on (00 00 01 00) | NOT USED IN NES
func Sed(self *Cpu) {
	fmt.Println("Running Op Sed")
	self.SetFlag(Status_D, true)
}

//Set Status Flag of I - Interupt Disable to on (00 00 01 00)
func Sei(self *Cpu) {
	fmt.Println("Running Op Sei")
	self.SetFlag(Status_I, true)
}

// STA  Store Accumulator in Memory
func Sta(self *Cpu) {
	fmt.Println("Running Op Sta")
	self.WriteMemory(self.address, self.A)
}
func Stx(self *Cpu) {
	fmt.Println("Running Op Stx")
	self.WriteMemory(self.address, self.X)
}
func Sty(self *Cpu) {
	fmt.Println("Running Op Sty")
	self.WriteMemory(self.address, self.Y)
}

//TAX Transfer Accumulator into index X
func Tax(self *Cpu) {
	fmt.Println("Running Op Tax")
	self.X = self.A
	self.CheckNZ(self.X)
}

//TAY Transfer Accumulator into index Y
func Tay(self *Cpu) {
	fmt.Println("Running Op Tay")
	self.Y = self.A
	self.CheckNZ(self.Y)
}

//TSX Transfer Stack Pointer into Index X
func Tsx(self *Cpu) {
	fmt.Printf("Running Op Tsx - Copying sp:%d to x:%d", self.SP, self.X)
	self.X = self.SP
	self.CheckNZ(self.Y)
}

//TXA  Transfer Index X to Accumulator
func Txa(self *Cpu) {
	fmt.Println("Running Op Txa - copying x: %d to a: %d", self.X, self.A)
	self.A = self.X
	self.CheckNZ(self.A)
}

//TXS  Transfer Index X to Stack Register
func Txs(self *Cpu) {
	fmt.Println("Running Op Txs")
	self.SP = self.X
	self.CheckNZ(self.SP)
}

// TYA  Transfer Index Y to Accumulator
func Tya(self *Cpu) {
	fmt.Println("Running Op Tya")
	self.A = self.Y
	self.CheckNZ(self.A)
}
