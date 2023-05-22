What is an emulator?

A piece of software that simulates a piece of hardware functions, as accurately as possible so that the software can run?

Is it used for things other than pirating old games?

100%. In the world of music plugins and synthesiser, lots of old hardware has be digitially re-created to keep those sounds alive. It can be used to keep old hardware alive, or develop on a new platform that you might not be running, like android / ios simulators, or Qemu which can be used for running for mac os run times or emulating old hardware like a palm pilot.

Why do I James want to explore emulation?

For those tech savvy video game players who had early internet access, there was a magical time in the early days of the internet, when one first discovered you could play nintendo games, pokemon and ocarina of time on a computer (if your computer was fast enough). This was something of a radical idea, as it was turning current or older hardware that usually had a high cost to purchase and own into a free piece of software.

It also opened the world to more games, such as rare titles not released in a region, some of which were translated into english for the first time by fan communities (Like Final Fantasy V, Pokemon Green).

Emulation became a staple of what I would do with my computing devices. I got Snes/GB game running on a palm pilot, an old Nokia phone, a PSP, the first iMac, a PS3/PS4, Android Phones, iPad as well as custom Android devices. If I have a devices with an open-ness I want it to be able to run old games.

When I learnt to code, I eventually looked at some source code and it looked like this.

I put this into the realm of god-like programmers from era's by gone that can do the times table in binary, order a drink in assemble and buy something off ebay using telnet.

I eventually discovered a youtuber, bisqwit, who speaks like a speak and spell, is from finland and broken down how to create a first "easy emulator" in the form of the chip8. I watched the video, read a bunch up on it, and after about 2 months of coding finally got games working on it. It was challenging, but it was 100 tiny challenges that when strung together produced a great output.

I might just address, why Golang? Honestly, just because I wanted to. Golang is fast to build and run, and is statically typed. I could classify it as powerful and low level as C, with the Garbage Collection of Java, maybe the best standard library of any programming language I've used, and is great at producing simple fast builds. Because most of what I would be interacting with is bytes, occasionally characters, a language like JS wouldn't really be suitable, or rather, that fun because a "Number" isn't really a suitable tool for doing binary operations. Secondly whilst JS is plenty portable, I had already establishe with Gomobile which has OpenGL bindings I can run on windows, mac and android very easily with the only painpoint being its OpenGLES2 (shaders for everything). Thirdly, I love the testing tools, and because I won't have a bootable game in a while, simple and effective testing is important.

PIC

The NES became a bigger dream that would take a lot more time and energy as it had much more advanced capabilities, memory mapped catridges that could swap graphics slots, complexitity in the drawing of sprites, backgrounds, sound, etc. So without further adieu, let's get into it.

The Nintendo Entertainment System was the first western released console in the year 1985. It launched with Super Mario Bros and went onto to sell 62 million units.

Now remember, an emulator is software than pretends to be another pieces of hardware, so we have to understand this hardware incredibly well.

The NES has a series of core components, CPU, Memory, PPU (picture processor unit), APU (Audio Processor Unit), a Rom Loader, 2 controller inputs.

We'll focus on the CPU first as this is the heart of it's processing capability. The NES uses a 6502, same as the Apple II. It has a well defined and well documented set of instructions.

At startup, the ROM is loaded (from a file into memory, containing the sprite banks which goto the PPU and the instructions sets which goto the CPU.

Instructions are loaded into memory as binary from the ROM (cartridge), and these translate into operations the CPU can understand with a replacement of hex/binary to a specific instruction.

The NES CPU is controlled by the basic components. 64kb of RAM, A program counter (What instruction am I up to), a Stack Pointer (where am I up to in stack memory, a small subset of memory that acts like a stack), 4 registers for temorary Storage of values (accumulator, X,Y values and a single bytes whose bits store status flags.

In emulator code, we end up with this.

Now how does this actually execute?

$1000 $8D STA
$1001
$1002
$1003
$1004
$1005
