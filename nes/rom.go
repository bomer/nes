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
	control   byte
}

// Where the cartridge is read in
//Read file in curent dir into Memory
func LoadGame(filename string) {
	var header RomHeader
	rom, _ := ioutil.ReadFile(filename)
	rom_length := len(rom)
	if rom_length > 0 {
		fmt.Printf("Rom Length = %d\n", rom_length)
	}

	if rom[3] != 0x1a {
		panic("OH GOD WHAT IS THIS. I don't know how to read non iNes formatted roms")
	}
	header.PGR_banks = rom[4]
	header.CHR_banks = rom[5]
	header.control = rom[6]

	fmt.Printf("Nes Rom Info - %d PGR Banks, %d CHR Banks \n", header.PGR_banks, header.CHR_banks)

	// //If room to store ROM in RAM, start at 512 or 0x200
	// if (4096 - 512) > rom_length {
	// 	for i := 0; i < rom_length; i++ {
	// 		self.Memory[i+512] = rom[i]
	// 	}
	// }
}
