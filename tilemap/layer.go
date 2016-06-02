package tilemap

import (
	"encoding/csv"
	"strconv"
	"strings"
)

type Layer struct {
	Name string `xml:"name,attr"`

	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`

	Spacing int `xml:"spacing,attr"`
	Margin  int `xml:"margin,attr"`

	Data *Data `xml:"data"`

	Tiles [][]int
}

type Data struct {
	Encoding    string `xml:"encoding,attr"`
	Compression string `xml:"compression,attr"`

	Data string `xml:",chardata"`
}

func (l *Layer) Read() error {
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
