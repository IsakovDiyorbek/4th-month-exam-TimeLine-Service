package storage

import (
	"context"

	pb "github.com/Exam4/4th-month-exam-TimeLine-Service/genproto"
)

type StorageI interface {
	CustomEvent() CustomEventsI
	Milestone() MilestoneI
}

type StorageMongo interface {
	Historical() HistoricalI
	TimeLine() TimeLineI
}

type CustomEventsI interface {
	AddCustomEvent(ctx context.Context, req *pb.AddCustomEventRequest) (*pb.AddCustomEventResponse, error)
	UpdateCustomEvent(ctx context.Context, req *pb.UpdateCustomEventsRequest) (*pb.UpdateCustomEventsResponse, error)
	DeleteCustomEvent(ctx context.Context, req *pb.DeleteCustomEventsRequest) (*pb.DeleteCustomEventsResponse, error)
	GetAllCustomEvents(ctx context.Context, req *pb.GetAllEventsRequest) (*pb.GetAllEventsResponse, error)
	GetByIdCustomEvent(ctx context.Context, req *pb.GetByIdEvetsRequest) (*pb.GetByIdEvetsResponse, error)
}

type MilestoneI interface {
	AddMilestone(ctx context.Context, req *pb.AddMilestonesRequest) (*pb.AddMilestonesResponse, error)
	GetAllMilestone(ctx context.Context, req *pb.GetAllMilestonesRequest) (*pb.GetAllMilestonesResponse, error)
	UpdateMilestone(ctx context.Context, req *pb.UpdateMilestonesRequest) (*pb.UpdateMilestonesResponse, error)
	DeleteMilestone(ctx context.Context, req *pb.DeleteMilestonesRequest) (*pb.DeleteMilestonesResponse, error)
}

type HistoricalI interface {
	GetAllHistoricalEvents(ctx context.Context, req *pb.GetAllHistoricalRequest) (*pb.GetAllHistoricalResponse, error)
	AddHistoricalEvent(ctx context.Context, req *pb.AddHistoricalEventRequest) (*pb.AddHistoricalEventResponse, error)
	DeleteHistoricalEvent(ctx context.Context, req *pb.DeleteHistoricalEventRequest) (*pb.DeleteHistoricalEventResponse, error);
}

type TimeLineI interface {
	AddTimeLine(ctx context.Context, req *pb.AddTimeLineRequest) (*pb.AddTimeLineResponse, error)
	GetEvent(ctx context.Context, req *pb.GetUserEventsRequest) (*pb.GetUserEventsResponse, error)
	SearchTimeLine(ctx context.Context, req *pb.SearchTimeLineRequest) (*pb.SearchTimeLineResponse, error)
	DeleteTimeLine(ctx context.Context, req *pb.DeleteTimeLineRequest) (*pb.DeleteTimeLineResponse, error)
}
