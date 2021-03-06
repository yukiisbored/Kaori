package tilemap

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/yukiisbored/Kaori/texture"
)

// Tileset is a set of sprites or textures for the tiled map
type Tileset struct {
	Name string `xml:"name,attr"`

	TileWidth  int `xml:"tilewidth,attr"`
	TileHeight int `xml:"tileheight,attr"`

	TileCount int `xml:"tilecount,attr"`

	Columns int `xml:"columns,attr"`

	Spacing int `xml:"spacing,attr"`
	Margin  int `xml:"margin,attr"`

	FirstID int `xml:"firstgid,attr"`

	Image Image `xml:"image"`
}

// Image is the actual image for the spritesheet or texture
type Image struct {
	Source string `xml:"source,attr"`

	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`
}

// Load loads the image needed as an SDL Texture via Kaori's texture package
func (t *Tileset) Load(renderer *sdl.Renderer, folder string) error {
	fileName := t.Image.Source
	path := folder + "/" + fileName

	return texture.Load(renderer, path, t.Name)
}

// Free free the loaded image texture
func (t *Tileset) Free() {
	texture.Free(t.Name)
}

// DrawTile draw a specific tile to the renderer at a specific location
// The tile number is a continuous number starting from the tile set's first id
func (t *Tileset) DrawTile(renderer *sdl.Renderer, x, y int32, tile int) {
	// Subtract the value by the tileset's first id
	tile = tile - t.FirstID + 1

	// Just ignore when it's lower than 0
	if tile < 0 {
		return
	}

	row := int32(math.Ceil(float64(tile / t.Columns)))
	frame := tile % t.Columns

	texture.DrawFrame(renderer, t.Name, x, y,
		int32(t.TileWidth), int32(t.TileHeight),
		int32(row), int32(frame),
		int32(t.Spacing), int32(t.Margin), 0, sdl.FLIP_NONE)
}
