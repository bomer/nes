package nes

type Nes struct {
	Cpu Cpu
}

func (nes *Nes) Init() {
	nes.Cpu.Init()
	LoadGame("mario.nes")

}
