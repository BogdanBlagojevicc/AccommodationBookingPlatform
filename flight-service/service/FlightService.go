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

func NewFlightService(l *log.Logger, r *repository.FlightRepository) *FlightService {
	return &FlightService{l, r}
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

func (fs *FlightService) GetFlights(departure string, departurePlace string, arrivalPlace string, noOfSeats int) (model.Flights, error) {
	return fs.Repo.GetAll(departure, departurePlace, arrivalPlace, noOfSeats)
}

func (fs *FlightService) GetNumberOfFreeSeats(id string) (uint64, error) {
	flight, err := fs.Repo.GetFlightById(id)
	if err != nil {
		return 0, err
	}
	return flight.NumberOfFreeSeats, nil
}

func (fs *FlightService) Update(id string, newNumberOfFreeSeats uint64) error {
	return fs.Repo.Update(id, newNumberOfFreeSeats)
}
