// Package texture provides a texture management system for Kaori
package texture

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

var textureMap map[string]*sdl.Texture = make(map[string]*sdl.Texture)

// Load loads the image into a Texture
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

// Draw draws the loaded texture
func Draw(renderer *sdl.Renderer, id string, x, y, width, height int32, angle float64, flip sdl.RendererFlip) {
	src := sdl.Rect{0, 0, width, height}
	dst := sdl.Rect{x, y, width, height}

	renderer.CopyEx(textureMap[id], &src, &dst, angle, nil, flip)
}

// DrawFrame draws a part of the loaded texture
// A Frame is simply a cropped section of the texture by the width and height provided
func DrawFrame(renderer *sdl.Renderer, id string, x, y, width, height, currentRow, currentFrame, spacing, margin int32, angle float64, flip sdl.RendererFlip) {
	src := sdl.Rect{margin + spacing*currentFrame + width*currentFrame, margin + height*(currentRow-1), width, height}
	dst := sdl.Rect{x, y, width, height}

	renderer.CopyEx(textureMap[id], &src, &dst, angle, nil, flip)
}

// Free frees the texture from memory
func Free(id string) {
	textureMap[id].Destroy()
	textureMap[id] = nil
}

// Clean frees every loaded texture
func Clean() {
	for _, v := range textureMap {
		v.Destroy()
	}
}
