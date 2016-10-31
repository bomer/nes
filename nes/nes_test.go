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

func TestSei(t *testing.T) {
	// Mode_Immediate
	Setup()
	nes.Cpu.PC = 32768
	nes.Cpu.EmulateCycle()

	if nes.Cpu.Memory[35000] != 132 {
		t.Error("Mario not loaded!!")
	}
}
