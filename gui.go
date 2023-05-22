package main

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/bomer/nes/nes"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/exp/app/debug"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/gl"
)

var guiNes nes.Nes

var (
	images   *glutil.Images
	fps      *debug.FPS
	program  gl.Program
	position gl.Attrib
	offset   gl.Uniform
	color    gl.Uniform
	buf      gl.Buffer
	green    float32
	touchX   float32
	touchY   float32

	img glutil.Image
)

func main() {

	fmt.Printf("Initing...")
	guiNes.Cpu.Quiet = true
	guiNes.Rom.LoadGame("mario.nes", &guiNes)
	// guiNes.Cpu.Quiet = false
	go guiNes.Init()
	// guiNes.Cpu

	// for {
	// 	guiNes.Cpu.EmulateCycle()
	// }

	app.Main(func(a app.App) {
		var glctx gl.Context
		var sz size.Event
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					glctx, _ = e.DrawContext.(gl.Context)
					onStart(glctx)
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					onStop(glctx)
					glctx = nil
				}
			case size.Event:
				sz = e
				touchX = float32(sz.WidthPx / 2)
				touchY = float32(sz.HeightPx / 2)
			case paint.Event:
				if glctx == nil || e.External {
					// As we are actively painting as fast as
					// we can (usually 60 FPS), skip any paint
					// events sent by the system.
					continue
				}
				onPaint(glctx, sz)
				a.Publish()
				// Drive the animation by preparing to paint the next frame
				// after this one is shown.
				a.Send(paint.Event{})
			case touch.Event:
				touchX = e.X
				touchY = e.Y
			}
		}
	})
}
func onStart(glctx gl.Context) {
	var err error
	program, err = glutil.CreateProgram(glctx, vertexShader, fragmentShader)
	if err != nil {
		log.Printf("error creating GL program: %v", err)
		return
	}
	buf = glctx.CreateBuffer()
	// glctx.BindBuffer(gl.ARRAY_BUFFER, buf)
	// glctx.BufferData(gl.ARRAY_BUFFER, triangleData, gl.STATIC_DRAW)
	position = glctx.GetAttribLocation(program, "position")
	color = glctx.GetUniformLocation(program, "color")
	offset = glctx.GetUniformLocation(program, "offset")
	images = glutil.NewImages(glctx)
	fps = debug.NewFPS(images)

	//Buffer for display Buffer
	img = *images.NewImage(256, 240)
}
func onStop(glctx gl.Context) {
	glctx.DeleteProgram(program)
	glctx.DeleteBuffer(buf)
	fps.Release()
	images.Release()
}

func onPaint(glctx gl.Context, sz size.Event) {

	glctx.ClearColor(1, 1, 1, 1)
	glctx.Clear(gl.COLOR_BUFFER_BIT)

	glctx.UseProgram(program)

	glctx.BindBuffer(gl.ARRAY_BUFFER, buf)

	//Draw Pixels onto screen
	// for i := 0; i < 64; i++ {
	// for j := 0; j < 32; j++ {
	// if myChip8.Gfx[(j*64)+i] == 0 {
	// img.RGBA.Set(i, j, image.Black)
	// }

	// }
	// }

	// green := image.NewUniform(color.RGBA{0x00, 0x1f, 0x00, 0xff})
	//For each 256 Sprites > Sprite = [8][8]uint8
	count := 0
	// countx := 0
	county := 0
	// xoffset := 0
	xoffset := 0 //((count % 2) == 0) * 8
	for _, sprite := range guiNes.Ppu.TileMap {
		//For each row of pixels
		for rowindex, arrayOfRows := range sprite {
			//For each pixels in each row...
			for pixelindex, pixelvalue := range arrayOfRows {
				// fmt.Printf("Reading value of of $@ ", pixelvalue)
				if pixelvalue != 0 {
					img.RGBA.Set(pixelindex+xoffset, rowindex+county, guiNes.Ppu.GetColorFromPalette(int(pixelvalue+1)))
				}
			}
		}
		count += 1
		// if (count % 2) == 0 {
		xoffset += 8
		// }
		if (count%16) == 0 && count > 0 {
			county += 8
			xoffset = 0
		}

	}

	//Draw over whole screen
	//Changed to widthPT which gives the real edge of the screen instead of pixels.
	tl := geom.Point{0, 0}
	tr := geom.Point{geom.Pt(sz.WidthPt), 0}
	bl := geom.Point{0, geom.Pt(sz.HeightPt)}
	img.Upload()

	// Set up the texture
	glctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	glctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	glctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	glctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	img.Draw(sz, tl, tr, bl, img.RGBA.Bounds())
	// fps.Draw(sz)

	//cleanup every  frame
	img.Release()
	img = *images.NewImage(256, 224)

}

const squareoffset = 0.057

var triangleData = f32.Bytes(binary.LittleEndian,
	0.0, squareoffset, 0.0, // top left
	0.0, 0.0, 0.0, // bottom left
	squareoffset, 0.0, 0.0, // bottom right
	squareoffset, squareoffset, 0.0,
)

const (
	coordsPerVertex = 3
	vertexCount     = 4
)

const vertexShader = `#version 100
uniform vec2 offset;

attribute vec4 position;
void main() {
	// offset comes in with x/y values between 0 and 1.
	// position bounds are -1 to 1.
	vec4 offset4 = vec4(2.0*offset.x-1.0, 1.0-2.0*offset.y, 0, 0);
	gl_Position = position + offset4;
}`

const fragmentShader = `#version 100
precision mediump float;
uniform vec4 color;
void main() {
	gl_FragColor = color;
}`
