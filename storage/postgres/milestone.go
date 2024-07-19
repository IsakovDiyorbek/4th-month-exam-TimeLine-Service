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

type MilestoneRepo struct {
	Db *sql.DB
}

func NewMilestoneRepo(db *sql.DB) *MilestoneRepo {
	return &MilestoneRepo{Db: db}
}

func (m *MilestoneRepo) AddMilestone(ctx context.Context, req *pb.AddMilestonesRequest) (*pb.AddMilestonesResponse, error) {
	query := `insert into milestones(id, user_id, title, date, category) values($1, $2, $3, $4, $5)`

	req.Id = uuid.NewString()
	_, err := m.Db.ExecContext(ctx, query, req.Id, req.UserId, req.Title, req.Date, req.Category)
	if err != nil {
		log.Printf("Error while creating milestone: %v\n", err)
		return nil, err
	}
	return &pb.AddMilestonesResponse{}, nil
}

func (m *MilestoneRepo) GetAllMilestone(ctx context.Context, req *pb.GetAllMilestonesRequest) (*pb.GetAllMilestonesResponse, error) {
	query := `select id, user_id, title, date, category from milestones`

	param := make(map[string]interface{})
	filter := ` where deleted_at = 0`
	if req.UserId != "" {
		param["user_id"] = req.UserId
		filter += ` and user_id = :user_id`
	}

	if req.Category != "" {
		param["category"] = req.Category
		filter += ` and category = :category`
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

	rows, err := m.Db.QueryContext(ctx, query, arr...)
	if err != nil {
		log.Printf("Error while getting all milestones: %v\n", err)
		return nil, err
	}

	defer rows.Close()
	var milestones []*pb.Milestones
	for rows.Next() {
		var milestone pb.Milestones
		err := rows.Scan(&milestone.Id, &milestone.UserId, &milestone.Title, &milestone.Date, &milestone.Category)
		if err != nil {
			log.Printf("Error while scanning milestone: %v\n", err)
			return nil, err
		}
		milestones = append(milestones, &milestone)
	}
	return &pb.GetAllMilestonesResponse{Milestone: milestones}, nil
}


func (m *MilestoneRepo) UpdateMilestone(ctx context.Context, req *pb.UpdateMilestonesRequest) (*pb.UpdateMilestonesResponse, error) {
	query := `update milestones set title = $1, date = $2  where id = $3`
	_, err := m.Db.ExecContext(ctx, query, req.Title, req.Date, req.Id)
	if err != nil {
		log.Printf("Error while updating milestone: %v\n", err)
		return nil, err
	}
	return &pb.UpdateMilestonesResponse{}, nil
}

func (m *MilestoneRepo) DeleteMilestone(ctx context.Context, req *pb.DeleteMilestonesRequest) (*pb.DeleteMilestonesResponse, error) {
	query := `update milestones set deleted_at = $1 where id = $2`
	_, err := m.Db.ExecContext(ctx, query, time.Now().Unix(), req.Id)
	if err != nil {
		log.Printf("Error while deleting milestone: %v\n", err)
		return nil, err
	}
	return &pb.DeleteMilestonesResponse{}, nil
}
