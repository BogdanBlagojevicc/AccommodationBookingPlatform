package model

import "github.com/google/uuid"

type Flight struct {
	ID             uuid.UUID
	Departure      string
	DeparturePlace string
	ArrivalPlace   string
	Price          uint64
	NumberOfSeats  uint64
}
