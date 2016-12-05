package nes

const (
	// _ = iota
	// /7 6 5 4 3 2 1 0
	//  N V   B D I Z C
	// 0-C=Carry
	// 1 Z=Zero Flag
	// 2 I=Interupt Disable
	// 3 D=Decimal
	// 4 B=Brk/software interupt
	// 5
	// 6 V Overflow Flag
	// 7 N Negative Flag
	Status_C = iota //Carry Flag
	Status_Z        //Zero Flag
	Status_I        //Interupt Disable
	Status_D        //Decimale Flag, no system need but coded anyway
	Status_B        //If break executed, causing an IRQ
	_               //Ignored
	Status_V        //Overflow
	Status_N        //Negative flag (1 for negative numb)
)
