package nes

import "image/color"

type Ppu struct {
}

//Function to grab an color.RGBA value from it's hex index of 64 colors.
//This is mostly used for some basic tests of the default color Palette.
func (p *Ppu) GetColorFromPalette(c int) color.RGBA {
	return (Palette[c])
}

func tick() {

}
