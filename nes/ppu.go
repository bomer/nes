package nes

import "image/color"

type Ppu struct {
	Cycle    int // 0-340
	ScanLine int // 0-261, 0-239=visible, 240=post, 241-260=vblank, 261=pre
	Frame    uint64
}

//GetColorFromPalette is used to grab an color.RGBA value from it's hex index of 64 colors.
//This is mostly used for some basic tests of the default color Palette.
func (p *Ppu) GetColorFromPalette(c int) color.RGBA {
	return (Palette[c])
}

func (p *Ppu) EmulateCycle() {
	//Do one pixel
}
