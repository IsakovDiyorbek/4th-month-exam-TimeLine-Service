package service

import (
	"context"

	"github.com/Exam4/4th-month-exam-TimeLine-Service/genproto"
	"github.com/Exam4/4th-month-exam-TimeLine-Service/storage"
)

type MilestoneService struct {
	storage storage.StorageI
	genproto.UnimplementedMilestoneServiceServer
}

func NewMilestoneService(storage storage.StorageI) *MilestoneService {
	return &MilestoneService{storage: storage}
}

func (s *MilestoneService) AddMilestone(ctx context.Context, req *genproto.AddMilestonesRequest) (*genproto.AddMilestonesResponse, error) {
	return s.storage.Milestone().AddMilestone(ctx, req)
}

func (s *MilestoneService) GetAllMilestone(ctx context.Context, req *genproto.GetAllMilestonesRequest) (*genproto.GetAllMilestonesResponse, error) {
	return s.storage.Milestone().GetAllMilestone(ctx, req)
}

func (s *MilestoneService) UpdateMilestone(ctx context.Context, req *genproto.UpdateMilestonesRequest) (*genproto.UpdateMilestonesResponse, error) {
	return s.storage.Milestone().UpdateMilestone(ctx, req)
}

func (s *MilestoneService) DeleteMilestone(ctx context.Context, req *genproto.DeleteMilestonesRequest) (*genproto.DeleteMilestonesResponse, error) {
	return s.storage.Milestone().DeleteMilestone(ctx, req)
}
