package anime

import (
	"context"
	"log"

	"github.com/go-kit/kit/log/level"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB Collection Interface

func (svc *Service) initIndex() {
	// Set Options for the Index
	opt := options.Index()
	opt.SetName("title_search")
	opt.SetWeights(bson.M{
		"titles.title": 1, // Index each Title: { "titles": [{"title": "..."}]}
	})
	// Read: https://stackoverflow.com/a/24905566
	opt.SetDefaultLanguage("en")  // Set Default Language for all languages MongoDB does not support
	opt.SetLanguageOverride("en") // Override language Field for unsupported languages to "en"

	// MongoDB indexes the deepest key (titles.title => title)
	// So we specify 'title' as key and want that as "text" so it gets properly
	// tokenized
	index := mongo.IndexModel{Keys: bson.M{
		"title": "text",
	}, Options: opt}

	// Create the Index
	if _, err := svc.collection.Indexes().CreateOne(context.TODO(), index); err != nil {
		log.Println("Could not create text index:", err)
	}
}

// Create adds a new Anime to the MongoDB Collection
func (svc *Service) Create(anime Anime) {
	res, _ := svc.collection.InsertOne(nil, anime)
	level.Debug(svc.logger).Log("msg", "Inserted new Anime", "_id", res.InsertedID)
}

// Read retrieves an Anime from the MongoDB Collection
func (svc *Service) Read(id int) {
	svc.logger.Log("Read in Collection")
	return
}

// Update upserts the Anime in the Collection
func (svc *Service) Update(anime Anime) {
	svc.logger.Log("Update in Collection")
	return
}

// Delete removes the Anime from the Collection
func (svc *Service) Delete(id int) {
	svc.logger.Log("Delete in Collection")
	return
}

// Find Searches for Animes matching the query
// XXX: When searching with an index, we only get the ObjectIDs in return.
// We Need to fix that
func (svc *Service) Find(t string) []Anime {
	// https://wb.id.au/computer/golang-and-mongodb-using-the-mongo-go-driver/
	ctx := context.TODO()

	// Empty Slice for output
	animes := []Anime{}

	// Create a Flter with t(=term)
	filter := bson.M{"$text": bson.M{"$search": t}}
	findOptions := options.Find()
	findOptions.SetLimit(100) // Limit?
	// XXX: Projection allows us to show scoring, if we need that..!
	// This might be useful for ranked search results!
	// findOptions.SetProjection(bson.M{
	// 	"titles": 1,
	// 	"score":  bson.M{"$meta": "textScore"},
	// })
	// findOptions.SetSort(bson.M{"score": bson.M{"$meta": "textScore"}})
	cur, err := svc.collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatalln("Error during text search:", err)
	}
	defer cur.Close(ctx)

	//Loops over the cursor stream and appends to result array
	for cur.Next(ctx) {
		var a Anime
		cur.Decode(&a)
		animes = append(animes, a)
	}

	if err != nil {
		level.Error(svc.logger).Log("err", err)
	}
	return animes
}

// FindAll retrieves all Animes
func (svc *Service) FindAll() []*Anime {
	// https://wb.id.au/computer/golang-and-mongodb-using-the-mongo-go-driver/
	ctx := context.TODO()

	// Empty Slice for output
	animes := []*Anime{}

	findOptions := options.Find()
	findOptions.SetLimit(100)
	// findOptions.SetSort(bson.M{"score": bson.M{"$meta": "textScore"}})
	cur, err := svc.collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatalln("Error during text search:", err)
	}
	defer cur.Close(ctx)
	err = cur.All(nil, &animes)
	if err != nil {
		level.Error(svc.logger).Log("err", err)
	}
	return animes
}
