package game

import (
	"io/ioutil"
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/yukiisbored/Kaori/texture"
	"github.com/yukiisbored/Kaori/tilemap"
)

// Demo Scene is an example use of a Scene in this case drawing the logo
type DemoScene struct {
	testMap *tilemap.Map
}

func (s *DemoScene) Enter() {
	// Show a warm welomce message
	log.Println("Demo // Welcome to Kaori's Demo Scene!")
	log.Println("Demo // Loading Logo as Texture ...")

	// Load the logo as a texture using Kaori's texture module
	err := texture.Load(renderer, "./assets/kaori.png", "kaori")

	if err != nil {
		log.Fatalln("Demo // Oh no, can't load logo :(")
		log.Panic(err)
	}

	// Load the map file
	data, err := ioutil.ReadFile("./assets/map.tmx")

	if err != nil {
		log.Fatalln("Demo // Oh no, can't load map :(")
		log.Panic(err)
	}

	err = tilemap.Unmarshal(data, s.testMap)

	if err != nil {
		log.Fatalln("Demo // Oh no, can't parse map :(")
		log.Panic(err)
	}

	for _, ts := range s.testMap.Tilesets {
		ts.Load(renderer, "./assets")
	}
}

func (s *DemoScene) Update() {
	// Since this scene doesn't have user interaction at all we leave this empty
}

func (s *DemoScene) Draw(r *sdl.Renderer) {
	// Get the current size of the window
	w, h := window.GetSize()

	// Move the logo relatively to how long it's been running
	// but don't move outisde the window
	xLogo := tick * 4 % w
	yLogo := tick * 4 % h

	// Rotate the logo relatively to how long it's been running
	rot := tick * 4 % 360

	// Render Map
	//s.testMap.Draw(r, 0, 0)

	// Draw the logos!
	texture.Draw(renderer, "kaori",
		int32(xLogo), int32(yLogo),
		474, 167, float64(rot), sdl.FLIP_NONE)

	texture.Draw(renderer, "kaori",
		int32(w-xLogo), int32(h-yLogo),
		474, 167, float64(360-rot), sdl.FLIP_NONE)
}

func (s *DemoScene) HandleEvents(e sdl.Event) {
}

func (s *DemoScene) Exit() {
	log.Println("Demo // Freeing Texture")

	// Free the logo texture
	texture.Free("kaori")

	log.Println("Demo // Bye :(")
}
