package ball

import (
	"BreakoutGame/block"
	"BreakoutGame/platform"
	"BreakoutGame/sound"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Ball struct {
	PosX       int
	PosY       int
	Radius     int
	Dirrection int
	Speed      int
	Running    bool
	Platform   *platform.Platform
	Blocks     []block.Block
	Surface    *sdl.Surface
	Window     *sdl.Window
}

func (b *Ball) Place() {
	rect := sdl.Rect{int32(b.PosX), int32(b.PosY), int32(b.Radius), int32(b.Radius)}
	colour := sdl.Color{R: 192, G: 192, B: 192, A: 255} //silver color
	pixel := sdl.MapRGBA(b.Surface.Format, colour.R, colour.G, colour.B, colour.A)
	b.Surface.FillRect(&rect, pixel)
	b.Window.UpdateSurface()

	// fmt.Printf("[i] ball placed ad %d %d wtih radius of %d\n", b.PosX, b.PosY, b.Radius)
}

func (b *Ball) KickBall() {
	for b.Running {
		time.Sleep(time.Duration(b.Speed))
		//removing old ball
		rect := sdl.Rect{int32(b.PosX), int32(b.PosY), int32(b.Radius), int32(b.Radius)}
		colour := sdl.Color{R: 0, G: 0, B: 0, A: 255} //silver color
		pixel := sdl.MapRGBA(b.Surface.Format, colour.R, colour.G, colour.B, colour.A)
		b.Surface.FillRect(&rect, pixel)

		b.UpdateBallDir()
		b.PushBall()
		b.Place()

		if len(b.Blocks) == 0 {
			go sound.PlaySound("sound/win.wav")
			b.Running = false
		}
	}

}

func (b *Ball) PushBall() {
	pushVal := 5
	switch b.Dirrection {
	case 1:
		b.PosX += pushVal
		b.PosY += pushVal
	case 2:
		b.PosX += pushVal
	case 3:
		b.PosX += pushVal
		b.PosY -= pushVal
	case 4:
		b.PosY -= pushVal
	case 5:
		b.PosX -= pushVal
		b.PosY -= pushVal
	case 6:
		b.PosX -= pushVal
	case 7:
		b.PosX -= pushVal
		b.PosY += pushVal
	case 8:
		b.PosY += pushVal
	}
}

func (b *Ball) UpdateBallDir() {
	x := b.PosX
	y := b.PosY
	maxX, maxY := b.Window.GetSize()

	//right side of game
	if x+b.Radius == int(maxX) && b.Dirrection == 3 {
		b.Dirrection = 5
	}

	if x+b.Radius == int(maxX) && b.Dirrection == 1 {
		b.Dirrection = 7
	}

	//left side of game
	if x == 0 && b.Dirrection == 5 {
		b.Dirrection = 3
	}

	if x == 0 && b.Dirrection == 7 {
		b.Dirrection = 1
	}

	//top side of game
	if y == 0 && b.Dirrection == 5 {
		b.Dirrection = 7
	}

	if y == 0 && b.Dirrection == 3 {
		b.Dirrection = 1
	}

	if y == 0 && b.Dirrection == 4 {
		b.Dirrection = 8
	}

	//bottom side of game
	if y > int(maxY) {
		go sound.PlaySound("sound/lose.wav")
		b.eraseAll()
		b.Running = false
	}

	//blocks
	for i, v := range b.Blocks {
		if v.Collision(x, y) {
			v.Erase(true)
			b.Blocks = append(b.Blocks[:i], b.Blocks[i+1:]...)
			b.Dirrection = 1
			b.Speed = b.Speed - 1_000_000
		} else if v.Collision(x+b.Radius, y) {
			v.Erase(true)
			b.Blocks = append(b.Blocks[:i], b.Blocks[i+1:]...)
			b.Dirrection = 7
			b.Speed = b.Speed - 1_000_000
		} else if v.Collision(x+b.Radius, y+b.Radius) {
			v.Erase(true)
			b.Blocks = append(b.Blocks[:i], b.Blocks[i+1:]...)
			b.Dirrection = 1
			b.Speed = b.Speed - 1_000_000
		} else if v.Collision(x, y+b.Radius) {
			v.Erase(true)
			b.Blocks = append(b.Blocks[:i], b.Blocks[i+1:]...)
			b.Dirrection = 7
			b.Speed = b.Speed - 1_000_000
		}
	}

	//platform
	if b.Platform.Collision(x, y+b.Radius) || b.Platform.Collision(x+b.Radius, y+b.Radius) {
		relBallPos := x - b.Platform.X

		if relBallPos < b.Platform.Width/3 {
			b.Dirrection = 5
		} else if relBallPos < b.Platform.Width*2/3 {
			b.Dirrection = 4
		} else {
			b.Dirrection = 3
		}
	}
}

func (b *Ball) eraseAll() {
	for _, v := range b.Blocks {
		time.Sleep(1000000000)
		v.Erase(false)
	}
}
