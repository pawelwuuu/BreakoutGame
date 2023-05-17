package platform

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Platform struct {
	X       int
	Y       int
	Width   int
	Height  int
	Surface *sdl.Surface
	Window  *sdl.Window
}

func (p *Platform) Place() {
	rect := sdl.Rect{int32(p.X), int32(p.Y), int32(p.Width), int32(p.Height)}
	colour := sdl.Color{R: 255, G: 255, B: 255, A: 255} //white color
	pixel := sdl.MapRGBA(p.Surface.Format, colour.R, colour.G, colour.B, colour.A)
	p.Surface.FillRect(&rect, pixel)

	p.Window.UpdateSurface()

	fmt.Printf("[i] placing platform to x: %d y: %d\n", p.X, p.Y)
}

func (p *Platform) Move(moveX int) {
	winWidht, _ := p.Window.GetSize()
	if p.X+moveX >= -90 && p.X+moveX < int(winWidht) {
		rect := sdl.Rect{int32(p.X), int32(p.Y), int32(p.Width), int32(p.Height)}
		colour := sdl.Color{R: 0, G: 0, B: 0, A: 255} //white color
		pixel := sdl.MapRGBA(p.Surface.Format, colour.R, colour.G, colour.B, colour.A)
		p.Surface.FillRect(&rect, pixel)

		p.X += moveX
		rect = sdl.Rect{int32(p.X), int32(p.Y), int32(p.Width), int32(p.Height)}
		colour = sdl.Color{R: 255, G: 255, B: 255, A: 255} //white color
		pixel = sdl.MapRGBA(p.Surface.Format, colour.R, colour.G, colour.B, colour.A)
		p.Surface.FillRect(&rect, pixel)

		fmt.Printf("[i] placing platform to x: %d y: %d\n", p.X, p.Y)

		p.Window.UpdateSurface()
	}
}

func (p *Platform) Collision(x, y int) bool {
	if p.X <= x && x <= p.X+p.Width && y >= p.Y && y <= p.Y+p.Height {
		return true
	}
	return false
}
