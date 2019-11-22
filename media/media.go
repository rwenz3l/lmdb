package media

// Title is a Wrapper for all Variants
// XML Output = <title type="..." language="...">...</title>
type Title struct {
	Title    string `xml:",chardata" json:"title"`
	Type     string `xml:"type,attr" json:"type"`
	Language string `xml:"language,attr" json:"language"`
}

// Repository provides an Interface for SQL Data I/O
type Repository struct {
}
