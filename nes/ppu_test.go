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
