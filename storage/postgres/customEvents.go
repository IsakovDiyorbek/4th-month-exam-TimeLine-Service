package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

	pb "github.com/Exam4/4th-month-exam-TimeLine-Service/genproto"
	"github.com/Exam4/4th-month-exam-TimeLine-Service/helper"
	"github.com/google/uuid"
)

type CustomEventsRepo struct {
	db *sql.DB
}

func NewCustomEventsRepo(db *sql.DB) *CustomEventsRepo {
	return &CustomEventsRepo{db: db}
}

func (c *CustomEventsRepo) AddCustomEvent(ctx context.Context, req *pb.AddCustomEventRequest) (*pb.AddCustomEventResponse, error) {
	query := `insert into custom_events(id, user_id, title, description, date, category) values($1, $2, $3, $4, $5, $6)`

	req.Id = uuid.NewString()
	_, err := c.db.ExecContext(ctx, query, req.Id, req.UserId, req.Title, req.Description, req.Date, req.Category)
	if err != nil {
		log.Printf("Error while creating milestone: %v\n", err)
		return nil, err
	}
	return &pb.AddCustomEventResponse{}, nil
}


func (c *CustomEventsRepo) UpdateCustomEvent(ctx context.Context, req *pb.UpdateCustomEventsRequest) (*pb.UpdateCustomEventsResponse, error) {
	query := `update custom_events set title = $1, description = $2 where id = $3`
	_, err := c.db.ExecContext(ctx, query, req.Title, req.Description, req.EventId)
	if err != nil {
		log.Printf("Error while updating custom event: %v\n", err)
		return nil, err
	}
	return &pb.UpdateCustomEventsResponse{}, nil
}

func (c *CustomEventsRepo) DeleteCustomEvent(ctx context.Context, req *pb.DeleteCustomEventsRequest) (*pb.DeleteCustomEventsResponse, error) {
	query := `update custom_events set deleted_at = $1 where id = $2`
	_, err := c.db.ExecContext(ctx, query, time.Now().Unix(), req.EventId)
	if err != nil {
		log.Printf("Error while deleting custom event: %v\n", err)
		return nil, err
	}
	return &pb.DeleteCustomEventsResponse{}, nil
}

func (c *CustomEventsRepo) GetAllCustomEvents(ctx context.Context, req *pb.GetAllEventsRequest) (*pb.GetAllEventsResponse, error) {
	query := `select id, user_id, title, description, date, category, created_at from custom_events`

	param := make(map[string]interface{})
	filter := ` where deleted_at = 0`
	if req.UserId != "" {
		param["user_id"] = req.UserId
		filter += ` and user_id = :user_id`
	}

	if req.Date != "" {
		param["date"] = req.Date
		filter += ` and date = :date`
	}

	if req.Title != "" {
		param["title"] = req.Title
		filter += ` and title = :title`
	}


	query += filter

	query, arr := helper.ReplaceQueryParams(query, param)

	rows, err := c.db.QueryContext(ctx, query, arr...)
	if err != nil {
		log.Printf("Error while getting all custom events: %v\n", err)
		return nil, err
	}

	defer rows.Close()
	var customEvents []*pb.CustomEvents
	for rows.Next() {
		var customEvent pb.CustomEvents
		err := rows.Scan(&customEvent.Id, &customEvent.UserId, &customEvent.Title, &customEvent.Description, &customEvent.Date, &customEvent.Category, &customEvent.CreatedAt)
		if err != nil {
			log.Printf("Error while scanning custom event: %v\n", err)
			return nil, err
		}
		customEvents = append(customEvents, &customEvent)
	}
	return &pb.GetAllEventsResponse{Event: customEvents}, nil
}


func (c *CustomEventsRepo) GetByIdCustomEvent(ctx context.Context, req *pb.GetByIdEvetsRequest) (*pb.GetByIdEvetsResponse, error){
	query := `select id, user_id, title, description, date, category, created_at from custom_events where id = $1`
	var customEvent pb.CustomEvents
	err := c.db.QueryRowContext(ctx, query, req.Id).Scan(&customEvent.Id, &customEvent.UserId, &customEvent.Title, &customEvent.Description, &customEvent.Date, &customEvent.Category, &customEvent.CreatedAt)
	if err != nil {
		log.Printf("Error while getting custom event by id: %v\n", err)
		return nil, err
	}
	return &pb.GetByIdEvetsResponse{Event: &customEvent}, nil
}
