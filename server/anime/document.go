package anime

import "go.mongodb.org/mongo-driver/bson/primitive"

// Document defines the MongoDB Collection Document Structure

// Title defines the Name of an Anime, that varies in Type and Language
type Title struct {
	Title    string `bson:"title"`
	Type     string `bson:"type"`
	Language string `bson:"language"`
}

// Anime is a Document stored in MongoDB
type Anime struct {
	ID     *primitive.ObjectID `bson:"_id,omitempty"`
	Titles []Title             `bson:"titles"` // List of alternative titles
	Year   int                 `bson:"year"`
	Genres []string            `bson:"genres"`
}
