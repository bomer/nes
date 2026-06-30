package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/bomer/nes/nes"
)

var myNes nes.Nes

// Your custom helper
func debugf(format string, args ...any) {
	slog.Logf(context.Background(), slog.LevelDebug, format, args...)
}

func main() {

	fmt.Printf("Initing...")
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
