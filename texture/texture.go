package texture

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

var textureMap map[string]*sdl.Texture

func Load(renderer *sdl.Renderer, fileName, id string) error {
	tempSurface, err := img.Load(fileName)

	if err != nil {
		return err
	}

	defer tempSurface.Free()

	texture, err := renderer.CreateTextureFromSurface(tempSurface)

	if err != nil {
		return err
	}

	textureMap[id] = texture

	return nil
}

func Draw(renderer *sdl.Renderer, id string, x, y, width, height int32, flip sdl.RendererFlip) {
	src := sdl.Rect{0, 0, width, height}
	dst := sdl.Rect{x, y, width, height}

	renderer.CopyEx(textureMap[id], &src, &dst, 0, nil, flip)
}

func DrawFrame(renderer *sdl.Renderer, id string, x, y, width, height, currentRow, currentFrame int32, flip sdl.RendererFlip) {
	src := sdl.Rect{width * currentFrame, height * (currentRow - 1), width, height}
	dst := sdl.Rect{x, y, width, height}

	renderer.CopyEx(textureMap[id], &src, &dst, 0, nil, flip)
}

func Free(id string) {
	textureMap[id].Destroy()
}

func Clean() {
	for _, v := range textureMap {
		v.Destroy()
	}
}
