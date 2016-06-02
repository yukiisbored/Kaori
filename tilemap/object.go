package tilemap

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
