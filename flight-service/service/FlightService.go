package service

import (
	"flight-service/repository"
)

type FlightService struct {
	Repo *repository.FlightRepository
}
