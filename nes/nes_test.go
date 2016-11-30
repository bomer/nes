package nes_test

import (
	"fmt"
	Nes "github.com/bomer/nes/nes"
	"testing"
)

var nes Nes.Nes

func Setup() {
	nes.Rom.LoadGame("../mario.nes", &nes.Cpu)
	nes.Cpu.Init()
}

func TestLoad(t *testing.T) {
	// Mode_Immediate
	Setup()
	fmt.Printf("AAaaand%+v", nes.Cpu.Memory[35000])
	if nes.Cpu.Memory[35000] != 132 {
		t.Error("Mario not loaded!!")
	}
}

// 0x78
func TestSei(t *testing.T) {
	// Mode_Immediate
	Setup()

	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x78
	if nes.Cpu.S != 0 {
		t.Error("Status flag not init'd correctly")
	}
	nes.Cpu.EmulateCycle()

	if nes.Cpu.S != 4 {
		t.Error(" I Status flag not Updated correctly")
	}
}
