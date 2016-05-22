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

	tick int
)

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

func Render() {
	renderer.Clear()

	currentScene.Draw(renderer)

	renderer.Present()
}

func Update() {
	tick++

	currentScene.Update()
}

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

func Clean() {
	log.Println("Game // Cleaning up ...")

	currentScene.Exit()
	texture.Clean()
	input.Clean()
	renderer.Destroy()
	window.Destroy()
}

func Running() bool {
	return running
}

func Window() *sdl.Window {
	return window
}

func Renderer() *sdl.Renderer {
	return renderer
}

func Stop() {
	running = false
}

func ChangeScene(s scene.Scene) {
	currentScene.Exit()
	currentScene = s
	currentScene.Enter()
}
