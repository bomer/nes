package nes_test

import (
	"testing"

	Nes "github.com/bomer/nes/nes"
)

var ppuNes Nes.Nes

func Test0x00Color(t *testing.T) {
	color := ppuNes.Ppu.GetColorFromPalette(0x00)

	if color.R != 117 && color.G != 117 && color.B != 117 {
		t.Error("Hex value issue with 0x00")
	}
}

func Test0x15Color(t *testing.T) {
	color := ppuNes.Ppu.GetColorFromPalette(0x15)

	if color.R != 231 && color.G != 0 && color.B != 91 {
		t.Error("Hex value issue with 0x15")
	}
}

func TestVblankFlags(t *testing.T) {
	ppuNes.Ppu.PPUSTATUS = 0
	ppuNes.Ppu.ScanLine = 241
	ppuNes.Ppu.Cycle = 1
	if ppuNes.Ppu.PPUSTATUS != 0 {
		t.Errorf("Error, Bad init state on vblank, got 0x%x", ppuNes.Ppu.PPUSTATUS)
	}
	ppuNes.Ppu.EmulateCycle()
	if ppuNes.Ppu.PPUSTATUS != 0x80 {
		t.Errorf("Error, Did not update PPU Status, got 0x%x", ppuNes.Ppu.PPUSTATUS)
	}
}
