package mongo

import (
	"context"
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
	var (
		limit  int
		offset int
		err    error
	)

	if req.GetLimit() != "" {
		limit, err = strconv.Atoi(req.GetLimit())
		if err != nil {
			return nil, err
		}
	} else {
		limit = 0
	}

	if req.GetOfset() != "" {
		offset, err = strconv.Atoi(req.GetOfset())
		if err != nil {
			return nil, err
		}
	} else {
		offset = 0
	}

	filter := bson.M{}
	if req.GetDate() != "" {
		date, err := time.Parse("2006-01-02", req.GetDate())
		if err != nil {
			return nil, err
		}
		filter["date"] = bson.M{"$gte": date}
	}

	findOptions := options.Find()
	if limit > 0 {
		findOptions.SetLimit(int64(limit))
	}
	findOptions.SetSkip(int64(offset))

	cursor, err := s.mongo.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	var events []*pb.Events
	for cursor.Next(ctx) {
		var event pb.Events
		if err := cursor.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	return &pb.GetUserEventsResponse{Events: events}, nil
}
func (s *TimelineRepo) SearchTimeLine(ctx context.Context, req *pb.SearchTimeLineRequest) (*pb.SearchTimeLineResponse, error) {
	startDate, err := time.Parse("2006-01-02", req.GetStartDate())
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02", req.GetEndDate())
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"date": bson.M{
			"$gte": startDate,
			"$lte": endDate,		
		},
	}

	cursor, err := s.mongo.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var events []*pb.Events
	for cursor.Next(ctx) {
		var event pb.Events
		if err := cursor.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, &event)	
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
