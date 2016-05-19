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

	currentScene scene.Scene = EmptyScene{}
)

func Init(title string, xpos, ypos, height, width int, fullscreen bool) {
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

	w, err := sdl.CreateWindow(title, xpos, ypos, height, width, flags)

	if err != nil {
		log.Fatalln("Game // Error while creating window")
		log.Panic(err)
	}

	log.Println("Game // Created Window")

	r, err := sdl.CreateRenderer(w, -1, 0)

	if err != nil {
		log.Fatalln("Game // Error while creating renderer")
		log.Panic(err)
	}

	r.SetDrawColor(0, 0, 0, 0)

	log.Println("Game // Created Renderer")

	log.Println("Game // Init Success")

	window = w
	renderer = r

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
