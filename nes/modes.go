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

	0x00: {Mode_Implied, BRK, 1, 7, 0, Brk},
	0x01: {Mode_IndirectX, ORA, 2, 6, 0, Ora},
	0x05: {Mode_ZeroPage, ORA, 2, 3, 0, Ora},
	0x06: {Mode_ZeroPage, ASL, 2, 5, 0, Asl},
	0x08: {Mode_Implied, PHP, 1, 3, 0, Php},
	0x09: {Mode_Immediate, ORA, 2, 2, 0, Ora},
	0x10: {Mode_Relative, BPL, 2, 2, 1, Bpl},
	0x11: {Mode_IndirectY, ORA, 2, 5, 1, Ora},
	0x15: {Mode_ZeroPageX, ORA, 2, 4, 0, Ora},
	0x16: {Mode_ZeroPageX, ASL, 2, 6, 0, Asl},
	0x18: {Mode_Implied, CLC, 1, 2, 0, Clc},
	0x19: {Mode_AbsoluteY, ORA, 3, 4, 1, Ora},
	0x20: {Mode_Absolute, JSR, 3, 6, 0, Jsr},
	0x21: {Mode_IndirectX, AND, 2, 6, 0, And},
	0x24: {Mode_ZeroPage, BIT, 2, 3, 0, Bit},
	0x25: {Mode_ZeroPage, AND, 2, 3, 0, And},
	0x26: {Mode_ZeroPage, ROL, 2, 5, 0, Rol},
	0x28: {Mode_Implied, PLP, 1, 4, 0, Plp},
	0x29: {Mode_Immediate, AND, 2, 2, 0, And},
	0x30: {Mode_Relative, BMI, 2, 2, 1, Bmi},
	0x31: {Mode_IndirectY, AND, 2, 5, 1, And},
	0x35: {Mode_ZeroPageX, AND, 2, 4, 0, And},
	0x36: {Mode_ZeroPageX, ROL, 2, 6, 0, Rol},
	0x38: {Mode_Implied, SEC, 1, 2, 0, Sec},
	0x39: {Mode_AbsoluteY, AND, 3, 4, 1, And},
	0x40: {Mode_Implied, RTI, 1, 6, 0, Rti},
	0x41: {Mode_IndirectX, EOR, 2, 6, 0, Eor},
	0x45: {Mode_ZeroPage, EOR, 2, 3, 0, Eor},
	0x46: {Mode_ZeroPage, LSR, 2, 5, 0, Lsr},
	0x48: {Mode_Implied, PHA, 1, 3, 0, Pha},
	0x49: {Mode_Immediate, EOR, 2, 2, 0, Eor},
	0x50: {Mode_Relative, BVC, 2, 2, 1, Bvc},
	0x51: {Mode_IndirectY, EOR, 2, 5, 1, Eor},
	0x55: {Mode_ZeroPageX, EOR, 2, 4, 0, Eor},
	0x56: {Mode_ZeroPageX, LSR, 2, 6, 0, Lsr},
	0x58: {Mode_Implied, CLI, 1, 2, 0, Cli},
	0x59: {Mode_AbsoluteY, EOR, 3, 4, 1, Eor},
	0x60: {Mode_Implied, RTS, 1, 6, 0, Rts},
	0x61: {Mode_IndirectX, ADC, 2, 6, 0, Adc},
	0x65: {Mode_ZeroPage, ADC, 2, 3, 0, Adc},
	0x66: {Mode_ZeroPage, ROR, 2, 5, 0, Ror},
	0x68: {Mode_Implied, PLA, 1, 4, 0, Pla},
	0x69: {Mode_Immediate, ADC, 2, 2, 0, Adc},
	0x70: {Mode_Relative, BVS, 2, 2, 1, Bvs},
	0x71: {Mode_IndirectY, ADC, 2, 5, 1, Adc},
	0x75: {Mode_ZeroPageX, ADC, 2, 4, 0, Adc},
	0x76: {Mode_ZeroPageX, ROR, 2, 6, 0, Ror},
	0x78: {Mode_Implied, SEI, 1, 2, 0, Sei},
	0x79: {Mode_AbsoluteY, ADC, 3, 4, 1, Adc},
	0x81: {Mode_IndirectX, STA, 2, 6, 0, Sta},
	0x84: {Mode_ZeroPage, STY, 2, 3, 0, Sty},
	0x85: {Mode_ZeroPage, STA, 2, 3, 0, Sta},
	0x86: {Mode_ZeroPage, STX, 2, 3, 0, Stx},
	0x88: {Mode_Implied, DEY, 1, 2, 0, Dey},
	0x90: {Mode_Relative, BCC, 2, 2, 1, Bcc},
	0x91: {Mode_IndirectY, STA, 2, 6, 0, Sta},
	0x94: {Mode_ZeroPageX, STY, 2, 4, 0, Sty},
	0x95: {Mode_ZeroPageX, STA, 2, 4, 0, Sta},
	0x96: {Mode_ZeroPageY, STX, 2, 4, 0, Stx},
	0x98: {Mode_Implied, TYA, 1, 2, 0, Tya},
	0x99: {Mode_AbsoluteY, STA, 3, 5, 0, Sta},
	0x0A: {Mode_Accumulator, ASL, 1, 2, 0, Asl},
	0x0D: {Mode_Absolute, ORA, 3, 4, 0, Ora},
	0x0E: {Mode_Absolute, ASL, 3, 6, 0, Asl},
	0x1D: {Mode_AbsoluteX, ORA, 3, 4, 1, Ora},
	0x1E: {Mode_AbsoluteX, ASL, 3, 7, 0, Asl},
	0x2A: {Mode_Accumulator, ROL, 1, 2, 0, Rol},
	0x2C: {Mode_Absolute, BIT, 3, 4, 0, Bit},
	0x2D: {Mode_Absolute, AND, 3, 4, 0, And},
	0x2E: {Mode_Absolute, ROL, 3, 6, 0, Rol},
	0x3D: {Mode_AbsoluteX, AND, 3, 4, 1, And},
	0x3E: {Mode_AbsoluteX, ROL, 3, 7, 0, Rol},
	0x4A: {Mode_Accumulator, LSR, 1, 2, 0, Lsr},
	0x4C: {Mode_Absolute, JMP, 3, 3, 0, Jmp},
	0x4D: {Mode_Absolute, EOR, 3, 4, 0, Eor},
	0x4E: {Mode_Absolute, LSR, 3, 6, 0, Lsr},
	0x5D: {Mode_AbsoluteX, EOR, 3, 4, 1, Eor},
	0x5E: {Mode_AbsoluteX, LSR, 3, 7, 0, Lsr},
	0x6A: {Mode_Accumulator, ROR, 1, 2, 0, Ror},
	0x6C: {Mode_Indirect, JMP, 3, 5, 0, Jmp},
	0x6D: {Mode_Absolute, ADC, 3, 4, 0, Adc},
	0x6E: {Mode_Absolute, ROR, 3, 6, 0, Ror},
	0x7D: {Mode_AbsoluteX, ADC, 3, 4, 1, Adc},
	0x7E: {Mode_AbsoluteX, ROR, 3, 7, 0, Ror},
	0x8A: {Mode_Implied, TXA, 1, 2, 0, Txa},
	0x8C: {Mode_Absolute, STY, 3, 4, 0, Sty},
	0x8D: {Mode_Absolute, STA, 3, 4, 0, Sta},
	0x8E: {Mode_Absolute, STX, 3, 4, 0, Stx},
	0x9A: {Mode_Implied, TXS, 1, 2, 0, Txs},
	0x9D: {Mode_AbsoluteX, STA, 3, 5, 0, Sta},
	0xA0: {Mode_Immediate, LDY, 2, 2, 0, Ldy},
	0xA1: {Mode_IndirectX, LDA, 2, 6, 0, Lda},
	0xA2: {Mode_Immediate, LDX, 2, 2, 0, Ldx},
	0xA4: {Mode_ZeroPage, LDY, 2, 3, 0, Ldy},
	0xA5: {Mode_ZeroPage, LDA, 2, 3, 0, Lda},
	0xA6: {Mode_ZeroPage, LDX, 2, 3, 0, Ldx},
	0xA8: {Mode_Implied, TAY, 1, 2, 0, Tay},
	0xA9: {Mode_Immediate, LDA, 2, 2, 0, Lda},
	0xAA: {Mode_Implied, TAX, 1, 2, 0, Tax},
	0xAC: {Mode_Absolute, LDY, 3, 4, 0, Ldy},
	0xAD: {Mode_Absolute, LDA, 3, 4, 0, Lda},
	0xAE: {Mode_Absolute, LDX, 3, 4, 0, Ldx},
	0xB0: {Mode_Relative, BCS, 2, 2, 1, Bcs},
	0xB1: {Mode_IndirectY, LDA, 2, 5, 1, Lda},
	0xB4: {Mode_ZeroPageX, LDY, 2, 4, 0, Ldy},
	0xB5: {Mode_ZeroPageX, LDA, 2, 4, 0, Lda},
	0xB6: {Mode_ZeroPageY, LDX, 2, 4, 0, Ldx},
	0xB8: {Mode_Implied, CLV, 1, 2, 0, Clv},
	0xB9: {Mode_AbsoluteY, LDA, 3, 4, 1, Lda},
	0xBA: {Mode_Implied, TSX, 1, 2, 0, Tsx},
	0xBC: {Mode_AbsoluteX, LDY, 3, 4, 1, Ldy},
	0xBD: {Mode_AbsoluteX, LDA, 3, 4, 1, Lda},
	0xBE: {Mode_AbsoluteY, LDX, 3, 4, 1, Ldx},
	0xC0: {Mode_Immediate, CPY, 2, 2, 0, Cpy},
	0xC1: {Mode_IndirectX, CMP, 2, 6, 0, Cmp},
	0xC4: {Mode_ZeroPage, CPY, 2, 3, 0, Cpy},
	0xC5: {Mode_ZeroPage, CMP, 2, 3, 0, Cmp},
	0xC6: {Mode_ZeroPage, DEC, 2, 5, 0, Dec},
	0xC8: {Mode_Implied, INY, 1, 2, 0, Iny},
	0xC9: {Mode_Immediate, CMP, 2, 2, 0, Cmp},
	0xCA: {Mode_Implied, DEX, 1, 2, 0, Dex},
	0xCC: {Mode_Absolute, CPY, 3, 4, 0, Cpy},
	0xCD: {Mode_Absolute, CMP, 3, 4, 0, Cmp},
	0xCE: {Mode_Absolute, DEC, 3, 3, 0, Dec},
	0xD0: {Mode_Relative, BNE, 2, 2, 1, Bne},
	0xD1: {Mode_IndirectY, CMP, 2, 5, 1, Cmp},
	0xD5: {Mode_ZeroPageX, CMP, 2, 4, 0, Cmp},
	0xD6: {Mode_ZeroPageX, DEC, 2, 6, 0, Dec},
	0xD8: {Mode_Implied, CLD, 1, 2, 0, Cld},
	0xD9: {Mode_AbsoluteY, CMP, 3, 4, 1, Cmp},
	0xDD: {Mode_AbsoluteX, CMP, 3, 4, 1, Cmp},
	0xDE: {Mode_AbsoluteX, DEC, 3, 7, 0, Dec},
	0xE0: {Mode_Immediate, CPX, 2, 2, 0, Cpx},
	0xE1: {Mode_IndirectX, SBC, 2, 6, 0, Sbc},
	0xE4: {Mode_ZeroPage, CPX, 2, 3, 0, Cpx},
	0xE5: {Mode_ZeroPage, SBC, 2, 3, 0, Sbc},
	0xE6: {Mode_ZeroPage, INC, 2, 5, 0, Inc},
	0xE8: {Mode_Implied, INX, 1, 2, 0, Inx},
	0xE9: {Mode_Immediate, SBC, 2, 2, 0, Sbc},
	0xEA: {Mode_Implied, NOP, 1, 2, 0, Nop},
	0xEC: {Mode_Absolute, CPX, 3, 4, 0, Cpx},
	0xED: {Mode_Absolute, SBC, 3, 4, 0, Sbc},
	0xEE: {Mode_Absolute, INC, 3, 6, 0, Inc},
	0xF0: {Mode_Relative, BEQ, 2, 2, 1, Beq},
	0xF1: {Mode_IndirectY, SBC, 2, 5, 1, Sbc},
	0xF5: {Mode_ZeroPageX, SBC, 2, 4, 0, Sbc},
	0xF6: {Mode_ZeroPageX, INC, 2, 6, 0, Inc},
	0xF8: {Mode_Implied, SED, 1, 2, 0, Sed},
	0xF9: {Mode_AbsoluteY, SBC, 3, 4, 1, Sbc},
	0xFD: {Mode_AbsoluteX, SBC, 3, 4, 1, Sbc},
	0xFE: {Mode_AbsoluteX, INC, 3, 7, 0, Inc},
}
