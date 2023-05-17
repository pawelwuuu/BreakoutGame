package main

import (
	"BreakoutGame/ball"
	"BreakoutGame/block"
	"BreakoutGame/platform"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	var WINDOW_WIDTH int32 = 800
	var WINDOW_HEIGHT int32 = 700

	var blocks []block.Block

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("BreakOut Game", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		WINDOW_WIDTH, WINDOW_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	for height := 5; height <= 135; height += 60 {
		for width := 11; width < 780; width += 157 {
			rect := sdl.Rect{int32(width), int32(height), 151, 50}
			colour := sdl.Color{R: 250, G: uint8(height), B: 0, A: 255} // purple
			pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
			surface.FillRect(&rect, pixel)

			blocks = append(blocks, block.Block{width, height, width + 151, height + 50, surface, window})
		}
	}

	println("cordinates of all blocks:")
	for _, v := range blocks {
		println(v.StartX, v.StartY, v.EndX, v.EndY)
	}
	window.UpdateSurface()

	platform := platform.Platform{int(WINDOW_WIDTH / 2), int(WINDOW_HEIGHT - 20), 100, 12, surface, window}
	platform.Place()

	ball := ball.Ball{int(WINDOW_WIDTH / 2), int(WINDOW_HEIGHT / 2), 15, 8, 19_000_000, true, &platform, blocks, surface, window}
	ball.Place()
	go ball.KickBall()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				keyCode := t.Keysym.Sym
				if keyCode == 1073741903 {
					println("[i] right arrow pressed")
					platform.Move(30)
				}

				if keyCode == 1073741904 {
					println("[i] left arrow pressed")
					platform.Move(-30)
				}
			}
		}
	}
}
