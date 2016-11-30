package nes

const (
	// _ = iota
	// /7 6 5 4 3 2 1 0
	//  Z V   B D I Z C
	//C=Carry, Z=Zero, I=Interupt, D=Decimal,B=Brk/software interupt, V-Overflow,S=Sign, 1=negative
	Status_C = iota //Carry Flag
	Status_Z        //Zero Flag
	Status_I        //Interupt Disable
	Status_D        //Ignored
	Status_B        //If break executed, causing an IRQ
	_               //Ignored
	Status_V        //Overflow
	Status_N        //Negative flag (1 for negative numb)
)
