package tilemap

import (
	"encoding/csv"
	"encoding/xml"
	"strconv"
	"strings"
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

type Tileset struct {
	Name string `xml:"name,attr"`

	TileWidth  int `xml:"tilewidth,attr"`
	TileHeight int `xml:"tileheight,attr"`

	TileCount int `xml:"tilecount,attr"`

	Columns int `xml:"columns,attr"`

	Image Image `xml:"image"`
}

type Image struct {
	Source string `xml:"source,attr"`

	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`
}

type Layer struct {
	Name string `xml:"name,attr"`

	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`

	Data Data `xml:"data"`

	Tiles [][]int
}

type Data struct {
	Encoding    string `xml:"encoding,attr"`
	Compression string `xml:"compression,attr"`

	Data string `xml:",chardata"`
}

type ObjectGroup struct {
	Name    string    `xml:"name,attr"`
	Objects []*Object `xml:"object"`
}

type Object struct {
	Id int `xml:"id,attr"`

	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`

	Gid int `xml:"gid,attr"`

	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`

	Polyline Polyline `xml:"polyline"`
}

type Polyline struct {
	Points string `xml:"points,attr"`
}

func Unmarshal(data []byte, tilemap *Map) error {
	err := xml.Unmarshal(data, tilemap)

	for _, l := range tilemap.Layers {
		l.Read()
	}

	return err
}

func Marshal(tilemap Map) ([]byte, error) {
	return xml.Marshal(tilemap)
}

func (l *Layer) Read() error {
	raw := l.Data.Data

	rdr := csv.NewReader(strings.NewReader(raw))

	rdr.TrimLeadingSpace = true

	records, err := rdr.ReadAll()

	if err != nil {
		return err
	}

	l.Tiles = make([][]int, len(records))

	for i, row := range records {
		l.Tiles[i] = make([]int, len(records[i]))

		for j, col := range row {
			tile, err := strconv.Atoi(col)

			if err != nil {
				return nil
			}

			l.Tiles[i][j] = tile
		}
	}

	return nil
}
