package sensor

import (
	"log"
	"math/rand"
	"time"
)

type SensorData struct {
	Status Status `json:"status"`
	// Condition string
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Service interface {
	// GetStatus() (Status, error)
	RandomizeStatus()
	StartAutoRandomize()
	// DecideCondition() string
}

type Repository interface {
	GetStatus() (*SensorData, error)
	SetStatus(*SensorData) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetStatus() (Status, error) {
	data, err := s.repo.GetStatus()
	if err != nil {
		return Status{}, err
	}
	return data.Status, nil
}

func (s *service) RandomizeStatus() {
	data, err := s.repo.GetStatus()
	if err != nil {
		return
	}
	data.Status.Water = rand.Intn(100)
	data.Status.Wind = rand.Intn(100)
	s.repo.SetStatus(data)
	log.Println("Randomized status")
}

func (s *service) StartAutoRandomize() {
	// Call RandomizeStatus() every 15 seconds
	for {
		s.RandomizeStatus()
		time.Sleep(15 * time.Second)
	}

}
