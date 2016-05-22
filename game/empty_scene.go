package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

type EmptyScene struct {
}

func (s *EmptyScene) Enter() {
}

func (s *EmptyScene) Update() {
}

func (s *EmptyScene) Draw(r *sdl.Renderer) {
}

func (s *EmptyScene) HandleEvents(e sdl.Event) {
}

func (s *EmptyScene) Exit() {
}
