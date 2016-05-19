package game

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/yukiisbored/Kaori/scene"
	"github.com/yukiisbored/Kaori/texture"
)

type DemoScene struct {
	scene.Scene
}

func (s DemoScene) Enter() {
	log.Println("Demo // Welcome to Kaori's Demo Scene!")
	log.Println("Demo // Loading Logo as Texture ...")

	err := texture.Load(renderer, "./assets/kaori.png", "kaori")

	if err != nil {
		log.Fatalln("Demo // Oh no, can't load logo :(")
		log.Panic(err)
	}
}

func (s DemoScene) Update() {
}

func (s DemoScene) Draw(r *sdl.Renderer) {
	w, h := window.GetSize()

	xLogo := tick * 4 % w
	yLogo := tick * 4 % h
	rot := tick * 4 % 360

	texture.Draw(renderer, "kaori",
		int32(xLogo), int32(yLogo),
		474, 167, float64(rot), sdl.FLIP_NONE)

	texture.Draw(renderer, "kaori",
		int32(w-xLogo), int32(h-yLogo),
		474, 167, float64(360-rot), sdl.FLIP_NONE)
}

func (s DemoScene) HandleEvents(e sdl.Event) {
}

func (s DemoScene) Exit() {
	log.Println("Demo // Freeing Texture")
	texture.Free("kaori")
	log.Println("Demo // Bye :(")
}
