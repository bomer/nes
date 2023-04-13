package nes

import (
	"fmt"
	"io/ioutil"
)

type Rom struct {
}

type RomHeader struct {
	PGR_banks byte
	CHR_banks byte
	Control   byte
	Control2  byte
	Ram_banks byte
}

// Where the cartridge is read in
// Read file in curent dir into Memory of NES CPU
func (self *Rom) LoadGame(filename string, nes *Nes) {
	var header RomHeader
	rom, _ := ioutil.ReadFile(filename)
	rom_length := len(rom)
	if rom_length > 0 {
		fmt.Printf("Rom Length = %d\n", rom_length)
	}

	//Mario Header
	//0 1  2 3  4 5  6 7  8 9  1011 1213 1415
	//4e45 531a 0201 0100 0000 0000 0000 0000
	//NES   |   P R
	if rom[3] != 0x1a {
		panic("OH GOD WHAT IS THIS. I don't know how to read non iNes formatted roms")
	}
	header.PGR_banks = rom[4]
	header.CHR_banks = rom[5]
	//TODO - read bitewise values from controls, get memory mapper
	header.Control = rom[6]
	header.Control2 = rom[7]
	header.Ram_banks = rom[8] //if 0, set to 1, else value

	if header.Ram_banks == 0 {
		header.Ram_banks = 1
	}

	fmt.Printf("Nes Rom Info - %d PGR Banks, %d CHR Banks \n", header.PGR_banks, header.CHR_banks)
	fmt.Printf("Control = %d", header.Control)

	//Assuming MM0
	PGRBytes := (1024 * 16) * int(header.PGR_banks) // + 16 to ignore header? Removed
	fmt.Printf("Writing %d\n", PGRBytes)
	i := 0
	for i = 0; i < PGRBytes; i++ {
		nes.Cpu.WriteMemory(uint16(i+0x8000), rom[i+16])
	}
	fmt.Printf("wrote %d bytes \n", i)
	//Read bytes i to i + 8192

	ppuRam := rom[i+16 : i+8192+16]
	println("Read CHR Bank of")
	fmt.Printf("ppu BANK %x\n", ppuRam)

	copy(nes.Ppu.Memory[:], ppuRam)

	nes.Ppu.GetInfoForPatternTable()
	// for j := i; j < finish; j++ {

	// }

	// Now read 8k bytes and save to PPU ram.

	// fmt.Printf("%%", cpu.Memory)
	// if (4096 - 512) > 2^14 {
	// 	for i := 0; i < rom_length; i++ {
	// 		self.Memory[i+512] = rom[i]
	// 	}
	// }
}
