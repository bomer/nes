# Nes

My second emulator in Golang. The Nintendo Entertainment System

## About

This a project that's been done many times before but in the realm of aspiriational for me. Good excuse to get more comfortable with Golang and with thinking about programming closer to the metal. Hopefully this'll end up running with gomobile.

## Update 2023.

I have resumed this, after some on and off attempts over the years and am taking it more seriously.

Done so far this year:

- Loaded PPU Ram from the ROM which I previously had skipped
- Wrote code to pull out memory of 2x CHR rom banks and load into arrays of 8x8 sprites to render in a test GUI
- Added base gomobile GL Renderer to test real display out.

# TODO

- [x] Start doing PPU Tick Processing, building a buffer for display and processing background and sprites - Handy Reference - https://austinmorlan.com/posts/nes_rendering_overview/
- [ ] Load Background onto buffer and render
- [ ] Memory Mirroring for PPU
- [x] Load Sprites
- [ ] Memory mapping to push updates to the PPU Memory, specifically OAM + attribute tables else screen will never update
- [ ] Controls mapping
- [ ] Making mario fully playable

# Far Future:

- Proper Memory Mapping (Basically hard coded MM0/1)
- Sound
