package game

import (
	"io/ioutil"
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/yukiisbored/Kaori/texture"
	"github.com/yukiisbored/Kaori/tilemap"
)

// DemoScene is an example use of a Scene in this case drawing the logo
type DemoScene struct {
	testMap *tilemap.Map
}

// Enter executes when the scene is starting
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

	log.Println("Demo // Loading Tiled Map ...")

	// Load the map file
	data, err := ioutil.ReadFile("./assets/map.tmx")

	if err != nil {
		log.Fatalln("Demo // Oh no, can't load map :(")
		log.Panic(err)
	}

	// Parse tmx data
	log.Println("Demo // Parsing TMX Data...")

	s.testMap = new(tilemap.Map)
	err = tilemap.Unmarshal(data, s.testMap)

	if err != nil {
		log.Fatalln("Demo // Oh no, can't parse map :(")
		log.Panic(err)
	}

	// Load Tile Map's Tilesets
	log.Println("Demo // Loading Tiled Map's Tilesets ...")

	for _, ts := range s.testMap.Tilesets {
		ts.Load(renderer, "./assets")
	}
}

// Update executes when a game update is being executed
func (s *DemoScene) Update() {
	// Since this scene doesn't have user interaction at all we leave this empty
}

// Draw executes when a game render is being executed
func (s *DemoScene) Draw(r *sdl.Renderer) {
	// Get the current size of the window
	w, h := window.GetSize()

	// Move the logo relatively to how long it's been running
	// but don't move outisde the window
	xLogo := tick * 4 % w
	yLogo := tick * 4 % h

	// Rotate the logo relatively to how long it's been running
	rot := tick * 4 % 360

	// Draw Map

	s.testMap.Draw(r, 0, -800)

	// Draw the logos!
	texture.Draw(renderer, "kaori",
		int32(xLogo), int32(yLogo),
		474, 167, float64(rot), sdl.FLIP_NONE)

	texture.Draw(renderer, "kaori",
		int32(w-xLogo), int32(h-yLogo),
		474, 167, float64(360-rot), sdl.FLIP_NONE)
}

// HandleEvents executes when there's an SDL Event from the Event Poll
func (s *DemoScene) HandleEvents(e sdl.Event) {
}

// Exit executes when the scene is being changed or the game is closing
func (s *DemoScene) Exit() {
	log.Println("Demo // Freeing Texture")

	// Free the logo texture
	texture.Free("kaori")

	log.Println("Demo // Bye :(")
}
