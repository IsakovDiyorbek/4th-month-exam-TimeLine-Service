package service

import (
	"context"

	"github.com/Exam4/4th-month-exam-TimeLine-Service/genproto"
	"github.com/Exam4/4th-month-exam-TimeLine-Service/storage"
)

type TimeLineService struct {
	storage storage.StorageMongo
	genproto.UnimplementedTimeLineServiceServer
}

func NewTimeLineService(storage storage.StorageMongo) *TimeLineService {
	return &TimeLineService{storage: storage}
}

func (s *TimeLineService) AddTimeLine(ctx context.Context, req *genproto.AddTimeLineRequest) (*genproto.AddTimeLineResponse, error) {
	return s.storage.TimeLine().AddTimeLine(ctx, req)
}

func (s *TimeLineService) GetEvent(ctx context.Context, req *genproto.GetUserEventsRequest) (*genproto.GetUserEventsResponse, error) {
	return s.storage.TimeLine().GetEvent(ctx, req)
}

func (s *TimeLineService) SearchTimeLine(ctx context.Context, req *genproto.SearchTimeLineRequest) (*genproto.SearchTimeLineResponse, error) {
	return s.storage.TimeLine().SearchTimeLine(ctx, req)
}

func (s *TimeLineService) DeleteTimeLine(ctx context.Context, req *genproto.DeleteTimeLineRequest) (*genproto.DeleteTimeLineResponse, error) {
	return s.storage.TimeLine().DeleteTimeLine(ctx, req)
}
