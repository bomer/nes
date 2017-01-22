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
	Setup()
	fmt.Printf("AAaaand%+v", nes.Cpu.Memory[35000])
	if nes.Cpu.Memory[35000] != 132 {
		t.Error("Mario not loaded!!")
	}
}

// 0x78
func TestSei(t *testing.T) {
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
	Setup()
	nes.Cpu.PC = 32770
	nes.Cpu.EmulateCycle()

	if nes.Cpu.A != 16 {
		t.Error("Failed to load into Acuumulator Value correctly")
	}
}

// 0xa9 - Load Accumlator into X . Testing just in immediate so the next value is used, aa will read ab
func TestLdx(t *testing.T) {
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

// 0xAA- Store Accumulator in Memory
func TestTax(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xAA
	nes.Cpu.A = 222
	if nes.Cpu.A != 222 {
		t.Error("Failed to setup A correctly")
	}
	if nes.Cpu.GetFlag(Nes.Status_N) == true {
		t.Error("Bady setup Flags!")
	}
	nes.Cpu.EmulateCycle()
	if nes.Cpu.X != 222 {
		t.Error("Failed to Copy A -> X correctly")
	}

	if nes.Cpu.GetFlag(Nes.Status_N) != true {
		t.Error("Bady calculated Flags!")
	}
}

// 0xA8- Transfer Accumulator to Y
func TestTaY(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xA8
	nes.Cpu.A = 222
	if nes.Cpu.A != 222 {
		t.Error("Failed to setup A correctly")
	}
	if nes.Cpu.GetFlag(Nes.Status_N) == true {
		t.Error("Bady setup Flags!")
	}
	nes.Cpu.EmulateCycle()
	if nes.Cpu.Y != 222 {
		t.Error("Failed to Copy A -> X correctly")
	}

	if nes.Cpu.GetFlag(Nes.Status_N) != true {
		t.Error("Bady calculated Flags!")
	}
}

// 0xBA- Transfer Accumulator to Y
func TestTSX(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xBA
	nes.Cpu.SP = 222
	nes.Cpu.Y = 5
	//before checl
	if nes.Cpu.Y != 5 {
		t.Error("Failed to setup Y correctly")
	}
	if nes.Cpu.SP != 222 {
		t.Error("U DONE BAD")
	}
	if nes.Cpu.GetFlag(Nes.Status_N) != false {
		t.Error("Bady setup Flags!")
	}

	//Run and test
	nes.Cpu.EmulateCycle()
	if nes.Cpu.X != 222 {
		t.Error("Failed to Copy SP -> Y correctly")
	}

	if nes.Cpu.GetFlag(Nes.Status_N) != false {
		t.Error("Bady calculated Flags!")
	}
}

// 0x8A Transfer Index X to Accumulator
func TestTXA(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x8A
	nes.Cpu.X = 222
	nes.Cpu.A = 5
	//before checl
	if nes.Cpu.A != 5 {
		t.Error("Failed to setup Accumulator correctly")
	}
	if nes.Cpu.X != 222 {
		t.Error("U DONE BAD")
	}
	if nes.Cpu.GetFlag(Nes.Status_N) != false {
		t.Error("Bady setup Flags!")
	}

	//Run and test
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 222 {
		t.Error("Failed to Copy X to A correctly")
	}

	if nes.Cpu.GetFlag(Nes.Status_N) != true {
		t.Error("Bady calculated Flags!")
	}
}

// 0x9a TXS  Transfer Index X to Stack Register
func TestTXS(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x9A
	nes.Cpu.X = 222
	nes.Cpu.SP = 5
	//before checl
	if nes.Cpu.SP != 5 {
		t.Error("Failed to setup Accumulator correctly")
	}
	if nes.Cpu.X != 222 {
		t.Error("U DONE BAD")
	}
	if nes.Cpu.GetFlag(Nes.Status_N) != false {
		t.Error("Bady setup Flags!")
	}

	//Run and test
	nes.Cpu.EmulateCycle()
	if nes.Cpu.SP != 222 {
		t.Error("Failed to Copy X to A correctly")
	}

	if nes.Cpu.GetFlag(Nes.Status_N) != true {
		t.Error("Bady calculated Flags!")
	}
}

// 0x98 TYA  Transfer Index Y to Accumulator
func TestTYA(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x98
	nes.Cpu.Y = 222
	nes.Cpu.A = 5
	//before checl
	if nes.Cpu.A != 5 {
		t.Error("Failed to setup Accumulator correctly")
	}
	if nes.Cpu.Y != 222 {
		t.Error("U DONE BAD")
	}
	if nes.Cpu.GetFlag(Nes.Status_N) != false {
		t.Error("Bady setup Flags!")
	}

	//Run and test
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 222 {
		t.Error("Failed to Copy X to A correctly")
	}

	if nes.Cpu.GetFlag(Nes.Status_N) != true {
		t.Error("Bady calculated Flags!")
	}
}

//0x69 ADC Add M To accumulator + C
func TestADC(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x69
	nes.Cpu.Memory[0xab] = 100
	nes.Cpu.A = 50

	//First test, 50  + 100, a=150, overflow is true, carry false
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 150 {
		t.Error("Failed to add 50 and 100")
	}
	if nes.Cpu.GetFlag(Nes.Status_V) != true {
		t.Error("Failed to add 50 and 100, the overflow flag came back wrong")
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != false {
		t.Error("Failed to add 50 and 100, the Carry flag came back wrong")
	}

	//Second test, 1  + 1 + c of 1, a=3, overflow is false, carry false
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xab] = 1
	nes.Cpu.A = 1
	nes.Cpu.SetFlag(Nes.Status_C, true)

	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 3 {
		t.Error("Failed to add 1 and 1 and C of 1")
	}
	if nes.Cpu.GetFlag(Nes.Status_V) != false {
		t.Error("Failed to add 1 and 1 and C of 1, the overflow flag came back wrong")
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != false {
		t.Error("Failed to add 1 and 1 and C of 1, the Carry flag came back wrong")
	}

	//Second test, 100  + 100, should equal 200..
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xab] = 200
	nes.Cpu.A = 200

	nes.Cpu.EmulateCycle()
	fmt.Println(nes.Cpu.A)
	if nes.Cpu.A != 144 {
		t.Error("Failed to add 200 and 200")
	}
	if nes.Cpu.GetFlag(Nes.Status_V) != false {
		t.Error("Failed to add 200 and 200, the overflow flag came back wrong")
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != true {
		t.Error("Failed to add 200 and 200, the Carry flag came back wrong")
	}

}

//0xe9 SBC subtract M from accumulator + 1-C
func TestSBC(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xE9
	nes.Cpu.Memory[0xab] = 50
	nes.Cpu.A = 100
	nes.Cpu.SetFlag(Nes.Status_C, true)

	//First test, 50  + 100, a=150, overflow is true, carry false
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 50 {
		t.Error("Failed to minus 100 and 50")
	}
	if nes.Cpu.GetFlag(Nes.Status_V) != false {
		t.Error("Failed to minus 100 and 50, the overflow flag came back wrong")
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != true {
		t.Error("Failed to minus 100 and 50, the Carry flag came back wrong")
	}

	//200-200=0
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xab] = 200
	nes.Cpu.A = 200
	nes.Cpu.SetFlag(Nes.Status_C, true)
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 0 {
		t.Error("Failed to minus 100 and 50")
	}
	if nes.Cpu.GetFlag(Nes.Status_V) != false {
		t.Error("Failed to minus 100 and 50, the overflow flag came back wrong")
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != true {
		t.Error("Failed to minus 100 and 50, the Carry flag came back wrong")
	}
	if nes.Cpu.GetFlag(Nes.Status_Z) != true {
		t.Error("Failed to minus 100 and 50, the Carry flag came back wrong")
	}

	// 0-200=56
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xab] = 200
	nes.Cpu.A = 0
	nes.Cpu.SetFlag(Nes.Status_C, true)
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 56 {
		t.Error("Failed to minus 100 and 50, got instead%d", nes.Cpu.A)
	}

	// -50 - 50 =-100
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xab] = 50
	nes.Cpu.A = 40
	nes.Cpu.SetFlag(Nes.Status_C, false)
	// nes.Cpu.SetFlag(Nes.Status_N, true)
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 245 {
		t.Error("Failed to minus 100 and 50, got instead%d", nes.Cpu.A)
	}

}
