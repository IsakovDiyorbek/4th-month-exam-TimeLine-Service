package mongo

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "github.com/Exam4/4th-month-exam-TimeLine-Service/genproto"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TimelineRepo struct {
	mongo *mongo.Collection
}

func NewTimelineRepo(mongo *mongo.Database) *TimelineRepo {
	return &TimelineRepo{mongo: mongo.Collection("timelines")}
}
func (s *TimelineRepo) AddTimeLine(ctx context.Context, req *pb.AddTimeLineRequest) (*pb.AddTimeLineResponse, error) {
	req.Id = uuid.NewString()

	timeline := pb.TimeLine{
		Id:     uuid.NewString(),
		UserId: req.UserId,
		Events: []*pb.Events{
			{
				Id:      uuid.NewString(),
				Title:   req.Events.Title,
				Type:    req.Events.Type,
				Date:    req.Events.Date,
				Preview: req.Events.Preview,
			},
		},
		LastUpdated: req.LastUpdated,
	}

	_, err := s.mongo.InsertOne(context.TODO(), timeline)
	if err != nil {
		return nil, err
	}

	return &pb.AddTimeLineResponse{}, nil
}

func (s *TimelineRepo) GetEvent(ctx context.Context, req *pb.GetUserEventsRequest) (*pb.GetUserEventsResponse, error) {
	filter := bson.D{}
	opts := options.Find()

	if req.Limit != "" {
		limit, err := strconv.ParseInt(req.Limit, 10, 64)
		if err != nil {
			return nil, err
		}
		opts.SetLimit(limit)
	}
	if req.Ofset != "" {
		offset, err := strconv.ParseInt(req.Ofset, 10, 64)
		if err != nil {
			return nil, err
		}
		opts.SetSkip(offset)
	}
	if req.Date != "" {
		date, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			return nil, fmt.Errorf("invalid date format: %v", err)
		}
		filter = append(filter, bson.E{"date", date})
	}

	cursor, err := s.mongo.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var events []*pb.Events
	for cursor.Next(ctx) {
		var event pb.Events
		err := cursor.Decode(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.GetUserEventsResponse{Events: events}, nil
}

func (s *TimelineRepo) SearchTimeLine(ctx context.Context, req *pb.SearchTimeLineRequest) (*pb.SearchTimeLineResponse, error) {
	filter := bson.D{}
	opts := options.Find()

	var startDate, endDate time.Time
	var err error
	if req.StartDate != "" {
		startDate, err = time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			return nil, fmt.Errorf("invalid start_date format: %v", err)
		}
		filter = append(filter, bson.E{"date", bson.D{{"$gte", startDate}}})
	}
	if req.EndDate != "" {
		endDate, err = time.Parse("2006-01-02", req.EndDate)
		if err != nil {	
			return nil, fmt.Errorf("invalid end_date format: %v", err)
		}
		filter = append(filter, bson.E{"date", bson.D{{"$lte", endDate}}})
	}
	
	cursor, err := s.mongo.Find(ctx, filter, opts)
	if err != nil {
		log.Printf("Error while searching timeline: %v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var events []*pb.Events
	for cursor.Next(ctx) {
		var event pb.Events
		err := cursor.Decode(&event)
		if err != nil {
			log.Printf("Error while decoding event: %v\n", err)
			return nil, err
		}
		events = append(events, &event)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.SearchTimeLineResponse{Event: events}, nil
}
func (s *TimelineRepo) DeleteTimeLine(ctx context.Context, req *pb.DeleteTimeLineRequest) (*pb.DeleteTimeLineResponse, error) {
	filter := bson.D{{"id", req.Id}}

	_, err := s.mongo.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteTimeLineResponse{}, nil
}
