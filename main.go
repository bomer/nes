package main

import (
	"log/slog"

	"github.com/bomer/nes/nes"
)

var myNes nes.Nes

func main() {

	slog.SetLogLoggerLevel(slog.LevelError)

	nes.Debugf("Initing...")
	myNes.Cpu.Quiet = true
	myNes.Rom.LoadGame("mario.nes", &myNes)
	// myNes.Cpu.Quiet = false
	myNes.Cpu.Debug = false
	myNes.Cpu.DebugLines = 128
	myNes.Cpu.System = &myNes
	myNes.Ppu.System = &myNes
	myNes.Init()
	// myNes.Cpu

	// for {
	// 	myNes.Cpu.EmulateCycle()
	// }

}
