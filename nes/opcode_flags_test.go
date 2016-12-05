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
	nes.Cpu.S = 0
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
		t.Error(" I Status flag not SET correctly")
	}

	//Now Clear it
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x58
	nes.Cpu.EmulateCycle()

	if nes.Cpu.S != 0 {
		t.Error("I Status flag not RESET correctly")
	}
}

// 0xF8 && 0xD8 - Set & Clear Decimale flag
func TestSedAdnCld(t *testing.T) {
	// Mode_Immediate
	Setup()
	//Set decimal flag
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xF8
	if nes.Cpu.S != 0 {
		t.Error("Status flag not init'd correctly")
	}
	nes.Cpu.EmulateCycle()
	if fmt.Sprintf("%b", nes.Cpu.S) != "1000" {
		t.Error("D Status flag not setting Correctly!")
	}

	// Now Clear Decimale Flag
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xD8

	nes.Cpu.EmulateCycle()
	// fmt.Printf("3--- Should go back %s.", fmt.Sprintf("d", nes.Cpu.S))
	fmt.Printf("!!! Should go back %s ZEND", fmt.Sprintf("d", nes.Cpu.S))
	if fmt.Sprintf("%b", nes.Cpu.S) != "0" {
		t.Error("D Status flag not returning back to 0")
	}
}

// 0x38 && 0x18 - Set & Clear Carry flag
func TestSecAdnClc(t *testing.T) {
	Setup()
	//Set clear flag
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x38
	if nes.Cpu.S != 0 { //inital check
		t.Error("Carry flag not init'd correctly")
	}
	nes.Cpu.EmulateCycle()
	if fmt.Sprintf("%b", nes.Cpu.S) != "1" {
		t.Error("C Carry flag not setting Correctly!")
	}

	// Now Clear Carry Flag
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x18
	nes.Cpu.EmulateCycle()
	if fmt.Sprintf("%b", nes.Cpu.S) != "0" {
		t.Error("D Carry flag not returning back to 0")
	}
}

// 0xB8- Clear Overflow Flag, so set because that only happens during ops running
func TestClv(t *testing.T) {
	Setup()
	//Set clear flag
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xB8

	nes.Cpu.S = 64
	nes.Cpu.EmulateCycle()
	if fmt.Sprintf("%b", nes.Cpu.S) != "0" {
		t.Error("Overflow flag not setting Correctly!")
	}
	//Second scenrio for safety sake
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xB8

	nes.Cpu.S = 255
	nes.Cpu.EmulateCycle()
	if fmt.Sprintf("%b", nes.Cpu.S) != "10111111" {
		t.Error("Overflow flag not setting Correctly!")
	}
}
