package main

import (
	"log"
	"net"

	"github.com/Exam4/4th-month-exam-TimeLine-Service/config"
	"github.com/Exam4/4th-month-exam-TimeLine-Service/genproto"
	"github.com/Exam4/4th-month-exam-TimeLine-Service/service"
	"github.com/Exam4/4th-month-exam-TimeLine-Service/storage/mongo"
	"github.com/Exam4/4th-month-exam-TimeLine-Service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	Psql, err := postgres.DbConnection()
	if err != nil {
		
		panic(err)
	}

	mongo, err := mongo.SetupMongoDBConnection()
	if err != nil {
		panic(err)
	}

	liss, err := net.Listen("tcp", cfg.HTTPPort)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	genproto.RegisterHistoricalServiceServer(s, service.NewHistoricalService(mongo))
	genproto.RegisterCustomEventsServiceServer(s, service.NewCustomEventsService(Psql))
	genproto.RegisterMilestoneServiceServer(s, service.NewMilestoneService(Psql))
	genproto.RegisterTimeLineServiceServer(s, service.NewTimeLineService(mongo))

	log.Printf("Server started on port: %v", cfg.HTTPPort)
	if err := s.Serve(liss); err != nil {
		log.Fatal("error while serving: %v", err)
	}
}
