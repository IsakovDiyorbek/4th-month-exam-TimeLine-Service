package service

import (
	"context"

	"github.com/Exam4/4th-month-exam-TimeLine-Service/genproto"
	"github.com/Exam4/4th-month-exam-TimeLine-Service/storage"
)

type CustomEventsService struct {
	storage storage.StorageI
	genproto.UnimplementedCustomEventsServiceServer
}

func NewCustomEventsService(storage storage.StorageI) *CustomEventsService {
	return &CustomEventsService{storage: storage}
}

func (s *CustomEventsService) AddCustomEvent(ctx context.Context, req *genproto.AddCustomEventRequest) (*genproto.AddCustomEventResponse, error) {
	return s.storage.CustomEvent().AddCustomEvent(ctx, req)
}

func (s *CustomEventsService) UpdateCustomEvent(ctx context.Context, req *genproto.UpdateCustomEventsRequest) (*genproto.UpdateCustomEventsResponse, error) {
	return s.storage.CustomEvent().UpdateCustomEvent(ctx, req)
}

func (s *CustomEventsService) DeleteCustomEvent(ctx context.Context, req *genproto.DeleteCustomEventsRequest) (*genproto.DeleteCustomEventsResponse, error) {
	return s.storage.CustomEvent().DeleteCustomEvent(ctx, req)
}

func (s *CustomEventsService) GetAllCustomEvents(ctx context.Context, req *genproto.GetAllEventsRequest) (*genproto.GetAllEventsResponse, error) {
	return s.storage.CustomEvent().GetAllCustomEvents(ctx, req)
}

func (s *CustomEventsService) GetByIdCustomEvent(ctx context.Context, req *genproto.GetByIdEvetsRequest) (*genproto.GetByIdEvetsResponse, error) {
	return s.storage.CustomEvent().GetByIdCustomEvent(ctx, req)
}

