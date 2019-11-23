package anime

import (
	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/mongo"
)

// Service provides necessary dependencies
type Service struct {
	logger     log.Logger
	collection *mongo.Collection
}

// NewService creates a new Instance
func NewService(logger log.Logger, collection *mongo.Collection) *Service {
	svc := &Service{
		logger:     logger,
		collection: collection,
	}
	svc.initIndex()
	return svc
}
