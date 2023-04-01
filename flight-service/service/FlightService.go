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

	//user, err := fs.Repo.GetByEmail(newUser.Email)

	//if user != nil {
	//	fs.Logger.Println("User with this email already exists!")
	//	}

	return fs.Repo.Insert(newFlight) //newUser
}

func (fs *FlightService) GetNumberOfFreeSeatsById(id string) (*uint64, error) {
	return fs.Repo.GetNumberOfFreeSeatsById(id)
}
