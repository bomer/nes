package nes

import (
	"fmt"
	"log"
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

	// System References
	RomReader Rom
	System    *Nes

	//non Global items, used to track instruction info
	instruction      byte       //Current value stored at memory[pc]
	info             OpCodeInfo //Stores information about how to read the full op code
	CycleCount       uint64
	FrameCount       uint64
	InstructionCount uint64 // Not useful number, but for my debugging

	Quiet      bool // Debug Logs
	Debug      bool
	DebugLines uint64
}

// SetFlag was Made possible by http://stackoverflow.com/questions/47981/how-do-you-set-clear-and-toggle-a-single-bit-in-c-c
func (self *Cpu) SetFlag(flag int, tovalue bool) {
	fmt.Printf("Setting flag at positon %d to %t - ", flag, tovalue)
	fmt.Printf("Before - %b ", self.S)
	//check
	// n |= (1 << self.S)
	if tovalue {
		self.S |= 1 << uint8(flag)
	} else {
		self.S &= ^(1 << uint8(flag))
	}
	fmt.Printf("After - %b \n", (self.S))

}

// GetFlag Made possible by http://stackoverflow.com/questions/47981/how-do-you-set-clear-and-toggle-a-single-bit-in-c-c
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

// Read Memory from an address. This needs to take into account the full memory map.
// Todo - Add full memory mapp as per below.

// Address range	Size	Device
// $0000–$07FF	$0800	2 KB internal RAM
// $0800–$0FFF	$0800	Mirrors of $0000–$07FF
// $1000–$17FF	$0800
// $1800–$1FFF	$0800
// $2000–$2007	$0008	NES PPU registers
// $2008–$3FFF	$1FF8	Mirrors of $2000–$2007 (repeats every 8 bytes)
// $4000–$4017	$0018	NES APU and I/O registers
// $4018–$401F	$0008	APU and I/O functionality that is normally disabled. See CPU Test Mode.
// $4020–$FFFF
// • $6000–$7FFF
// • $8000–$FFFF	$BFE0
// $2000
// $8000	Unmapped. Available for cartridge use.
// Usually cartridge RAM, when present.
// Usually cartridge ROM and mapper registers.
// PPU Bytes
// PPUCTRL	$2000	VPHB SINN	W	NMI enable (V), PPU master/slave (P), sprite height (H), background tile select (B), sprite tile select (S), increment mode (I), nametable select / X and Y scroll bit 8 (NN)
// PPUMASK	$2001	BGRs bMmG	W	color emphasis (BGR), sprite enable (s), background enable (b), sprite left column enable (M), background left column enable (m), greyscale (G)
// PPUSTATUS	$2002	VSO- ----	R	vblank (V), sprite 0 hit (S), sprite overflow (O); read resets write pair for $2005/$2006
// OAMADDR	$2003	AAAA AAAA	W	OAM read/write address
// OAMDATA	$2004	DDDD DDDD	RW	OAM data read/write
// PPUSCROLL	$2005	XXXX XXXX YYYY YYYY	Wx2	X and Y scroll bits 7-0 (two writes: X scroll, then Y scroll)
// PPUADDR	$2006	..AA AAAA AAAA AAAA	Wx2	VRAM address (two writes: most significant byte, then least significant byte)
// PPUDATA	$2007	DDDD DDDD	RW	VRAM data read/write
// OAMDMA	$4014	AAAA AAAA	W	OAM DMA high address
func (self *Cpu) ReadAddressByte(start uint16) uint8 {

	if start == 0x2002 {
		fmt.Printf("Got PPU Memory for PPUCTRL %d", self.System.Ppu.PPUCTRL)
		return self.System.Ppu.PPUCTRL
	}

	return uint8(self.Memory[start])
}

func (self *Cpu) ReadAddress(start uint16) uint16 {
	b1 := uint16(self.Memory[start])
	b2 := uint16(self.Memory[start+1])
	fmt.Printf("Op Code %02x , B1=%02x B2=%02x\n", self.instruction, b1, b2)
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

// Push16Bit Push a unsigned 16 bit integer into the stack.
func (self *Cpu) Push16Bit(value uint16) {
	low := value & 0xff
	high := value >> 8 //push back 8 bits to the front
	self.Push(byte(high))
	self.Push(byte(low))
}

func (self *Cpu) Pull() byte {
	self.SP++
	pos := uint16(0x100) + uint16(self.SP)
	ret := self.ReadAddressByte(pos)
	return ret
}
func (self *Cpu) Pull16Bit() uint16 {
	low := uint16(self.Pull())
	high := uint16(self.Pull())
	fmt.Printf("low - %02x , high - %02x ", low, high)
	fmt.Printf(" high << %02x ", high<<4)
	ret := high<<4 | low

	return ret
}

// Check for negative & zero, common on sets +  calculations
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
	fmt.Printf("\n=====\nAbout to run instruction at %04x\n", self.PC)
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

	case Mode_AbsoluteX:
		address = self.ReadAddress(self.PC+1) + uint16(self.X)

	case Mode_AbsoluteY:
		address = self.ReadAddress(self.PC+1) + uint16(self.Y)

	case Mode_Indirect: // TODO, need to do indirect_X and Y. Contains bug
		address = self.ReadWrappedAddress(self.ReadAddress(self.PC + 1))

	case Mode_IndirectX:
		address = self.ReadWrappedAddress(self.ReadAddress(self.PC+1) + uint16(self.X))

	case Mode_IndirectY:
		address = self.ReadWrappedAddress(self.ReadAddress(self.PC+1) + uint16(self.Y))

	case Mode_Immediate:
		address = self.PC + 1

	case Mode_Accumulator:
		address = 0

	case Mode_Implied:
		address = 0

	case Mode_Relative: //Crazy one
		offset := uint16(self.Memory[self.PC+1])
		if offset < 0x80 {
			address = self.PC + 2 + offset
		} else {
			address = self.PC + 2 + offset - 0x100
		}

	case Mode_ZeroPage: //Read only one one byte refference as 16 bit
		address = uint16(self.Memory[self.PC+1])

	case Mode_ZeroPageX:
		address = uint16(uint16(self.Memory[self.PC+1]) + uint16(self.X))

	case Mode_ZeroPageY:
		address = uint16(uint16(self.Memory[self.PC+1]) + uint16(self.Y))

	}
	fmt.Printf("Got Address %02x", address)

	//Moving Increnement PC before??
	self.PC += uint16(self.info.No_Bytes)
	self.address = address

	if self.address == 0x2007 || self.address == 0x2006 {
		log.Fatalln("Interacting with VRM which I haven't coded shit")
	}
	//Run Operation
	self.info.RunOperation(self)
	self.CycleCount += uint64(self.info.No_Cycles)
	self.InstructionCount++
	fmt.Printf("Op Executed - Cycle Count = %d. Instruction Count: %d \n", self.CycleCount, self.InstructionCount)
	fmt.Println("Done with this op.")
	fmt.Println("=====")

	if self.Debug && self.InstructionCount%self.DebugLines == 0 {
		Pause()
	}

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
	self.SP = 0xfD
	self.X = 0
	self.Y = 0

}
func (self *Cpu) EmulateCycle() {
	self.DecodeInstruction()

}

// Adc Add Memory to Accumulator with Carry
// Must set Carry and Overflow Flag
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

// And is AND Memory with Accumulator
// A AND M -> A
func And(self *Cpu) {
	fmt.Println("Running Op And")
	m := self.ReadAddressByte(self.address)
	self.A &= m
	self.CheckNZ(self.A)

}

// Asl  Shift Left One Bit (Memory or Accumulator)
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

// BEQ Branch on Result Zero
func Beq(self *Cpu) {
	fmt.Println("Running Op Beq")
	if self.GetFlag(Status_Z) == true {
		self.PC = self.address
	}
}

// TODO OH GOD WHAT IS THIS!
func Bit(self *Cpu) {
	log.Fatalln("Missing Op Code")
	fmt.Println("Running Op Bit")
}

// BMI  Branch on Result Minus
func Bmi(self *Cpu) {
	fmt.Println("Running Op Bmi")
	if self.GetFlag(Status_N) == true {
		self.PC = self.address
	}
}

// BNE  Branch on Result not Zero
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

// // BRK - Force Interupt
// Chat GPT's description
// The BRK (Break) instruction on the NES (Nintendo Entertainment System) CPU is used to signal an interrupt and halt the execution of the current program.
// The BRK instruction consists of a single byte opcode (0x00) and is typically used for software debugging purposes.
// When the CPU encounters the BRK instruction, it performs the following steps:
// Push the address of the next instruction on the stack.
// Set the interrupt flag (I flag) in the status register to prevent any further interrupts.
// Push the status register on the stack.
// Load the interrupt vector at address 0xFFFE and jump to that address.
// The interrupt vector at 0xFFFE is a 16-bit address that points to the location of the interrupt service routine (ISR) for the BRK instruction.
// The ISR typically contains code that handles the software debugging operations and then returns control to the main program.
// Once the ISR is finished executing, it will load the program counter with the address that was pushed onto the stack during step 1, then restore the status register from the stack during step 3.
// The CPU will then continue executing instructions from the program counter that was loaded, as if the BRK instruction never happened.
// In summary, the NES BRK instruction is used to signal an interrupt and halt the execution of the current program, allowing the CPU to perform a software debugging operation before returning to the main program.
// push PC+2, push SR
func Brk(self *Cpu) {
	fmt.Println("Running Op Brk, cpu info  - ")
	fmt.Printf("%+v", self.info)
	self.Push16Bit(self.PC)
	Php(self)
	Sei(self)
	self.PC = self.ReadAddress(0xFFFE)

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

// Clear Decimal Flag
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
	log.Fatalln("Missing Op Code")
	fmt.Println("Running Op Cmp")
}
func Cpx(self *Cpu) {
	log.Fatalln("Missing Op Code")
	fmt.Println("Running Op Cpx")
}
func Cpy(self *Cpu) {
	log.Fatalln("Missing Op Code")
	fmt.Println("Running Op Cpy")
}

// DEC  Decrement Memory by One
func Dec(self *Cpu) {
	fmt.Println("Running Op Dec")
	m := self.ReadAddressByte(self.address)
	m--
	self.WriteMemory(self.address, m)
	self.CheckNZ(m)
}

// DEX  Decrement Index X by One
func Dex(self *Cpu) {
	fmt.Println("Running Op Dex")
	self.X--
	self.CheckNZ(self.X)
}

// DEY  Decrement Index Y by One
func Dey(self *Cpu) {
	fmt.Println("Running Op Dey")
	self.Y--
	self.CheckNZ(self.Y)
}

// EOR , Excluse OR
// A EOR M -> A
func Eor(self *Cpu) {
	fmt.Println("Running Op Eor")
	m := self.ReadAddressByte(self.address)
	self.A ^= m
	self.CheckNZ(self.A)
}

// INC  Increment Memory by One
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

// INY  Increment Y Reg by One
func Iny(self *Cpu) {
	fmt.Println("Running Op Iny")
	self.Y++
	self.CheckNZ(self.Y)
}

// JMP  Jump to new location
func Jmp(self *Cpu) {
	fmt.Println("Running Op Jmp")
	self.PC = self.address

}

// JSR  Jump and Store Current Position“
func Jsr(self *Cpu) {
	fmt.Println("Running Op Jsr")
	self.Push16Bit(self.PC - 1) //-1 because uses 3 bytes so it moves PC +3 already.
	self.PC = self.address
}

// Load memory (M) from Address (self.address) into Accumulator
func Lda(self *Cpu) {
	fmt.Println("Running Op Lda")
	self.A = self.ReadAddressByte(self.address)
	fmt.Printf("Setting Accumulator to value stored in self.address %x .. %d (hex val of %x) \n ", self.address, self.A, self.A)
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
	if self.info.Mode == Mode_Accumulator { // Interact with self.A
		carry := self.A & 1 //Firt Bit, captured before hand and saved into carry flag.
		if carry > 0 {
			self.SetFlag(Status_C, true)
		} else {
			self.SetFlag(Status_C, false)
		}
		self.A = self.A >> 1
		self.CheckNZ(self.A)
	} else { // Interact with memory read byte, read first, then modify, and write back
		m := self.ReadAddressByte(self.address)
		fmt.Printf("Read %v, contians %v, result should be %v", self.address, m, m>>1)
		carry := m & 1
		if carry > 0 {
			self.SetFlag(Status_C, true)
		} else {
			self.SetFlag(Status_C, false)
		}
		m = m >> 1
		self.WriteMemory(self.address, m)
		self.CheckNZ(m)
	}
}
func Nop(self *Cpu) {
	fmt.Println("Running Op Nop")
	fmt.Println("Doing nothing. Fucken Spot on Darryl.")
}

// ORA, Or memory with accumulator
// A OR M -> A
func Ora(self *Cpu) {
	fmt.Println("Running Op Ora")
	m := self.ReadAddressByte(self.address)
	self.A |= m
	self.CheckNZ(self.A)
}

// PHA  Push Accumulator on Stack
func Pha(self *Cpu) {
	fmt.Println("Running Op Pha")
	self.Push(self.A)
}

// PHP -  Push Processor Status on Stack
func Php(self *Cpu) {
	fmt.Println("Running Op Php")
	self.Push(self.S)
}
func Pla(self *Cpu) {
	fmt.Println("Running Op Pla")
	self.A = self.Pull()
}
func Plp(self *Cpu) {
	fmt.Println("Running Op Plp")
	self.S = self.Pull()
}
func Rol(self *Cpu) {
	log.Fatalln("Missing Op Code")
	fmt.Println("Running Op Rol")
}
func Ror(self *Cpu) {
	log.Fatalln("Missing Op Code")
	fmt.Println("Running Op Ror")
}

// RTI - Return from Interupt
// pull SR, pull PC
func Rti(self *Cpu) {
	fmt.Println("Running Op Rti")
	self.S = self.Pull()
	self.PC = self.Pull16Bit()
}

// RTS  Return from Subroutine
// pull PC, PC+1 -> PC
func Rts(self *Cpu) {
	fmt.Println("Running Op Rts")
	self.PC = self.Pull16Bit() + 1

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

// Set Status Flag of C - Carry Flag to on (00 10 00 00)
func Sec(self *Cpu) { //Set Carry Flag
	fmt.Println("Running Op Sec")
	self.SetFlag(Status_C, true)
}

// Set Status Flag of D - Decimal Flag to on (00 00 01 00) | NOT USED IN NES
func Sed(self *Cpu) {
	fmt.Println("Running Op Sed")
	self.SetFlag(Status_D, true)
}

// Set Status Flag of I - Interupt Disable to on (00 00 01 00)
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

// TAX Transfer Accumulator into index X
func Tax(self *Cpu) {
	fmt.Println("Running Op Tax")
	self.X = self.A
	self.CheckNZ(self.X)
}

// TAY Transfer Accumulator into index Y
func Tay(self *Cpu) {
	fmt.Println("Running Op Tay")
	self.Y = self.A
	self.CheckNZ(self.Y)
}

// TSX Transfer Stack Pointer into Index X
func Tsx(self *Cpu) {
	fmt.Printf("Running Op Tsx - Copying sp:%d to x:%d", self.SP, self.X)
	self.X = self.SP
	self.CheckNZ(self.Y)
}

// TXA  Transfer Index X to Accumulator
func Txa(self *Cpu) {
	fmt.Printf("Running Op Txa - copying x: %d to a: %d\n", self.X, self.A)
	self.A = self.X
	self.CheckNZ(self.A)
}

// TXS  Transfer Index X to Stack Register
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
