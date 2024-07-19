package mongo

import (
	"context"

	"github.com/Exam4/4th-month-exam-TimeLine-Service/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StorageM struct {
	mongo      *mongo.Database
	timeline   storage.TimeLineI
	historical storage.HistoricalI
}

func SetupMongoDBConnection() (storage.StorageMongo, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	mongo := client.Database("memories")

	time := NewTimelineRepo(mongo)
	history := NewHistoricalRepo(mongo)
	return &StorageM{
		mongo:      mongo,
		timeline:   time,
		historical: history,
	}, nil
}

func (s *StorageM) TimeLine() storage.TimeLineI {
	if s.timeline == nil {
		s.timeline = NewTimelineRepo(s.mongo)
	}
	return s.timeline
}

func (s *StorageM) Historical() storage.HistoricalI {
	if s.historical == nil {
		s.historical = NewHistoricalRepo(s.mongo)
	}
	return s.historical
}
