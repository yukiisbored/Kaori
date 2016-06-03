package tilemap

import (
	"encoding/csv"
	"strconv"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

// Layer is a structure for TMX Tilemap Layers
type Layer struct {
	Parent *Map

	Name string `xml:"name,attr"`

	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`

	Spacing int `xml:"spacing,attr"`
	Margin  int `xml:"margin,attr"`

	Data *Data `xml:"data"`

	Tiles [][]int `xml:"-"`
}

// Data is a structure for Layer's Tile Data
type Data struct {
	Encoding    string `xml:"encoding,attr"`
	Compression string `xml:"compression,attr"`

	Data string `xml:",chardata"`
}

// Read parses the layer's tile data into two dimension array of integers
func (l *Layer) Read() error {
	// Support CSV for now
	// TODO: XML, base64, base64 + zlib
	return readCSV(l)
}

func readCSV(l *Layer) error {
	// Tile Maps are weird
	// Horrible.
	raw := l.Data.Data[0:len(l.Data.Data)-1] + ","

	rdr := csv.NewReader(strings.NewReader(raw))

	rdr.TrimLeadingSpace = true

	records, err := rdr.ReadAll()

	if err != nil {
		return err
	}

	l.Tiles = make([][]int, len(records))

	for i, row := range records {
		l.Tiles[i] = make([]int, len(records[i])-1)

		for j, col := range row {
			// I hate this thing, but it's the only thing I can think of right now
			if j == len(records[i])-1 {
				break
			}

			tile, err := strconv.Atoi(col)

			if err != nil {
				return err
			}

			l.Tiles[i][j] = tile
		}
	}

	return nil
}

// Draw draws the layer to the renderer with the starting location of x and y
func (l *Layer) Draw(renderer *sdl.Renderer, x, y int32) {
	for yTile, r := range l.Tiles {
		for xTile, t := range r {
			var tileset *Tileset = l.Parent.Tilesets[0]

			for _, ts := range l.Parent.Tilesets {
				if ts.FirstID > t {
					continue
				}

				tileset = ts
				return
			}

			tileset.DrawTile(renderer, x+int32(xTile*tileset.TileWidth), y+int32(yTile*tileset.TileHeight), t)
		}
	}
}
