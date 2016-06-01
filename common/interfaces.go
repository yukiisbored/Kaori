// Package common provides shared interfaces / structures
// between Kaori's Modules
package common

import "github.com/veandco/go-sdl2/sdl"

// Updater provides a function which will be used when
// a game update event happens
type Updater interface {
	Update()
}

// Drawer provides a Draw function which will be used when
// a game draw event happens.
// This function also provides access to the game's renderer
type Drawer interface {
	Draw(*sdl.Renderer)
}

// EventHandler provides a function which will be used when
// a certain SDL Event is received
type EventHandler interface {
	HandleEvents(sdl.Event)
}

// Cleaner provides a function which will be used when
// a game is closing. This function should be used for
// freeing SDL's objects and/or cleaning resources
type Cleaner interface {
	Clean()
}

// GameObject provides the functions which are needed
// for a game object.
type GameObject interface {
	Updater
	Drawer
	Cleaner
}
