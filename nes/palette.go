package nes

import "image/color"

//Palette contains 64 colors, referenced in Hex for rendering.
//This is the default palette, I think you can have others?
//No tests for this file, as go coverage requires functions for test coverage.
var Palette = map[int]color.RGBA{
	0x00: {0x75, 0x75, 0x75, 0xff},
	0x01: {0x27, 0x1B, 0x8F, 0xff},
	0x02: {0x00, 0x00, 0xAB, 0xff},
	0x03: {0x47, 0x00, 0x9F, 0xff},
	0x04: {0x8F, 0x00, 0x77, 0xff},
	0x05: {0xAB, 0x00, 0x13, 0xff},
	0x06: {0xA7, 0x00, 0x00, 0xff},
	0x07: {0x7F, 0x0B, 0x00, 0xff},
	0x08: {0x43, 0x2F, 0x00, 0xff},
	0x09: {0x00, 0x47, 0x00, 0xff},
	0x0A: {0x00, 0x51, 0x00, 0xff},
	0x0B: {0x00, 0x3F, 0x17, 0xff},
	0x0C: {0x1B, 0x3F, 0x5F, 0xff},
	0x0D: {0x00, 0x00, 0x00, 0xff},
	0x0E: {0x00, 0x00, 0x00, 0xff},
	0x0F: {0x00, 0x00, 0x00, 0xff},
	0x10: {0xBC, 0xBC, 0xBC, 0xff},
	0x11: {0x00, 0x73, 0xEF, 0xff},
	0x12: {0x23, 0x3B, 0xEF, 0xff},
	0x13: {0x83, 0x00, 0xF3, 0xff},
	0x14: {0xBF, 0x00, 0xBF, 0xff},
	0x15: {0xE7, 0x00, 0x5B, 0xff},
	0x16: {0xDB, 0x2B, 0x00, 0xff},
	0x17: {0xCB, 0x4F, 0x0F, 0xff},
	0x18: {0x8B, 0x73, 0x00, 0xff},
	0x19: {0x00, 0x97, 0x00, 0xff},
	0x1A: {0x00, 0xAB, 0x00, 0xff},
	0x1B: {0x00, 0x93, 0x3B, 0xff},
	0x1C: {0x00, 0x83, 0x8B, 0xff},
	0x1D: {0x00, 0x00, 0x00, 0xff},
	0x1E: {0x00, 0x00, 0x00, 0xff},
	0x1F: {0x00, 0x00, 0x00, 0xff},
	0x20: {0xFF, 0xFF, 0xFF, 0xff},
	0x21: {0x3F, 0xBF, 0xFF, 0xff},
	0x22: {0x5F, 0x97, 0xFF, 0xff},
	0x23: {0xA7, 0x8B, 0xFD, 0xff},
	0x24: {0xF7, 0x7B, 0xFF, 0xff},
	0x25: {0xFF, 0x77, 0xB7, 0xff},
	0x26: {0xFF, 0x77, 0x63, 0xff},
	0x27: {0xFF, 0x9B, 0x3B, 0xff},
	0x28: {0xF3, 0xBF, 0x3F, 0xff},
	0x29: {0x83, 0xD3, 0x13, 0xff},
	0x2A: {0x4F, 0xDF, 0x4B, 0xff},
	0x2B: {0x58, 0xF8, 0x98, 0xff},
	0x2C: {0x00, 0xEB, 0xDB, 0xff},
	0x2D: {0x00, 0x00, 0x00, 0xff},
	0x2E: {0x00, 0x00, 0x00, 0xff},
	0x2F: {0x00, 0x00, 0x00, 0xff},
	0x30: {0xFF, 0xFF, 0xFF, 0xff},
	0x31: {0xAB, 0xE7, 0xFF, 0xff},
	0x32: {0xC7, 0xD7, 0xFF, 0xff},
	0x33: {0xD7, 0xCB, 0xFF, 0xff},
	0x34: {0xFF, 0xC7, 0xFF, 0xff},
	0x35: {0xFF, 0xC7, 0xDB, 0xff},
	0x36: {0xFF, 0xBF, 0xB3, 0xff},
	0x37: {0xFF, 0xDB, 0xAB, 0xff},
	0x38: {0xFF, 0xE7, 0xA3, 0xff},
	0x39: {0xE3, 0xFF, 0xA3, 0xff},
	0x3A: {0xAB, 0xF3, 0xBF, 0xff},
	0x3B: {0xB3, 0xFF, 0xCF, 0xff},
	0x3C: {0x9F, 0xFF, 0xF3, 0xff},
	0x3D: {0x00, 0x00, 0x00, 0xff},
	0x3E: {0x00, 0x00, 0x00, 0xff},
	0x3F: {0x00, 0x00, 0x00, 0xff},
}
