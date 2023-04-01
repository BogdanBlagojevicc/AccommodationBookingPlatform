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

func (fs *FlightService) GetFlights(departure string, departurePlace string, arrivalPlace string, noOfSeats int) (model.Flights, error) {
	return fs.Repo.GetAll(departure, departurePlace, arrivalPlace, noOfSeats)
}

func (fs *FlightService) GetFlightById(id string) (*model.Flight, error) {
	return fs.Repo.GetFlightById(id)
}
