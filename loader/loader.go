package main

import (
	"log"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/yukiisbored/Kaori/game"
)

const FPS = 60
const DELAY_TIME = 1000 / FPS

var frameStart, frameTime uint32

func main() {
	log.Println("Welcome to Kaori")

	log.Printf("Compiled with %s for %s %s\n", runtime.Compiler, runtime.GOOS, runtime.GOARCH)

	game.Init("Kaori", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 800, 600, false)

	for game.Running() {
		frameStart = sdl.GetTicks()

		game.HandleEvents()
		game.Update()
		game.Render()

		frameTime = sdl.GetTicks() - frameStart

		if frameTime < DELAY_TIME {
			sdl.Delay(DELAY_TIME - frameTime)
		}
	}

	game.Clean()

	log.Println("Goodbye o/")
}
