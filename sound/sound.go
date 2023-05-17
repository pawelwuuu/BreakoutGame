package sound

import (
	"log"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func PlaySound(path string) {
	if err := mix.Init(mix.INIT_MP3); err != nil {
		log.Println(err)
		return
	}
	defer mix.Quit()

	if err := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Println(err)
		return
	}
	defer mix.CloseAudio()

	if music, err := mix.LoadMUS(path); err != nil {
		log.Println(err)
	} else if err = music.Play(1); err != nil {
		log.Println(err)
	} else {
		sdl.Delay(1000)
		music.Free()
	}
}
