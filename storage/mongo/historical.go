package mongo

import (
	"context"
	"log"
	"strconv"

	pb "github.com/Exam4/4th-month-exam-TimeLine-Service/genproto"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HistoricalRepo struct {
	client *mongo.Collection
}

func NewHistoricalRepo(mongo *mongo.Database) *HistoricalRepo {
	return &HistoricalRepo{client: mongo.Collection("historical")}
}

func (s *HistoricalRepo) AddHistoricalEvent(ctx context.Context, req *pb.AddHistoricalEventRequest) (*pb.AddHistoricalEventResponse, error) {
	event := pb.HistoricalEvents{
		Id:          uuid.NewString(),
		Title:       req.Title,
		Date:        req.Date,
		Category:    req.Category,
		Description: req.Description,
		SourceUrl:   req.SourceUrl,
		Time:        req.Time,
	}

	_, err := s.client.InsertOne(context.TODO(), event)
	if err != nil {
		log.Printf("Error inserting historical event: %v", err)
		return nil, err
	}

	return &pb.AddHistoricalEventResponse{}, nil
}

func (s *HistoricalRepo) GetAllHistoricalEvents(ctx context.Context, req *pb.GetAllHistoricalRequest) (*pb.GetAllHistoricalResponse, error) {
	filter := bson.D{}
	opts := options.Find()
	if req.Limit != "" {
		limit, err := strconv.ParseInt(req.Limit, 10, 64)
		if err != nil {
			log.Printf("Error parsing limit: %v", err)
			return nil, err
		}
		opts.SetLimit(limit)
	}
	if req.Ofset != "" {
		offset, err := strconv.ParseInt(req.Ofset, 10, 64)
		if err != nil {
			log.Printf("Error parsing offset: %v", err)
			return nil, err
		}
		opts.SetSkip(offset)
	}
	if req.Date != "" {
		filter = append(filter, bson.E{"date", req.Date})
	}

	cursor, err := s.client.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Printf("Error finding historical events: %v", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var events []*pb.HistoricalEvents
	for cursor.Next(context.TODO()) {
		var event pb.HistoricalEvents
		err := cursor.Decode(&event)
		if err != nil {
			log.Printf("Error decoding historical event: %v", err)
			return nil, err
		}
		events = append(events, &event)
	}

	return &pb.GetAllHistoricalResponse{Event: events}, nil
}

func (s *HistoricalRepo) DeleteHistoricalEvent(ctx context.Context, req *pb.DeleteHistoricalEventRequest) (*pb.DeleteHistoricalEventResponse, error) {

	filter := bson.D{{"id", req.Id}}

	_, err := s.client.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Printf("Error deleting historical event: %v", err)
		return nil, err
	}

	return &pb.DeleteHistoricalEventResponse{}, nil
}
