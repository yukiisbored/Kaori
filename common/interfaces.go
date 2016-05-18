package common

import "github.com/veandco/go-sdl2/sdl"

type Updater interface {
	Update()
}

type Drawer interface {
	Draw(*sdl.Renderer)
}

type EventHandler interface {
	HandleEvents(sdl.Event)
}

type Cleaner interface {
	Clean()
}

type GameObject interface {
	Updater
	Drawer
	Cleaner
}
