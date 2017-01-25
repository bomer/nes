package nes

// import "fmt"

// addressing modes
const (
	_ = iota
	Mode_Immediate
	Mode_Absolute

	Mode_AbsoluteX
	Mode_AbsoluteY

	Mode_Accumulator
	Mode_Implied
	Mode_Indirect
	Mode_IndirectX
	Mode_IndirectY
	Mode_Relative
	Mode_ZeroPage
	Mode_ZeroPageX
	Mode_ZeroPageY
)

const (
	_ = iota
	ADC
	AND
	ASL
	BCC
	BCS
	BEQ
	BIT
	BMI
	BNE
	BPL
	BRK
	BVC
	BVS
	CLC
	CLD
	CLI
	CLV
	CMP
	CPX
	CPY
	DEC
	DEX
	DEY
	EOR
	INC
	INX
	INY
	JMP
	JSR
	LDA
	LDX
	LDY
	LSR
	NOP
	ORA
	PHA
	PHP
	PLA
	PLP
	ROL
	ROR
	RTI
	RTS
	SBC
	SEC
	SED
	SEI
	STA
	STX
	STY
	TAX
	TAY
	TSX
	TXA
	TXS
	TYA
)

type fn func(cpu *Cpu) // fn = operation function

type OpCodeInfo struct {
	Mode          int
	Operation     int
	No_Bytes      byte
	No_Cycles     byte
	BoundaryCross byte
	Function      fn
}

func (o *OpCodeInfo) RunOperation(cpu *Cpu) {
	o.Function(cpu)
}

func (o *OpCodeInfo) ModeString() string {
	name := []string{"_", "Mode_Immediate", "Mode_Absolute", "Mode_AbsoluteX", "Mode_AbsoluteY", "Mode_Accumulator", "Mode_Implied", "Mode_Indirect", "Mode_IndirectX", "Mode_IndirectY", "Mode_Relative", "Mode_ZeroPage", "Mode_ZeroPageX", "Mode_ZeroPageY"}
	return name[o.Mode]
}

func (o *OpCodeInfo) OperationString() string {
	name := []string{"_", "ADC", "AND", "ASL", "BCC", "BCS", "BEQ", "BIT", "BMI", "BNE", "BPL", "BRK", "BVC", "BVS", "CLC", "CLD", "CLI", "CLV", "CMP", "CPX", "CPY", "DEC", "DEX", "DEY", "EOR", "INC", "INX", "INY", "JMP", "JSR", "LDA", "LDX", "LDY", "LSR", "NOP", "ORA", "PHA", "PHP", "PLA", "PLP", "ROL", "ROR", "RTI", "RTS", "SBC", "SEC", "SED", "SEI", "STA", "STX", "STY", "TAX", "TAY", "TSX", "TXA", "TXS", "TYA"}
	return name[o.Operation]
}

var OpTable = map[int]OpCodeInfo{
	// 0x00: OpCodeInfo{Mode_Implied, BRK, 1, 7, 0},

	0x00: OpCodeInfo{Mode_Implied, BRK, 1, 7, 0, Brk},
	0x01: OpCodeInfo{Mode_IndirectX, ORA, 2, 6, 0, Ora},
	0x05: OpCodeInfo{Mode_ZeroPage, ORA, 2, 3, 0, Ora},
	0x06: OpCodeInfo{Mode_ZeroPage, ASL, 2, 5, 0, Asl},
	0x08: OpCodeInfo{Mode_Implied, PHP, 1, 3, 0, Php},
	0x09: OpCodeInfo{Mode_Immediate, ORA, 2, 2, 0, Ora},
	0x10: OpCodeInfo{Mode_Relative, BPL, 2, 2, 1, Bpl},
	0x11: OpCodeInfo{Mode_IndirectY, ORA, 2, 5, 1, Ora},
	0x15: OpCodeInfo{Mode_ZeroPageX, ORA, 2, 4, 0, Ora},
	0x16: OpCodeInfo{Mode_ZeroPageX, ASL, 2, 6, 0, Asl},
	0x18: OpCodeInfo{Mode_Implied, CLC, 1, 2, 0, Clc},
	0x19: OpCodeInfo{Mode_AbsoluteY, ORA, 3, 4, 1, Ora},
	0x20: OpCodeInfo{Mode_Absolute, JSR, 3, 6, 0, Jsr},
	0x21: OpCodeInfo{Mode_IndirectX, AND, 2, 6, 0, And},
	0x24: OpCodeInfo{Mode_ZeroPage, BIT, 2, 3, 0, Bit},
	0x25: OpCodeInfo{Mode_ZeroPage, AND, 2, 3, 0, And},
	0x26: OpCodeInfo{Mode_ZeroPage, ROL, 2, 5, 0, Rol},
	0x28: OpCodeInfo{Mode_Implied, PHP, 1, 4, 0, Php},
	0x29: OpCodeInfo{Mode_Immediate, AND, 2, 2, 0, And},
	0x30: OpCodeInfo{Mode_Relative, BMI, 2, 2, 1, Bmi},
	0x31: OpCodeInfo{Mode_IndirectY, AND, 2, 5, 1, And},
	0x35: OpCodeInfo{Mode_ZeroPageX, AND, 2, 4, 0, And},
	0x36: OpCodeInfo{Mode_ZeroPageX, ROL, 2, 6, 0, Rol},
	0x38: OpCodeInfo{Mode_Implied, SEC, 1, 2, 0, Sec},
	0x39: OpCodeInfo{Mode_AbsoluteY, AND, 3, 4, 1, And},
	0x40: OpCodeInfo{Mode_Implied, RTI, 1, 6, 0, Rti},
	0x41: OpCodeInfo{Mode_IndirectX, EOR, 2, 6, 0, Eor},
	0x45: OpCodeInfo{Mode_ZeroPage, EOR, 2, 3, 0, Eor},
	0x46: OpCodeInfo{Mode_ZeroPage, LSR, 2, 5, 0, Lsr},
	0x48: OpCodeInfo{Mode_Implied, PHA, 1, 3, 0, Pha},
	0x49: OpCodeInfo{Mode_Immediate, EOR, 2, 2, 0, Eor},
	0x50: OpCodeInfo{Mode_Relative, BVC, 2, 2, 1, Bvc},
	0x51: OpCodeInfo{Mode_IndirectY, EOR, 2, 5, 1, Eor},
	0x55: OpCodeInfo{Mode_ZeroPageX, EOR, 2, 4, 0, Eor},
	0x56: OpCodeInfo{Mode_ZeroPageX, LSR, 2, 6, 0, Lsr},
	0x58: OpCodeInfo{Mode_Implied, CLI, 1, 2, 0, Cli},
	0x59: OpCodeInfo{Mode_AbsoluteY, EOR, 3, 4, 1, Eor},
	0x60: OpCodeInfo{Mode_Implied, RTS, 1, 6, 0, Rts},
	0x61: OpCodeInfo{Mode_IndirectX, ADC, 2, 6, 0, Adc},
	0x65: OpCodeInfo{Mode_ZeroPage, ADC, 2, 3, 0, Adc},
	0x66: OpCodeInfo{Mode_ZeroPage, ROR, 2, 5, 0, Ror},
	0x68: OpCodeInfo{Mode_Implied, PLA, 1, 4, 0, Pla},
	0x69: OpCodeInfo{Mode_Immediate, ADC, 2, 2, 0, Adc},
	0x70: OpCodeInfo{Mode_Relative, BVS, 2, 2, 1, Bvs},
	0x71: OpCodeInfo{Mode_IndirectY, ADC, 2, 5, 1, Adc},
	0x75: OpCodeInfo{Mode_ZeroPageX, ADC, 2, 4, 0, Adc},
	0x76: OpCodeInfo{Mode_ZeroPageX, ROR, 2, 6, 0, Ror},
	0x78: OpCodeInfo{Mode_Implied, SEI, 1, 2, 0, Sei},
	0x79: OpCodeInfo{Mode_AbsoluteY, ADC, 3, 4, 1, Adc},
	0x81: OpCodeInfo{Mode_IndirectX, STA, 2, 6, 0, Sta},
	0x84: OpCodeInfo{Mode_ZeroPage, STY, 2, 3, 0, Sty},
	0x85: OpCodeInfo{Mode_ZeroPage, STA, 2, 3, 0, Sta},
	0x86: OpCodeInfo{Mode_ZeroPage, STX, 2, 3, 0, Stx},
	0x88: OpCodeInfo{Mode_Implied, DEY, 1, 2, 0, Dey},
	0x90: OpCodeInfo{Mode_Relative, BCC, 2, 2, 1, Bcc},
	0x91: OpCodeInfo{Mode_IndirectY, STA, 2, 6, 0, Sta},
	0x94: OpCodeInfo{Mode_ZeroPageX, STY, 2, 4, 0, Sty},
	0x95: OpCodeInfo{Mode_ZeroPageX, STA, 2, 4, 0, Sta},
	0x96: OpCodeInfo{Mode_ZeroPageY, STX, 2, 4, 0, Stx},
	0x98: OpCodeInfo{Mode_Implied, TYA, 1, 2, 0, Tya},
	0x99: OpCodeInfo{Mode_AbsoluteY, STA, 3, 5, 0, Sta},
	0x0A: OpCodeInfo{Mode_Accumulator, ASL, 1, 2, 0, Asl},
	0x0D: OpCodeInfo{Mode_Absolute, ORA, 3, 4, 0, Ora},
	0x0E: OpCodeInfo{Mode_Absolute, ASL, 3, 6, 0, Asl},
	0x1D: OpCodeInfo{Mode_AbsoluteX, ORA, 3, 4, 1, Ora},
	0x1E: OpCodeInfo{Mode_AbsoluteX, ASL, 3, 7, 0, Asl},
	0x2A: OpCodeInfo{Mode_Accumulator, ROL, 1, 2, 0, Rol},
	0x2C: OpCodeInfo{Mode_Absolute, BIT, 3, 4, 0, Bit},
	0x2D: OpCodeInfo{Mode_Absolute, AND, 3, 4, 0, And},
	0x2E: OpCodeInfo{Mode_Absolute, ROL, 3, 6, 0, Rol},
	0x3D: OpCodeInfo{Mode_AbsoluteX, AND, 3, 4, 1, And},
	0x3E: OpCodeInfo{Mode_AbsoluteX, ROL, 3, 7, 0, Rol},
	0x4A: OpCodeInfo{Mode_Accumulator, LSR, 1, 2, 0, Lsr},
	0x4C: OpCodeInfo{Mode_Absolute, JMP, 3, 3, 0, Jmp},
	0x4D: OpCodeInfo{Mode_Absolute, EOR, 3, 4, 0, Eor},
	0x4E: OpCodeInfo{Mode_Absolute, LSR, 3, 6, 0, Lsr},
	0x5D: OpCodeInfo{Mode_AbsoluteX, EOR, 3, 4, 1, Eor},
	0x5E: OpCodeInfo{Mode_AbsoluteX, LSR, 3, 7, 0, Lsr},
	0x6A: OpCodeInfo{Mode_Accumulator, ROR, 1, 2, 0, Ror},
	0x6C: OpCodeInfo{Mode_Indirect, JMP, 3, 5, 0, Jmp},
	0x6D: OpCodeInfo{Mode_Absolute, ADC, 3, 4, 0, Adc},
	0x6E: OpCodeInfo{Mode_Absolute, ROR, 3, 6, 0, Ror},
	0x7D: OpCodeInfo{Mode_AbsoluteX, ADC, 3, 4, 1, Adc},
	0x7E: OpCodeInfo{Mode_AbsoluteX, ROR, 3, 7, 0, Ror},
	0x8A: OpCodeInfo{Mode_Implied, TXA, 1, 2, 0, Txa},
	0x8C: OpCodeInfo{Mode_Absolute, STY, 3, 4, 0, Sty},
	0x8D: OpCodeInfo{Mode_Absolute, STA, 3, 4, 0, Sta},
	0x8E: OpCodeInfo{Mode_Absolute, STX, 3, 4, 0, Stx},
	0x9A: OpCodeInfo{Mode_Implied, TXS, 1, 2, 0, Txs},
	0x9D: OpCodeInfo{Mode_AbsoluteX, STA, 3, 5, 0, Sta},
	0xA0: OpCodeInfo{Mode_Immediate, LDY, 2, 2, 0, Ldy},
	0xA1: OpCodeInfo{Mode_IndirectX, LDA, 2, 6, 0, Lda},
	0xA2: OpCodeInfo{Mode_Immediate, LDX, 2, 2, 0, Ldx},
	0xA4: OpCodeInfo{Mode_ZeroPage, LDY, 2, 3, 0, Ldy},
	0xA5: OpCodeInfo{Mode_ZeroPage, LDA, 2, 3, 0, Lda},
	0xA6: OpCodeInfo{Mode_ZeroPage, LDX, 2, 3, 0, Ldx},
	0xA8: OpCodeInfo{Mode_Implied, TAY, 1, 2, 0, Tay},
	0xA9: OpCodeInfo{Mode_Immediate, LDA, 2, 2, 0, Lda},
	0xAA: OpCodeInfo{Mode_Implied, TAX, 1, 2, 0, Tax},
	0xAC: OpCodeInfo{Mode_Absolute, LDY, 3, 4, 0, Ldy},
	0xAD: OpCodeInfo{Mode_Absolute, LDA, 3, 4, 0, Lda},
	0xAE: OpCodeInfo{Mode_Absolute, LDX, 3, 4, 0, Ldx},
	0xB0: OpCodeInfo{Mode_Relative, BCS, 2, 2, 1, Bcs},
	0xB1: OpCodeInfo{Mode_IndirectY, LDA, 2, 5, 1, Lda},
	0xB4: OpCodeInfo{Mode_ZeroPageX, LDY, 2, 4, 0, Ldy},
	0xB5: OpCodeInfo{Mode_ZeroPageX, LDA, 2, 4, 0, Lda},
	0xB6: OpCodeInfo{Mode_ZeroPageY, LDX, 2, 4, 0, Ldx},
	0xB8: OpCodeInfo{Mode_Implied, CLV, 1, 2, 0, Clv},
	0xB9: OpCodeInfo{Mode_AbsoluteY, LDA, 3, 4, 1, Lda},
	0xBA: OpCodeInfo{Mode_Implied, TSX, 1, 2, 0, Tsx},
	0xBC: OpCodeInfo{Mode_AbsoluteX, LDY, 3, 4, 1, Ldy},
	0xBD: OpCodeInfo{Mode_AbsoluteX, LDA, 3, 4, 1, Lda},
	0xBE: OpCodeInfo{Mode_AbsoluteY, LDX, 3, 4, 1, Ldx},
	0xC0: OpCodeInfo{Mode_Immediate, CPY, 2, 2, 0, Cpy},
	0xC1: OpCodeInfo{Mode_IndirectX, CMP, 2, 6, 0, Cmp},
	0xC4: OpCodeInfo{Mode_ZeroPage, CPY, 2, 3, 0, Cpy},
	0xC5: OpCodeInfo{Mode_ZeroPage, CMP, 2, 3, 0, Cmp},
	0xC6: OpCodeInfo{Mode_ZeroPage, DEC, 2, 5, 0, Dec},
	0xC8: OpCodeInfo{Mode_Implied, INY, 1, 2, 0, Iny},
	0xC9: OpCodeInfo{Mode_Immediate, CMP, 2, 2, 0, Cmp},
	0xCA: OpCodeInfo{Mode_Implied, DEX, 1, 2, 0, Dex},
	0xCC: OpCodeInfo{Mode_Absolute, CPY, 3, 4, 0, Cpy},
	0xCD: OpCodeInfo{Mode_Absolute, CMP, 3, 4, 0, Cmp},
	0xCE: OpCodeInfo{Mode_Absolute, DEC, 3, 3, 0, Dec},
	0xD0: OpCodeInfo{Mode_Relative, BNE, 2, 2, 1, Bne},
	0xD1: OpCodeInfo{Mode_IndirectY, CMP, 2, 5, 1, Cmp},
	0xD5: OpCodeInfo{Mode_ZeroPageX, CMP, 2, 4, 0, Cmp},
	0xD6: OpCodeInfo{Mode_ZeroPageX, DEC, 2, 6, 0, Dec},
	0xD8: OpCodeInfo{Mode_Implied, CLD, 1, 2, 0, Cld},
	0xD9: OpCodeInfo{Mode_AbsoluteY, CMP, 3, 4, 1, Cmp},
	0xDD: OpCodeInfo{Mode_AbsoluteX, CMP, 3, 4, 1, Cmp},
	0xDE: OpCodeInfo{Mode_AbsoluteX, DEC, 3, 7, 0, Dec},
	0xE0: OpCodeInfo{Mode_Immediate, CPX, 2, 2, 0, Cpx},
	0xE1: OpCodeInfo{Mode_IndirectX, SBC, 2, 6, 0, Sbc},
	0xE4: OpCodeInfo{Mode_ZeroPage, CPX, 2, 3, 0, Cpx},
	0xE5: OpCodeInfo{Mode_ZeroPage, SBC, 2, 3, 0, Sbc},
	0xE6: OpCodeInfo{Mode_ZeroPage, INC, 2, 5, 0, Inc},
	0xE8: OpCodeInfo{Mode_Implied, INX, 1, 2, 0, Inx},
	0xE9: OpCodeInfo{Mode_Immediate, SBC, 2, 2, 0, Sbc},
	0xEA: OpCodeInfo{Mode_Implied, NOP, 1, 2, 0, Nop},
	0xEC: OpCodeInfo{Mode_Absolute, CPX, 3, 4, 0, Cpx},
	0xED: OpCodeInfo{Mode_Absolute, SBC, 3, 4, 0, Sbc},
	0xEE: OpCodeInfo{Mode_Absolute, INC, 3, 6, 0, Inc},
	0xF0: OpCodeInfo{Mode_Relative, BEQ, 2, 2, 1, Beq},
	0xF1: OpCodeInfo{Mode_IndirectY, SBC, 2, 5, 1, Sbc},
	0xF5: OpCodeInfo{Mode_ZeroPageX, SBC, 2, 4, 0, Sbc},
	0xF6: OpCodeInfo{Mode_ZeroPageX, INC, 2, 6, 0, Inc},
	0xF8: OpCodeInfo{Mode_Implied, SED, 1, 2, 0, Sed},
	0xF9: OpCodeInfo{Mode_AbsoluteY, SBC, 3, 4, 1, Sbc},
	0xFD: OpCodeInfo{Mode_AbsoluteX, SBC, 3, 4, 1, Sbc},
	0xFE: OpCodeInfo{Mode_AbsoluteX, INC, 3, 7, 0, Inc},
}
