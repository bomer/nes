package nes

// addressing modes
const (
	_ = iota
	Mode_Immediate
	Mode_Absolute

	Mode_AbsoluteX
	Mode_AbsoluteY
	Mode_Accumulator
	Mode_Implied
	Mode_IndexedIndirect
	Mode_Indirect
	Mode_IndirectIndexed
	Mode_Relative
	Mode_ZeroPage
	Mode_ZeroPageX
	Mode_ZeroPageY
)
