package tilemap

import (
	"encoding/xml"

	"github.com/veandco/go-sdl2/sdl"
)

type Map struct {
	XMLName xml.Name `xml:"map"`

	Orientation string `xml:"orientation,attr"`
	RenderOrder string `xml:"renderorder,attr"`

	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`

	TileWidth  int `xml:"tilewidth,attr"`
	TileHeight int `xml:"tileheight,attr"`

	Tilesets []*Tileset `xml:"tileset"`

	Layers       []*Layer       `xml:"layer"`
	ObjectGroups []*ObjectGroup `xml:"objectgroup"`
}

func Unmarshal(data []byte, tilemap *Map) error {
	err := xml.Unmarshal(data, tilemap)

	for _, l := range tilemap.Layers {
		l.Parent = tilemap

		err := l.Read()

		if err != nil {
			return err
		}
	}

	return err
}

func Marshal(tilemap *Map) ([]byte, error) {
	return xml.Marshal(tilemap)
}

func (m *Map) Draw(renderer *sdl.Renderer, x, y int32) {
	for _, l := range m.Layers {
		l.Draw(renderer, x, y)
	}
}
