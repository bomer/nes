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
		t.Errorf("Mario not loaded!!")
	}
}

// 0x78
func TestSei(t *testing.T) {
	Setup()

	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x78
	if nes.Cpu.S != 0 {
		t.Errorf("Status flag not init'd correctly")
	}
	nes.Cpu.EmulateCycle()

	if nes.Cpu.S != 4 {
		t.Errorf(" I Status flag not SET correctly")
	}

	//Now Clear it
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x58
	nes.Cpu.EmulateCycle()

	if nes.Cpu.S != 0 {
		t.Errorf("I Status flag not RESET correctly")
	}
}

// 0xF8 && 0xD8 - Set & Clear Decimale flag
func TestSedAdnCld(t *testing.T) {
	Setup()
	//Set decimal flag
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xF8
	if nes.Cpu.S != 0 {
		t.Errorf("Status flag not init'd correctly")
	}
	nes.Cpu.EmulateCycle()
	if fmt.Sprintf("%b", nes.Cpu.S) != "1000" {
		t.Errorf("D Status flag not setting Correctly!")
	}

	// Now Clear Decimale Flag
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xD8

	nes.Cpu.EmulateCycle()
	// fmt.Printf("3--- Should go back %s.", fmt.Sprintf("d", nes.Cpu.S))
	fmt.Printf("!!! Should go back %s ZEND", fmt.Sprintf("%d", nes.Cpu.S))
	if fmt.Sprintf("%b", nes.Cpu.S) != "0" {
		t.Errorf("D Status flag not returning back to 0")
	}
}

// 0x38 && 0x18 - Set & Clear Carry flag
func TestSecAdnClc(t *testing.T) {
	Setup()
	//Set clear flag
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x38
	if nes.Cpu.S != 0 { //inital check
		t.Errorf("Carry flag not init'd correctly")
	}
	nes.Cpu.EmulateCycle()
	if fmt.Sprintf("%b", nes.Cpu.S) != "1" {
		t.Errorf("C Carry flag not setting Correctly!")
	}

	// Now Clear Carry Flag
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x18
	nes.Cpu.EmulateCycle()
	if fmt.Sprintf("%b", nes.Cpu.S) != "0" {
		t.Errorf("D Carry flag not returning back to 0")
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
		t.Errorf("Overflow flag not setting Correctly!")
	}
	//Second scenrio for safety sake
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xB8

	nes.Cpu.S = 255
	nes.Cpu.EmulateCycle()
	if fmt.Sprintf("%b", nes.Cpu.S) != "10111111" {
		t.Errorf("Overflow flag not setting Correctly!")
	}
}

// 0xa9 - Load Accumlator into A. THis test is a precise read from the rom reset section
func TestLda(t *testing.T) {
	Setup()
	nes.Cpu.PC = 32770
	nes.Cpu.EmulateCycle()

	if nes.Cpu.A != 16 {
		t.Errorf("Failed to load into Acuumulator Value correctly")
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
		t.Errorf("Failed to load into Register X Value correctly")
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
		t.Errorf("Failed to load into Register Y Value correctly")
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
		t.Errorf("Failed to load Memory with Accumulator Value correctly")
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
		t.Errorf("Failed to load Memory with Accumulator Value correctly")
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
		t.Errorf("Failed to load Memory with Accumulator Value correctly")
	}
}

// 0xAA- Store Accumulator in Memory
func TestTax(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xAA
	nes.Cpu.A = 222
	if nes.Cpu.A != 222 {
		t.Errorf("Failed to setup A correctly")
	}
	if nes.Cpu.GetFlag(Nes.Status_N) == true {
		t.Errorf("Bady setup Flags!")
	}
	nes.Cpu.EmulateCycle()
	if nes.Cpu.X != 222 {
		t.Errorf("Failed to Copy A -> X correctly")
	}

	if nes.Cpu.GetFlag(Nes.Status_N) != true {
		t.Errorf("Bady calculated Flags!")
	}
}

// 0xA8- Transfer Accumulator to Y
func TestTaY(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xA8
	nes.Cpu.A = 222
	if nes.Cpu.A != 222 {
		t.Errorf("Failed to setup A correctly")
	}
	if nes.Cpu.GetFlag(Nes.Status_N) == true {
		t.Errorf("Bady setup Flags!")
	}
	nes.Cpu.EmulateCycle()
	if nes.Cpu.Y != 222 {
		t.Errorf("Failed to Copy A -> X correctly")
	}

	if nes.Cpu.GetFlag(Nes.Status_N) != true {
		t.Errorf("Bady calculated Flags!")
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
		t.Errorf("Failed to setup Y correctly")
	}
	if nes.Cpu.SP != 222 {
		t.Errorf("U DONE BAD")
	}
	if nes.Cpu.GetFlag(Nes.Status_N) != false {
		t.Errorf("Bady setup Flags!")
	}

	//Run and test
	nes.Cpu.EmulateCycle()
	if nes.Cpu.X != 222 {
		t.Errorf("Failed to Copy SP -> Y correctly")
	}

	if nes.Cpu.GetFlag(Nes.Status_N) != false {
		t.Errorf("Bady calculated Flags!")
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
		t.Errorf("Failed to setup Accumulator correctly")
	}
	if nes.Cpu.X != 222 {
		t.Errorf("U DONE BAD")
	}
	if nes.Cpu.GetFlag(Nes.Status_N) != false {
		t.Errorf("Bady setup Flags!")
	}

	//Run and test
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 222 {
		t.Errorf("Failed to Copy X to A correctly")
	}

	if nes.Cpu.GetFlag(Nes.Status_N) != true {
		t.Errorf("Bady calculated Flags!")
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
		t.Errorf("Failed to setup Accumulator correctly")
	}
	if nes.Cpu.X != 222 {
		t.Errorf("U DONE BAD")
	}
	if nes.Cpu.GetFlag(Nes.Status_N) != false {
		t.Errorf("Bady setup Flags!")
	}

	//Run and test
	nes.Cpu.EmulateCycle()
	if nes.Cpu.SP != 222 {
		t.Errorf("Failed to Copy X to A correctly")
	}

	if nes.Cpu.GetFlag(Nes.Status_N) != true {
		t.Errorf("Bady calculated Flags!")
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
		t.Errorf("Failed to setup Accumulator correctly")
	}
	if nes.Cpu.Y != 222 {
		t.Errorf("U DONE BAD")
	}
	if nes.Cpu.GetFlag(Nes.Status_N) != false {
		t.Errorf("Bady setup Flags!")
	}

	//Run and test
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 222 {
		t.Errorf("Failed to Copy X to A correctly")
	}

	if nes.Cpu.GetFlag(Nes.Status_N) != true {
		t.Errorf("Bady calculated Flags!")
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
		t.Errorf("Failed to add 50 and 100")
	}
	if nes.Cpu.GetFlag(Nes.Status_V) != true {
		t.Errorf("Failed to add 50 and 100, the overflow flag came back wrong")
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != false {
		t.Errorf("Failed to add 50 and 100, the Carry flag came back wrong")
	}

	//Second test, 1  + 1 + c of 1, a=3, overflow is false, carry false
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xab] = 1
	nes.Cpu.A = 1
	nes.Cpu.SetFlag(Nes.Status_C, true)

	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 3 {
		t.Errorf("Failed to add 1 and 1 and C of 1")
	}
	if nes.Cpu.GetFlag(Nes.Status_V) != false {
		t.Errorf("Failed to add 1 and 1 and C of 1, the overflow flag came back wrong")
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != false {
		t.Errorf("Failed to add 1 and 1 and C of 1, the Carry flag came back wrong")
	}

	//Second test, 100  + 100, should equal 200..
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xab] = 200
	nes.Cpu.A = 200

	nes.Cpu.EmulateCycle()
	fmt.Println(nes.Cpu.A)
	if nes.Cpu.A != 144 {
		t.Errorf("Failed to add 200 and 200")
	}
	if nes.Cpu.GetFlag(Nes.Status_V) != false {
		t.Errorf("Failed to add 200 and 200, the overflow flag came back wrong")
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != true {
		t.Errorf("Failed to add 200 and 200, the Carry flag came back wrong")
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
		t.Errorf("Failed to minus 100 and 50")
	}
	if nes.Cpu.GetFlag(Nes.Status_V) != false {
		t.Errorf("Failed to minus 100 and 50, the overflow flag came back wrong")
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != true {
		t.Errorf("Failed to minus 100 and 50, the Carry flag came back wrong")
	}

	//200-200=0
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xab] = 200
	nes.Cpu.A = 200
	nes.Cpu.SetFlag(Nes.Status_C, true)
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 0 {
		t.Errorf("Failed to minus 100 and 50")
	}
	if nes.Cpu.GetFlag(Nes.Status_V) != false {
		t.Errorf("Failed to minus 100 and 50, the overflow flag came back wrong")
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != true {
		t.Errorf("Failed to minus 100 and 50, the Carry flag came back wrong")
	}
	if nes.Cpu.GetFlag(Nes.Status_Z) != true {
		t.Errorf("Failed to minus 100 and 50, the Carry flag came back wrong")
	}

	// 0-200=56
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xab] = 200
	nes.Cpu.A = 0
	nes.Cpu.SetFlag(Nes.Status_C, true)
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 56 {
		t.Errorf("Failed to minus 100 and 50, got instead%d", nes.Cpu.A)
	}

	// -50 - 50 =-100
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xab] = 50
	nes.Cpu.A = 40
	nes.Cpu.SetFlag(Nes.Status_C, false)
	// nes.Cpu.SetFlag(Nes.Status_N, true)
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 245 {
		t.Errorf("Failed to minus 100 and 50, got instead%d", nes.Cpu.A)
	}

}

//0x29 , a = AND memory and A
func TestAND(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0x29
	nes.Cpu.Memory[0x29] = 0x29
	nes.Cpu.Memory[0x2a] = 20
	nes.Cpu.A = 4
	nes.Cpu.SetFlag(Nes.Status_C, true)

	//First test, 50  + 100, a=150, overflow is true, carry false
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 4 {
		t.Errorf("Failed to and 20 and 4 = 4")
	}

	//second, 1111 &  0111, 15 7
	nes.Cpu.PC = 0x29
	nes.Cpu.Memory[0x29] = 0x29
	nes.Cpu.Memory[0x2a] = 7
	nes.Cpu.A = 15
	nes.Cpu.SetFlag(Nes.Status_C, true)

	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 7 {
		t.Errorf("Failed to and 15 and 7 = 7")
	}
}

//0x09 , a = OR memory and A
func TestORA(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x09
	nes.Cpu.Memory[0xab] = 20
	nes.Cpu.A = 4

	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 20 {
		t.Errorf("Failed to OR 20 and 4 = 20")
	}

	//second, 50 | 7
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x09
	nes.Cpu.Memory[0xab] = 7
	nes.Cpu.A = 50

	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 55 {
		t.Errorf("Failed to or 15 and 7 = 7, got %d", nes.Cpu.A)
	}
}

//0x49 , a = OR memory and A
func TestEOR(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x49
	nes.Cpu.Memory[0xab] = 20
	nes.Cpu.A = 4

	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 16 {
		t.Errorf("Failed to OR 20 and 4 = 20, got %d", nes.Cpu.A)
	}

	//second, 50 | 7
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x49
	nes.Cpu.Memory[0xab] = 1
	nes.Cpu.A = 32

	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 33 {
		t.Errorf("Failed to or 15 and 7 = 7, got %d", nes.Cpu.A)
	}

	//second, 50 | 7
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x49
	nes.Cpu.Memory[0xab] = 33
	nes.Cpu.A = 32

	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 1 {
		t.Errorf("Failed to or 15 and 7 = 7, got %d", nes.Cpu.A)
	}
}

//0x0A, ASL, shift left
func TestASL(t *testing.T) {
	Setup()

	//First Test, Accumuator shift 64 to 128
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x0A
	nes.Cpu.A = 64
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 128 {
		t.Errorf("Failed, did niot shift 64 to 128")
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != false {
		t.Errorf("Failed, carry wrong on shift 64 to 128")
	}

	//Second Test, Accumuator shift 192 shift
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x0A
	nes.Cpu.A = 192
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 128 {
		t.Errorf("Failed, did niot shift 192 to 128")
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != true {
		t.Errorf("Failed, carry wrong on shift 192 to 128")
	}

	//Now Memory modification ops = 0E
	//First Test, Accumuator shift 64 to 128
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[64] = 64
	nes.Cpu.Memory[0xaa] = 0x0E
	nes.Cpu.Memory[0xab] = 64
	nes.Cpu.Memory[0xac] = 0

	nes.Cpu.EmulateCycle()
	fmt.Printf("mem 0xab = %v", nes.Cpu.Memory[0xab])
	if nes.Cpu.Memory[64] != 128 {
		t.Errorf("Failed, did niot shift 64 to 128")
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != false {
		t.Errorf("Failed, carry wrong on shift 64 to 128")
	}

	//Now Memory modification ops = 0E
	//First Test, Accumuator shift 64 to 128
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[64] = 192
	nes.Cpu.Memory[0xaa] = 0x0E
	nes.Cpu.Memory[0xab] = 64
	nes.Cpu.Memory[0xac] = 0

	nes.Cpu.EmulateCycle()
	fmt.Printf("mem 0xab = %v", nes.Cpu.Memory[0xab])
	if nes.Cpu.Memory[64] != 128 {
		t.Errorf("Failed, did niot shift 192 to 128")
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != true {
		t.Errorf("Failed, carry wrong on shift 64 to 128")
	}

}

//0x90 Branch on Carry Clear
func TestBCC(t *testing.T) {
	Setup()
	//0x50 (80), increment by 8.
	nes.Cpu.PC = 80
	nes.Cpu.Memory[0x50] = 0x90
	nes.Cpu.Memory[0x51] = 8
	nes.Cpu.SetFlag(Nes.Status_C, false)

	nes.Cpu.EmulateCycle()

	if nes.Cpu.PC != 90 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to branch PC by offset of 8, got %d", nes.Cpu.PC)
	}

	//0x50 (80), increment by 8, but with carry on, so should only go +2
	nes.Cpu.PC = 80
	nes.Cpu.Memory[0x50] = 0x90
	nes.Cpu.Memory[0x51] = 8
	nes.Cpu.SetFlag(Nes.Status_C, true)

	nes.Cpu.EmulateCycle()

	if nes.Cpu.PC != 82 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to branch PC by offset of 8, got %d", nes.Cpu.PC)
	}

	//0x50 (80), increment by 200, negative number (-72)
	nes.Cpu.PC = 80
	nes.Cpu.Memory[0x50] = 0x90
	nes.Cpu.Memory[0x51] = 200
	nes.Cpu.SetFlag(Nes.Status_C, false)

	nes.Cpu.EmulateCycle()
	if nes.Cpu.PC != 26 { //80 -72(=8) + 2=10 (always add to to PC, pass or fail)
		t.Errorf("Failed to branch PC by offset of -72, got %d", nes.Cpu.PC)
	}
	//0x50 (80), DONT with negative number
	nes.Cpu.PC = 80
	nes.Cpu.Memory[0x50] = 0x90
	nes.Cpu.Memory[0x51] = 200
	nes.Cpu.SetFlag(Nes.Status_C, true)

	nes.Cpu.EmulateCycle()
	if nes.Cpu.PC != 82 { //80 -72(=8) + 2=10 (always add to to PC, pass or fail)
		t.Errorf("Failed to branch PC negative, got %d", nes.Cpu.PC)
	}
}

//0xB0 Branch on Carry Set
func TestBCS(t *testing.T) {
	Setup()
	//0x50 (80), increment by 8.
	nes.Cpu.PC = 80
	nes.Cpu.Memory[0x50] = 0xB0
	nes.Cpu.Memory[0x51] = 8
	nes.Cpu.SetFlag(Nes.Status_C, true)

	nes.Cpu.EmulateCycle()

	if nes.Cpu.PC != 90 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to branch PC by offset of 8, got %d", nes.Cpu.PC)
	}

	//0x50 (80), increment by 8, but with carry on, so should only go +2
	nes.Cpu.PC = 80
	nes.Cpu.Memory[0x50] = 0xB0
	nes.Cpu.Memory[0x51] = 8
	nes.Cpu.SetFlag(Nes.Status_C, false)

	nes.Cpu.EmulateCycle()

	if nes.Cpu.PC != 82 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to branch PC by offset of 8, got %d", nes.Cpu.PC)
	}

	//0x50 Move backwards, with negative number (-72, 200 in signed range)
	nes.Cpu.PC = 80
	nes.Cpu.Memory[0x50] = 0xB0
	nes.Cpu.Memory[0x51] = 200
	nes.Cpu.SetFlag(Nes.Status_C, true)

	nes.Cpu.EmulateCycle()
	if nes.Cpu.PC != 26 { //80 -72(=8) + 2=10 (always add to to PC, pass or fail)
		t.Errorf("Failed to branch PC by offset of -72, got %d", nes.Cpu.PC)
	}

	//0x50 (80), Dont increment by 200, negative number (-72), only go 2 up
	nes.Cpu.PC = 80
	nes.Cpu.Memory[0x50] = 0xB0
	nes.Cpu.Memory[0x51] = 200
	nes.Cpu.SetFlag(Nes.Status_C, false)

	nes.Cpu.EmulateCycle()
	if nes.Cpu.PC != 82 { //80 -72(=8) + 2=10 (always add to to PC, pass or fail)
		t.Errorf("Failed to branch PC by offset of -72, got %d", nes.Cpu.PC)
	}
}

//0xB0 Branch on Carry Set
func TestBEQ(t *testing.T) {
	Setup()
	//0x50 (80), increment by 8.
	nes.Cpu.PC = 80
	nes.Cpu.Memory[0x50] = 0xF0
	nes.Cpu.Memory[0x51] = 8
	nes.Cpu.SetFlag(Nes.Status_Z, true)

	nes.Cpu.EmulateCycle()

	if nes.Cpu.PC != 90 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to branch PC by offset of 8, got %d", nes.Cpu.PC)
	}
}

//0x30 BMI  Branch on Result Minus
func TestBMI(t *testing.T) {
	Setup()
	//0x50 (80), increment by 8.
	nes.Cpu.PC = 80
	nes.Cpu.Memory[0x50] = 0x30
	nes.Cpu.Memory[0x51] = 8
	nes.Cpu.SetFlag(Nes.Status_N, true)

	nes.Cpu.EmulateCycle()

	if nes.Cpu.PC != 90 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to branch PC by offset of 8, got %d", nes.Cpu.PC)
	}
}

//0xD0 BNE  Branch on Result not Zero
func TestBNE(t *testing.T) {
	Setup()
	//0x50 (80), increment by 8.
	nes.Cpu.PC = 80
	nes.Cpu.Memory[0x50] = 0xD0
	nes.Cpu.Memory[0x51] = 8
	nes.Cpu.SetFlag(Nes.Status_N, true)

	nes.Cpu.EmulateCycle()

	if nes.Cpu.PC != 90 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to branch PC by offset of 8, got %d", nes.Cpu.PC)
	}
}

//0x10 BPL  Branch on Result Plus
func TestBPL(t *testing.T) {
	Setup()
	//0x50 (80), increment by 8.
	nes.Cpu.PC = 80
	nes.Cpu.Memory[0x50] = 0x10
	nes.Cpu.Memory[0x51] = 8
	nes.Cpu.SetFlag(Nes.Status_N, false)

	nes.Cpu.EmulateCycle()

	if nes.Cpu.PC != 90 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to branch PC by offset of 8, got %d", nes.Cpu.PC)
	}
}

//0x50 BVC  Branch on Overflow Clear
func TestBVC(t *testing.T) {
	Setup()
	//0x50 (80), increment by 8.
	nes.Cpu.PC = 80
	nes.Cpu.Memory[0x50] = 0x50
	nes.Cpu.Memory[0x51] = 8
	nes.Cpu.SetFlag(Nes.Status_V, false)

	nes.Cpu.EmulateCycle()

	if nes.Cpu.PC != 90 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to branch PC by offset of 8, got %d", nes.Cpu.PC)
	}
}

//0x70 BVS  Branch on Overflow Set
func TestBvs(t *testing.T) {
	Setup()
	//0x50 (80), increment by 8.
	nes.Cpu.PC = 80
	nes.Cpu.Memory[0x50] = 0x70
	nes.Cpu.Memory[0x51] = 8
	nes.Cpu.SetFlag(Nes.Status_V, true)

	nes.Cpu.EmulateCycle()

	if nes.Cpu.PC != 90 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to branch PC by offset of 8, got %d", nes.Cpu.PC)
	}
}

//C6 DEC  Decrement Memory by One
func TestDEC(t *testing.T) {
	Setup()
	//0x50 (80), increment by 8.
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xC6
	nes.Cpu.Memory[0xab] = 100
	nes.Cpu.Memory[100] = 100
	nes.Cpu.SetFlag(Nes.Status_V, true)

	nes.Cpu.EmulateCycle()

	if nes.Cpu.Memory[100] != 99 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to DEC memory, got %d", nes.Cpu.Memory[100])
	}
}

//C6 DEX  Decrement X by One
func TestDEX(t *testing.T) {
	Setup()
	//0x50 (80), increment by 8.
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xCA
	nes.Cpu.X = 100
	nes.Cpu.SetFlag(Nes.Status_V, true)

	nes.Cpu.EmulateCycle()

	if nes.Cpu.X != 99 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to DEC memory, got %d", nes.Cpu.X)
	}
}

//88 DEX  Decrement Y by One
func TestDEY(t *testing.T) {
	Setup()
	//0x50 (80), increment by 8.
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x88
	nes.Cpu.Y = 100

	nes.Cpu.EmulateCycle()

	if nes.Cpu.Y != 99 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to DEC memory, got %d", nes.Cpu.Y)
	}
}

//0xE6 INC  Increment Memory by One
func TestINC(t *testing.T) {
	Setup()
	//Memory goes from 100 to 101
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xE6
	nes.Cpu.Memory[0xab] = 100
	nes.Cpu.Memory[100] = 100
	nes.Cpu.SetFlag(Nes.Status_V, true)

	nes.Cpu.EmulateCycle()

	if nes.Cpu.Memory[100] != 101 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to INC memory, got %d", nes.Cpu.Memory[100])
	}
}

//0xE8 INX  Increment X by One
func TestINX(t *testing.T) {
	Setup()
	//Memory goes from 100 to 101
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xE8
	nes.Cpu.X = 100
	nes.Cpu.SetFlag(Nes.Status_V, true)

	nes.Cpu.EmulateCycle()

	if nes.Cpu.X != 101 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to INC memory, got %d", nes.Cpu.X)
	}
}

//0xC8 INY  Increment Y by One
func TestINY(t *testing.T) {
	Setup()
	//Memory goes from 100 to 101
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xC8
	nes.Cpu.Y = 100

	nes.Cpu.EmulateCycle()

	if nes.Cpu.Y != 101 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to INC memory, got %d", nes.Cpu.Y)
	}
}

//0x48, PHA, push A. and 0x68, PULL A
func TestPHAandPLA(t *testing.T) {
	Setup()
	//Push 3 numbers into stack, then pop them off. 50, 60, 70. then, pop back
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x48
	nes.Cpu.Memory[0xab] = 0x48
	nes.Cpu.Memory[0xac] = 0x48

	nes.Cpu.Memory[0xad] = 0x68
	nes.Cpu.Memory[0xae] = 0x68
	nes.Cpu.Memory[0xaf] = 0x68

	nes.Cpu.A = 50
	nes.Cpu.EmulateCycle()
	nes.Cpu.A = 60
	nes.Cpu.EmulateCycle()
	nes.Cpu.A = 70
	nes.Cpu.EmulateCycle()
	nes.Cpu.A = 80 //set to 80, but dont push it into stack

	if nes.Cpu.A != 80 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Stack setup wrong")
	}
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 70 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to Pull memory 7, got %d", nes.Cpu.A)
	}
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 60 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to Pull memory 6, got %d", nes.Cpu.A)
	}
	nes.Cpu.EmulateCycle()
	if nes.Cpu.A != 50 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to Pull memory 5, got %d", nes.Cpu.A)
	}
}

//0x08, PHP, push A. and 0x28, PULL A
func TestPHPandPLP(t *testing.T) {
	Setup()
	//Push 3 numbers into stack, then pop them off. 50, 60, 70. then, pop back
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x08
	nes.Cpu.Memory[0xab] = 0x08
	nes.Cpu.Memory[0xac] = 0x08

	nes.Cpu.Memory[0xad] = 0x28
	nes.Cpu.Memory[0xae] = 0x28
	nes.Cpu.Memory[0xaf] = 0x28

	nes.Cpu.S = 50
	nes.Cpu.EmulateCycle()
	nes.Cpu.S = 60
	nes.Cpu.EmulateCycle()
	nes.Cpu.S = 70
	nes.Cpu.EmulateCycle()
	nes.Cpu.S = 80 //set to 80, but dont push it into stack

	if nes.Cpu.S != 80 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Stack setup wrong")
	}
	nes.Cpu.EmulateCycle()
	if nes.Cpu.S != 70 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to Pull memory 7, got %d", nes.Cpu.S)
	}
	nes.Cpu.EmulateCycle()
	if nes.Cpu.S != 60 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to Pull memory 6, got %d", nes.Cpu.S)
	}
	nes.Cpu.EmulateCycle()
	if nes.Cpu.S != 50 { //80 + 8 + 2 (always add to to PC, pass or fail)
		t.Errorf("Failed to Pull memory 5, got %d", nes.Cpu.S)
	}
}

//0x4A,  Logical Shift Right
func TestLSR(t *testing.T) {
	Setup()
	//Test 13 becomes...6 1101 > 0110 (carry on)
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x4A
	nes.Cpu.A = 13

	nes.Cpu.EmulateCycle()

	if nes.Cpu.A != 6 {
		t.Errorf("Failed to Shift A right, got %d", nes.Cpu.A)
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != true {
		t.Errorf("Failed to Shift A right, wrong carry set")
	}
	//Test 2, 8 becomes 4, 1000 > 0100 (carry off)
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x46
	nes.Cpu.Memory[0xab] = 0xac
	nes.Cpu.Memory[0xac] = 8

	nes.Cpu.EmulateCycle()

	if nes.Cpu.Memory[0xac] != 4 {
		t.Errorf("Failed to Shift A right, got %d", nes.Cpu.A)
	}
	if nes.Cpu.GetFlag(Nes.Status_C) != false {
		t.Errorf("Failed to Shift A right, wrong carry set")
	}
}

// 0x4C, JUMP
func TestJmpJsrAndPull(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x4C
	nes.Cpu.Memory[0xab] = 0xba
	nes.Cpu.Memory[0xac] = 0x00

	nes.Cpu.EmulateCycle()

	if nes.Cpu.PC != 0xba { //possible should be 0xbb.. not sure, if wrong, minus 2 from PC after setting
		t.Errorf("Failed to JMP to 0xb gota instead %x", nes.Cpu.PC)
		fmt.Printf("in hex %02x", nes.Cpu.PC)
	}

	//Now test JSR, store and jump 0x20
	nes.Cpu.Memory[0xbd] = 0x20
	nes.Cpu.Memory[0xbe] = 0xaa
	nes.Cpu.Memory[0xbf] = 0x00
	nes.Cpu.EmulateCycle()
	if nes.Cpu.PC != 0xfff0 {
		t.Errorf("Failed to JMP to 0xbb got instead %x", nes.Cpu.PC)
		fmt.Printf("in hex %02x", nes.Cpu.PC)
	}
	//Now Jump back to bd - RTS 0x60
	nes.Cpu.Memory[0xad] = 0x60
	nes.Cpu.EmulateCycle()
	if nes.Cpu.PC != 0xfff3 {
		t.Errorf("Failed to RTS to 0xbb got instead %x", nes.Cpu.PC)
		fmt.Printf("in hex %02x", nes.Cpu.PC)
	}

}

// 0xEA, NOP DO NOTHING
func TestNOP(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0xEA

	nes.Cpu.EmulateCycle()
	if nes.Cpu.PC != 0xAB {
		t.Errorf("You failed to do NOTHING. good job. No wonder the op code is EA.")
	}
}

//0x00 BRK and RTI
// Dec 2021, Moving back PC check, after moving 
func TestBRK(t *testing.T) {
	Setup()
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x00
	nes.Cpu.S = 11

	nes.Cpu.EmulateCycle()

	fmt.Printf("Checking %02x %02x ", nes.Cpu.Memory[0xFFFF], nes.Cpu.Memory[0xFFFE])

	if nes.Cpu.PC != 0xFFF0 {
		t.Errorf("Should moved to 0xFFF1, instead was %02x", nes.Cpu.PC)
	}

	//NOW RTI
	nes.Cpu.PC = 0xaa
	nes.Cpu.Memory[0xaa] = 0x40
	nes.Cpu.S = 0

	nes.Cpu.EmulateCycle()
	if nes.Cpu.S != 11 {
		t.Error("Did not restore S", nes.Cpu.PC)
	}

	if nes.Cpu.PC != 0xab {
		t.Errorf("Should moved to 0xab, instead was %02x", nes.Cpu.PC)
	}
}
