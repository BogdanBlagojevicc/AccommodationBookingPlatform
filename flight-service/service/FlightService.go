package service

import (
	"flight-service/model"
	"flight-service/repository"
	"log"
)

type FlightService struct {
	Logger *log.Logger
	Repo   *repository.FlightRepository
}

func NewFlightService(l *log.Logger, s *repository.FlightRepository) *FlightService {
	return &FlightService{l, s}
}
func (fs *FlightService) Insert(newFlight *model.Flight) (*model.Flight, error) {

	return fs.Repo.Insert(newFlight) //newUser
}

func (fs *FlightService) Delete(id string) error {

	return fs.Repo.Delete(id)
}

func (fs *FlightService) GetFlightById(id string) (*model.Flight, error) {
	return fs.Repo.GetFlightById(id)
}
