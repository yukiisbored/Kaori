package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

// EmptyScene is a skeleton for a Scene
type EmptyScene struct {
}

// Enter is used when the Scene is going to be used
// usually this function is used for loading assets, transitions, etc
func (s *EmptyScene) Enter() {
}

// Update is used when a game update event happens
// usually this function is used for moving the objects, check the input devices, etc
func (s *EmptyScene) Update() {
}

// Draw is used when a game render event happens
// usually this function is used for drawing the objects
func (s *EmptyScene) Draw(r *sdl.Renderer) {
}

// HandleEvents is used when a certain SDL Event is received
func (s *EmptyScene) HandleEvents(e sdl.Event) {
}

// Exit is used when the Scene is going to change or the game is going to close
// ( You can check what's going on with the running variable ) usually this
// function is used for freeing resources, cleaning objects, transitions, etc
func (s *EmptyScene) Exit() {
}
