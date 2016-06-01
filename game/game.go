// Package game provides an example of a usual Kaori game skeleton.
// Everything that are specific to the game should and only be in
// this class only.
package game

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/yukiisbored/Kaori/input"
	"github.com/yukiisbored/Kaori/scene"
	"github.com/yukiisbored/Kaori/texture"
)

var (
	window   *sdl.Window
	renderer *sdl.Renderer

	running bool

	currentScene scene.Scene = &DemoScene{}

	tick int // How many ticks has passed since the game was launched.
)

// Init is used for a starting procedure of the game which will do things
// such as creating the SDL Window, creating the SDL Renderer, Start the Scene, etc.
func Init(title string, x, y, width, height int, fullscreen bool) {
	var flags uint32

	if fullscreen {
		flags = sdl.WINDOW_FULLSCREEN
	}

	err := sdl.Init(sdl.INIT_EVERYTHING)

	if err != nil {
		log.Fatalln("Game // Error whie initialising SDL")
		log.Panic(err)
	}

	log.Println("Game // SDL Initialised")

	win, err := sdl.CreateWindow(title, x, y, width, height, flags)

	if err != nil {
		log.Fatalln("Game // Error while creating window")
		log.Panic(err)
	}

	log.Println("Game // Created Window")

	ren, err := sdl.CreateRenderer(win, -1, 0)

	if err != nil {
		log.Fatalln("Game // Error while creating renderer")
		log.Panic(err)
	}

	ren.SetDrawColor(255, 255, 255, 255)

	log.Println("Game // Created Renderer")

	log.Println("Game // Init Success")

	window = win
	renderer = ren

	input.Init()

	running = true

	currentScene.Enter()
}

// Render is used for drawing something to the renderer. Which will show
// up on the screen.
func Render() {
	renderer.Clear()

	currentScene.Draw(renderer)

	renderer.Present()
}

// Update is used to update and tick the game's component such as
// Moving the enemy, Update the score, Do physics, etc.
func Update() {
	tick++

	currentScene.Update()
}

// HandleEvents is used for receiving and handling SDL Events such as
// Update the mouse location, Check if the user is requesting to close the game, etc.
func HandleEvents() {
	event := sdl.PollEvent()

	switch event.(type) {
	case *sdl.QuitEvent:
		log.Println("Game // Got Quit Event. Stopping Game ...")
		running = false
		break
	default:
		input.HandleEvents(event)
		currentScene.HandleEvents(event)
		break
	}
}

// Clean is used for cleaning the game's resources such as
// Freeing the Textures, Clean the used resources, etc.
func Clean() {
	log.Println("Game // Cleaning up ...")

	currentScene.Exit()
	texture.Clean()
	input.Clean()
	renderer.Destroy()
	window.Destroy()
}

// Running returns the game's 'running' status.
func Running() bool {
	return running
}

// Window returns the SDL Window that's being used.
func Window() *sdl.Window {
	return window
}

// Renderer returns the SDL Renderer that's being used for the Window.
func Renderer() *sdl.Renderer {
	return renderer
}

// Stop literally stops the game
func Stop() {
	running = false
}

// ChangeScene changes the scene and runs the proper functions
// ( Run Exit on Old Scene, and Enter on New Scene )
func ChangeScene(s scene.Scene) {
	currentScene.Exit()
	currentScene = s
	currentScene.Enter()
}
