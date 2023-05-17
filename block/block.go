package block

import (
	"BreakoutGame/sound"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Block struct {
	StartX  int
	StartY  int
	EndX    int
	EndY    int
	Surface *sdl.Surface
	Window  *sdl.Window
}

func (b *Block) Collision(x, y int) bool {
	if b.StartX <= x && x <= b.EndX && y >= b.StartY && y <= b.EndY {
		return true
	}
	return false
}

func (b *Block) Erase(playsound bool) {
	rect := sdl.Rect{int32(b.StartX), int32(b.StartY), int32(b.EndX - b.StartX), int32(b.EndY - b.StartY)}
	colour := sdl.Color{R: 0, G: 0, B: 0, A: 255} //black color
	pixel := sdl.MapRGBA(b.Surface.Format, colour.R, colour.G, colour.B, colour.A)
	b.Surface.FillRect(&rect, pixel)

	b.Window.UpdateSurface()
	if playsound {
		go sound.PlaySound("sound/beep.mp3")
	}

	fmt.Printf("[i] block with initial cordinates %d %d erased\n", b.StartX, b.StartY)
}
