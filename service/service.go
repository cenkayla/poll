package service

import (
	"log"
	"net/http"
	"os"

	"github.com/cenkayla/votingservice/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func (a *Service) Open() error {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error when loading .env file")
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE")), &gorm.Config{})
	if err != nil {
		return err
	}
	a.DB = db
	a.ConfigureRouter()
	return a.DB.AutoMigrate(&models.Poll{}, &models.Choice{})
}

func (s *Service) ConfigureRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/createpoll", s.createPoll).Methods("POST")
	r.HandleFunc("/api/poll", s.votePoll).Methods("POST")
	r.HandleFunc("/api/getresult", s.getPoll).Methods("POST")
	return r
}
