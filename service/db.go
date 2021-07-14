package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cenkayla/votingservice/models"
	"gorm.io/gorm"
)

func (s *Service) getPoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := models.Poll{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.DB.Preload("Choice").First(&p).Where("id = ?", p.ID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&p)
}

func (s *Service) votePoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	d := models.Choice{}

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.DB.Model(&d).Where("name = ? AND poll_id = ?", d.Name, d.PollID).Update("votes", gorm.Expr("votes + ?", 1)).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(fmt.Sprintf("Successfully voted to %v", d.Name))
}

func (s *Service) createPoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := models.Poll{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.DB.Model(models.Poll{}).Preload("Choice").Create(&p).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Succesfully created.")
}
