package media

import "encoding/xml"

// Anime is a Metadata Container for all Animes
type Anime struct {
	XMLName xml.Name `xml:"anime" json:"-" yaml:"-"`    // Only for XML
	Titles  []Title  `xml:"titles>title" json:"titles"` // List of alternative titles
	Year    int      `xml:"year" json:"year" yaml:"year"`
	Genres  []string `xml:"genres" json:"genres"`
}
