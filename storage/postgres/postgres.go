package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Exam4/4th-month-exam-TimeLine-Service/storage"
	"github.com/Exam4/4th-month-exam-TimeLine-Service/config"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db         *sql.DB
	CustomEvents storage.CustomEventsI
	Milestones   storage.MilestoneI
}

func DbConnection() (storage.StorageI, error) {
	cfg := config.Load()
	con := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresDatabase, cfg.PostgresPassword, cfg.PostgresPort)
	db, err := sql.Open("postgres", con)
	if err != nil {
		log.Fatal("Error while db connection", err)
		return nil, nil
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error while db ping connection", err)
		return nil, nil

	}
	return &Storage{Db: db}, nil

}


func (s *Storage) CustomEvent() storage.CustomEventsI {
	if s.CustomEvents == nil {
		s.CustomEvents = NewCustomEventsRepo(s.Db)
	}
	return s.CustomEvents
}

func (s *Storage) Milestone() storage.MilestoneI {
	if s.Milestones == nil {
		s.Milestones = NewMilestoneRepo(s.Db)
	}
	return s.Milestones
}





