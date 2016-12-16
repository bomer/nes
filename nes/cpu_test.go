package nes_test

import (
	"fmt"
	Nes "github.com/bomer/nes/nes"
	"testing"
)

//TestNES - All unit tests are written in mariographical order.
//That is, in the order required to run Mario on my emulator
//Except in the case where there are mutliple accumulator versions, where I did those as well
var nes Nes.Nes

func Setup() {
	nes.Cpu.Quiet = true
	nes.Rom.LoadGame("../mario.nes", &nes.Cpu)
	nes.Cpu.Init()
	nes.Cpu.S = 0
	nes.Cpu.Quiet = false
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

// 0xa9 - Load Accumlator into A. THis test is a precise read from the rom reset section
func TestLda(t *testing.T) {
	// Mode_Immediate
	Setup()
	nes.Cpu.PC = 32770
	nes.Cpu.EmulateCycle()

	if nes.Cpu.A != 16 {
		t.Error("Failed to load into Acuumulator Value correctly")
	}
}

// 0xa9 - Load Accumlator into X . Testing just in immediate so the next value is used, aa will read ab
func TestLdx(t *testing.T) {
	// Mode_Immediate
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xA2
	nes.Cpu.Memory[0xab] = 222
	fmt.Printf("Checking some memory %d %d %d", nes.Cpu.Memory[0xaa], nes.Cpu.Memory[0xab], nes.Cpu.Memory[0xac])

	nes.Cpu.EmulateCycle()

	if nes.Cpu.X != 222 {
		t.Error("Failed to load into Register X Value correctly")
	}
}

// 0xa9 - Load Accumlator into T. Same as above
func TestLdy(t *testing.T) {
	// Mode_Immediate
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xA0
	nes.Cpu.Memory[0xab] = 222

	nes.Cpu.EmulateCycle()

	if nes.Cpu.Y != 222 {
		t.Error("Failed to load into Register Y Value correctly")
	}
}

// 0x8D - Store Accumulator in Memory
func TestSta(t *testing.T) {
	// Mode_Immediate
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.A = 111
	nes.Cpu.Memory[0xaa] = 0x8D
	nes.Cpu.Memory[0xab] = 222
	nes.Cpu.Memory[0xac] = 222

	nes.Cpu.EmulateCycle()

	if nes.Cpu.Memory[0xdede] != 111 {
		t.Error("Failed to load Memory with Accumulator Value correctly")
	}
}

// 0x8E - Store Accumulator in Memory
func TestStx(t *testing.T) {
	// Mode_Immediate
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.X = 111
	nes.Cpu.Memory[0xaa] = 0x8E
	nes.Cpu.Memory[0xab] = 222
	nes.Cpu.Memory[0xac] = 222

	nes.Cpu.EmulateCycle()

	if nes.Cpu.Memory[0xdede] != 111 {
		t.Error("Failed to load Memory with Accumulator Value correctly")
	}
}

// 0x8C- Store Accumulator in Memory
func TestSty(t *testing.T) {
	// Mode_Immediate
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Y = 111
	nes.Cpu.Memory[0xaa] = 0x8C
	nes.Cpu.Memory[0xab] = 222
	nes.Cpu.Memory[0xac] = 222

	nes.Cpu.EmulateCycle()

	if nes.Cpu.Memory[0xdede] != 111 {
		t.Error("Failed to load Memory with Accumulator Value correctly")
	}
}
