package nes

import (
	"fmt"
	"image/color"
)

// Ppu Memory Map.
// PPU memory map
// The PPU addresses a 14-bit (16kB) address space, $0000-3FFF, completely separate from the CPU's address bus.
// It is either directly accessed by the PPU itself, or via the CPU with memory mapped registers at $2006 and $2007.
// The NES has 2kB of RAM dedicated to the PPU, normally mapped to the nametable address space from $2000-2FFF, but this can be rerouted through custom cartridge wiring.

// Address range	Size	Description
// $0000-$0FFF	$1000	Pattern table 0
// $1000-$1FFF	$1000	Pattern table 1
// $2000-$23FF	$0400	Nametable 0
// $2400-$27FF	$0400	Nametable 1
// $2800-$2BFF	$0400	Nametable 2
// $2C00-$2FFF	$0400	Nametable 3
// $3000-$3EFF	$0F00	Mirrors of $2000-$2EFF
// $3F00-$3F1F	$0020	Palette RAM indexes
// $3F20-$3FFF	$00E0	Mirrors of $3F00-$3F1F

// In addition, the PPU internally contains 256 bytes of memory known as Object Attribute Memory which determines how sprites are rendered. The CPU can manipulate this memory through memory mapped registers at OAMADDR ($2003), OAMDATA ($2004), and OAMDMA ($4014). OAM can be viewed as an array with 64 entries. Each entry has 4 bytes: the sprite Y coordinate, the sprite tile number, the sprite attribute, and the sprite X coordinate.

// Address Low Nibble	Description
// $00, $04, $08, $0C	Sprite Y coordinate
// $01, $05, $09, $0D	Sprite tile #
// $02, $06, $0A, $0E	Sprite attribute
// $03, $07, $0B, $0F	Sprite X coordinate
// Hardware mapping
// The mappings above are the fixed addresses from which the PPU uses to fetch data during rendering. The actual device that the PPU fetches data from, however, may be configured by the cartridge.

// $0000-1FFF is normally mapped by the cartridge to a CHR-ROM or CHR-RAM, often with a bank switching mechanism.
// $2000-2FFF is normally mapped to the 2kB NES internal VRAM, providing 2 nametables with a mirroring configuration controlled by the cartridge, but it can be partly or fully remapped to RAM on the cartridge, allowing up to 4 simultaneous nametables.
// $3000-3EFF is usually a mirror of the 2kB region from $2000-2EFF. The PPU does not render from this address range, so this space has negligible utility.
// $3F00-3FFF is not configurable, always mapped to the internal palette control.

type Ppu struct {
	Cycle    int // 0-340
	ScanLine int // 0-261, 0-239=are visible frames, 240=post, 241-260=vblank, 261=pre
	Frame    uint64
	Memory   [0x3FFF + 1]byte // 16kb address space.
}

//Set Memory
// func (p *Ppu) SetMemory(chrbanks []byte) {
// 	// p.Memory = chrbanks
// 	copy(p.Memory[:], chrbanks)
// }

func (p *Ppu) GetInfoForPatternTable() {
	println("Checking at address 0x00, x:0,y:0")
	println(p.Memory[0x00:0x02])

	fmt.Printf("ppu BANK 0x%b\n", p.Memory[0:16])
	tile := p.Memory[0:16]
	for i, v := range tile {
		fmt.Printf("Tile: i: %d %b\n", i, v)
		//  128 64 32 16  8 4 2 1
		fmt.Printf("Bit 1 %d \n", v&1 == 1)
		fmt.Printf("Bit 2 %d \n", v&2)
		fmt.Printf("Bit 3 %d \n", v&4)
		fmt.Printf("Bit 4 %d \n", v&8)
		fmt.Printf("Bit 5 %d \n", v&16)
		fmt.Printf("Bit 6 %d \n", v&32)
		fmt.Printf("Bit 7 %d \n", v&128)
		fmt.Printf("Bit 8 %d \n", v&255)
		// fmt.Printf("Bit 1 %d \n", v&1)
	}
	//Now build a tile with values 0,1,2,3

	// fmt.Printf("ppu BANK 0x%d\n", p.Memory[0x01])
}

//GetColorFromPalette is used to grab an color.RGBA value from it's hex index of 64 colors.
//This is mostly used for some basic tests of the default color Palette.
func (p *Ppu) GetColorFromPalette(c int) color.RGBA {
	return (Palette[c])
}

// EmulateCycle is called 3 times for ever 1 CPU Cycle.
// There are 262 Scanlines per frame.
// Each scanline is 341 PPU Cycles
func (p *Ppu) EmulateCycle() {
	//Do one pixel
}
