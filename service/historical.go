package service

import (
	"context"

	"github.com/Exam4/4th-month-exam-TimeLine-Service/genproto"
	"github.com/Exam4/4th-month-exam-TimeLine-Service/storage"
)

type HistoricalService struct {
	storage storage.StorageMongo
	genproto.UnimplementedHistoricalServiceServer
}

func NewHistoricalService(storage storage.StorageMongo) *HistoricalService {
	return &HistoricalService{storage: storage}
}

func (s *HistoricalService) AddHistoricalEvent(ctx context.Context, req *genproto.AddHistoricalEventRequest) (*genproto.AddHistoricalEventResponse, error) {
	return s.storage.Historical().AddHistoricalEvent(ctx, req)
}


func (s *HistoricalService) GetAllHistoricalEvents(ctx context.Context, req *genproto.GetAllHistoricalRequest) (*genproto.GetAllHistoricalResponse, error) {
	return s.storage.Historical().GetAllHistoricalEvents(ctx, req)
}

func (s *HistoricalService) DeleteHistoricalEvent(ctx context.Context, req *genproto.DeleteHistoricalEventRequest) (*genproto.DeleteHistoricalEventResponse, error) {
	return s.storage.Historical().DeleteHistoricalEvent(ctx, req)
}