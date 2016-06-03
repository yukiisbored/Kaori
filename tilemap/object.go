package tilemap

// ObjectGroup is a group of tiled map objects
type ObjectGroup struct {
	Name    string    `xml:"name,attr"`
	Objects []*Object `xml:"object"`
}

// Object is a tile map object
// It can be an ellipse, rectangle, or a polyline
type Object struct {
	Id int `xml:"id,attr"`

	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`

	Gid int `xml:"gid,attr"`

	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`

	Polyline Polyline `xml:"polyline"`
}

// Polyline is a list of points representing a polygonal shape or line
type Polyline struct {
	Points string `xml:"points,attr"`
}
