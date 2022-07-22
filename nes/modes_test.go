package nes_test

import (
	"testing"

	Nes "github.com/bomer/nes/nes"
)

func TestModes(t *testing.T) {
	// Mode_Immediate
	// fmt.Printf("Mode_Absolute %+v \n", Nes.OpTable[0x00])
	info := Nes.OpTable[0x00]
	// fmt.Printf("Mode - %s, Operation - %s \n", info.ModeString(), info.OperationString())
	if info.ModeString() != "Mode_Implied" {
		t.Error("Wrong Mode Retrieved")
	}
	if info.OperationString() != "BRK" {
		t.Error("Wrong Operation Retrieved")
	}

	// 	0x61: OpCodeInfo{Mode_IndirectX, ADC, 2, 6, 0},
	info = Nes.OpTable[0x61]
	if info.ModeString() != "Mode_IndirectX" {
		t.Error("Wrong Mode Retrieved")
	}
	if info.OperationString() != "ADC" {
		t.Error("Wrong Operation Retrieved")
	}
}
