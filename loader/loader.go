// Loader is an example of a usual Kaori Game Loader for the PC
package main

import (
	"log"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/yukiisbored/Kaori/game"
)

const FPS = 60                // Limit the FPS to 60
const DELAY_TIME = 1000 / FPS // How long should we delay to get 60 frames in every second

var frameStart, frameTime uint32 // When did the frame start and How long does it take to do a game update and render

func main() {
	// Show us a welomce message
	log.Println("Welcome to Kaori")

	// Don't forget to say goodbye ;)
	defer log.Println("Goodbye o/")

	// Show information about the runtime
	log.Printf("Compiled with %s for %s %s\n", runtime.Compiler, runtime.GOOS, runtime.GOARCH)

	// Don't forget to clean the game after it's done
	game.Init("Kaori", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 800, 600, false)

	// Don't forget to clean the game after it's done
	defer game.Clean()

	for game.Running() {
		// Get the current time to mark a frame start
		frameStart = sdl.GetTicks()

		// Run all of the 'update' functions
		game.HandleEvents()
		game.Update()
		game.Render()

		// Record the time
		frameTime = sdl.GetTicks() - frameStart

		// Check if it's faster than delay time
		if frameTime < DELAY_TIME {
			// If it is faster, delay the game to prevent "speeding" on certain tick event
			sdl.Delay(DELAY_TIME - frameTime)
		}
	}
}
